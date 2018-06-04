package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

var (
	purposeLogin         = "login"
	purposeGetInvoices   = "invoices"
	purposeGetOneInvoice = "oneInvoice"
	initUrl              = "https://appstore.wipro.com/worklight/apps/services/preview/ContractPartner/common/0/default/index.html"
	apiUrl               = "https://appstore.wipro.com/worklight/apps/services/api/ContractPartner/common/query"
	usernameG            = "13618"
	passwordG            = "dummy"
	noOfProxy            = 6
	threadCnt            = 3
	waitThreshold        = 10
)

var client *http.Client
var proxy *utils.Proxy
var uas *utils.UA

func getCreateInvoicePurpose() string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\"}}","cinv"]`, usernameG)
}

func getLoginPurpose() string {

	x := fmt.Sprintf(`["{\"inputs\":{\"userId\":\"%s\",\"pwd\":\"%s\",\"userType\":\"PS\"}}","login"]`, usernameG, passwordG)
	fmt.Println(x)
	return x
}

func getInvoicesListPurpose(month, year string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"\",\"strinvoicestatus\":\"null\",\"strinvoicemonth\":\"%s\",\"strinvoiceyear\":\"%s\"}}","cinviewmul"]`, usernameG, month, year)
}

func getInvoicePurposePtax(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","cinviewptax"]`, usernameG, invoiceNo)
}

func getInvoicePurposePemp(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","cinviewpemp"]`, usernameG, invoiceNo)
}

func getInvoicePurposePid(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","cinviewpid"]`, usernameG, invoiceNo)
}

func _init_() {
	proxy = utils.NewProxy(noOfProxy, threadCnt, waitThreshold)
	color.Magenta("Fetching Proxy list")
	proxy.LoadProxies()
	color.Magenta("Ranking Proxies")
	proxy.RankProxy()

	color.Magenta("Setting up User Agents")
	uas = utils.NewRandomUA()
	uas.LoadDummyUserAgents()

}

func main() {
	color.Magenta("......Start.....")
	var monthNumber = flag.Int("month", 1, "Enter Month number for which you want the invoice i.e 1 -January, 2- February")
	var year = flag.String("year", "2018", "Enter year for which you want the invoice i.e 2017,2018")
	var outfilePath = flag.String("path", "./", "Enter path where you would like to get the generated file")
	var username = flag.String("username", "13618", "Enter Username")
	var password = flag.String("password", "acute#258", "Enter Password")
	var generateCreateInvFile = flag.Bool("generateCreateInvFile", false, "If need to generate create invoice file pass true")
	flag.Parse()
	usernameG = *username
	passwordG = *password
	month := time.Month(*monthNumber)
	_init_()
	client = initClient()
	cookie, err := loadPage()
	if err != nil {
		fmt.Println("Error Loading the Page: ", err)
		return
	}
	logintoken, err := performLogin1(cookie)
	if err != nil {
		fmt.Println("Error getting login token: ", err)
		return
	}
	res, err := performLogin2(cookie, logintoken)
	if err != nil {
		fmt.Println("Error Loggin in: ", err)
		return
	}
	if res == false {
		fmt.Println("Invalid Login credentails")
		return
	} else {
		if *generateCreateInvFile == true {
			getCreateInvoiceList(*outfilePath, month.String(), *year, cookie, logintoken)
			return
		}
		allInvoicesSummary, err := getInvoiceList(cookie, logintoken, month.String(), *year)
		if err != nil {
			fmt.Println("Error fetching invoices", err)
			return
		}
		motherOfAllInvoices := make(map[string]*FinalInvoices)
		for _, oneInvoiceSummary := range allInvoicesSummary.Result {
			oneInvoiceDetail, err := getSingleInvoice(cookie, logintoken, oneInvoiceSummary.InvoiceNumber)
			if err != nil {
				fmt.Println("error fetching details for invoice no: ", oneInvoiceSummary.InvoiceNumber)
				continue
			}
			motherOfAllInvoices[oneInvoiceSummary.InvoiceNumber] = oneInvoiceDetail
		}
		writeInvoicesToCsv(*outfilePath, month.String(), *year, allInvoicesSummary, motherOfAllInvoices)
		writeInvoicesToCsvOneAtATime(*outfilePath, month.String(), *year, allInvoicesSummary, motherOfAllInvoices)
		color.Magenta("......Done.....")
	}
}

func getCreateInvoiceList(filePath, month, year string, cookies, wlInst string) error {
	newReqFromData := formData()
	newReqFromData.Set("parameters", getCreateInvoicePurpose())
	startTime := time.Now()
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(newReqFromData.Encode()))
	if err != nil {
		return err
	}
	req.Header = getHeaders(false, cookies, wlInst)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	color.Yellow("Get Create Invoice Request \nPOST %s\nTime Taken:%s\n", apiUrl, time.Since(startTime))
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", body)
	responseStr := string(body)
	jsonExtract := gjson.Get(responseStr, "result")
	finalJson := fmt.Sprintf(`{"result": ` + jsonExtract.String() + `}`)
	allInvoices := AllCreateInvoiceEmpList{}
	err = json.Unmarshal([]byte(finalJson), &allInvoices)
	if err != nil {
		return fmt.Errorf("Error wrapping invoicesList response to JSON")
	}
	records := make([][]string, 10)
	headerRecord := []string{"ContractorID",
		"ContractorState",
		"EmpCompanyCode",
		"EmployeeName",
		"InvoiceAmount",
		"ProjectCode",
		"Rate",
		"RateTypeDesc",
		"ResumeNumber",
		"Rmdayshours",
		"TimeSheetMonth",
		"TimeSheetYear",
	}
	records = append(records, headerRecord)
	for _, oneInvoiceDetail := range allInvoices.Result {
		oneRecord := []string{
			oneInvoiceDetail.ContractorId,
			oneInvoiceDetail.ContractorState,
			oneInvoiceDetail.EmpCompanyCode,
			oneInvoiceDetail.EmployeeName,
			oneInvoiceDetail.InvoiceAmount,
			oneInvoiceDetail.ProjectCode,
			oneInvoiceDetail.Rate,
			oneInvoiceDetail.RateTypeDesc,
			oneInvoiceDetail.ResumeNumber,
			oneInvoiceDetail.Rmdayshours,
			oneInvoiceDetail.TimeSheetMonth,
			oneInvoiceDetail.TimesheetYear,
		}
		records = append(records, oneRecord)
	}
	finalFilePath := path.Join(filePath, fmt.Sprintf("./createInvoice-%s-%s.csv", month, year))
	color.Cyan("Creating file %s", finalFilePath)
	file, err := os.Create(finalFilePath)
	if err != nil {
		fmt.Println("Error creating file", err)
		return err
	}
	w := csv.NewWriter(file)
	fmt.Println(cap(records), len(records))
	for _, record := range records {
		if len(record) == 0 {
			continue
		}
		if err := w.Write(record); err != nil {
			fmt.Println("oops error writing to file", err)
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("error writing to file", err)
		return err
	}
	color.Cyan("Done creating file....")
	return nil
}

func writeInvoicesToCsvOneAtATime(filePath, month, year string, AllInvoiceSummary *AllInvoiceSummary, motherOfAllInvoices map[string]*FinalInvoices) {
	mapOfAllInvoices := make(map[string][][]string)
	for _, oneInvoiceSummary := range AllInvoiceSummary.Result {
		data := motherOfAllInvoices[oneInvoiceSummary.InvoiceNumber]
		records := make([][]string, 10)
		headerRecord := []string{"Name (EID)", "From", "To", "Days", "Amt"}
		records = append(records, headerRecord)
		for _, oneInvoiceDetail := range data.EmpInvoices {
			oneRecord := []string{
				oneInvoiceDetail.EmployeeName,
				oneInvoiceDetail.TimeSheetStDt,
				oneInvoiceDetail.TimeSheetEndDt,
				oneInvoiceDetail.RmDays + " Days",
				fmt.Sprintf("%f", oneInvoiceDetail.InvoiceAmount),
			}
			records = append(records, oneRecord)
			oneRecord = []string{oneInvoiceDetail.EmployeeId, "", "", "", ""}
			records = append(records, oneRecord)
		}
		/*
			oneRecord := []string{"Total", "", "", "", fmt.Sprintf("%f", oneInvoiceSummary.TotalInvoiceAmt)}
			records = append(records, oneRecord)
			oneRecord = []string{"Central GST @9%", "", "", "", fmt.Sprintf("%f", oneInvoiceSummary.TotalInvoiceAmt*9/100)}
			records = append(records, oneRecord)
			oneRecord = []string{"State GST @9%", "", "", "", fmt.Sprintf("%f", oneInvoiceSummary.TotalInvoiceAmt*9/100)}
			records = append(records, oneRecord)
			oneRecord = []string{"Total", "", "", "", fmt.Sprintf("%f", oneInvoiceSummary.TotalInvoiceAmt+(oneInvoiceSummary.TotalInvoiceAmt*18/100))}
			records = append(records, oneRecord)
		*/
		mapOfAllInvoices[oneInvoiceSummary.InvoiceNumber] = records
	}

	for invoiceNumber, invoiceLines := range mapOfAllInvoices {
		finalFilePath := path.Join(filePath, fmt.Sprintf("./%s-%s-%s.csv", invoiceNumber, month, year))
		color.Cyan("Creating file %s", finalFilePath)
		file, err := os.Create(finalFilePath)
		if err != nil {
			fmt.Println("Error creating file", err)
			return
		}
		w := csv.NewWriter(file)
		for _, record := range invoiceLines {
			if len(record) == 0 {
				continue
			}
			if err := w.Write(record); err != nil {
				fmt.Println("oops error writing to file", err)
				return
			}
		}
		w.Flush()
		if err := w.Error(); err != nil {
			fmt.Println("error writing to file", err)
			return
		}
		color.Cyan("Done creating file %s", finalFilePath)
	}
}

func writeInvoicesToCsv(filePath, month, year string, allInvoicesSummary *AllInvoiceSummary, motherOfAllInvoices map[string]*FinalInvoices) {
	records := make([][]string, 10)
	headerRecord := []string{
		"InvoiceNo", "InvoiceDt", "ServiceType", "ContractorState", "InvoiceYear", "InvoiceMonth",
		"InvoiceAmount", "TaxAmount", "WiproGSTNO", "ProjectCode", "SezIndicator", "WiproGSTNO2",
		"EmployeeName", "EmployeeId", "RmDays", "RmHours",
		"RateStr", "TimeSheetStDt", "TimeSheetEndDt"}
	records = append(records, headerRecord)
	for _, oneInvoiceSummary := range allInvoicesSummary.Result {
		data := motherOfAllInvoices[oneInvoiceSummary.InvoiceNumber]
		if data == nil {
			fmt.Println("No data found to invoice no", oneInvoiceSummary.InvoiceNumber)
			continue
		}
		invoicePid := data.OneInvoicePid
		for _, oneInvoiceDetail := range data.EmpInvoices {
			oneRecord := []string{
				invoicePid.InvoiceNo,
				invoicePid.InvoiceDt,
				invoicePid.SerivceType,
				oneInvoiceDetail.ContractorState,
				oneInvoiceDetail.InvoiceYear,
				oneInvoiceDetail.InoviceMonth,
				fmt.Sprintf("%f", oneInvoiceDetail.InvoiceAmount),
				fmt.Sprintf("%f", oneInvoiceDetail.TaxAmount),
				oneInvoiceDetail.WiproGSTNo,
				oneInvoiceDetail.ProjectCode,
				oneInvoiceDetail.SezIndicator,
				oneInvoiceDetail.WiproGSTNo,
				oneInvoiceDetail.EmployeeName,
				oneInvoiceDetail.EmployeeId,
				oneInvoiceDetail.RmDays,
				oneInvoiceDetail.RmHours,
				oneInvoiceDetail.RateStr,
				oneInvoiceDetail.TimeSheetStDt,
				oneInvoiceDetail.TimeSheetEndDt,
			}
			records = append(records, oneRecord)
		}
	}
	finalFilePath := path.Join(filePath, fmt.Sprintf("./viewInvoices-%s-%s.csv", month, year))
	color.Cyan("Creating file %s", finalFilePath)
	file, err := os.Create(finalFilePath)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	w := csv.NewWriter(file)
	fmt.Println(cap(records), len(records))
	for _, record := range records {
		if len(record) == 0 {
			continue
		}
		if err := w.Write(record); err != nil {
			fmt.Println("oops error writing to file", err)
			return
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("error writing to file", err)
		return
	}
	color.Cyan("Done creating file....")
}

func getInvoiceList(cookies, wlInst, month, year string) (*AllInvoiceSummary, error) {
	newReqFormData := formData()
	newReqFormData.Set("parameters", getInvoicesListPurpose(month, year))
	startTime := time.Now()
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(newReqFormData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookies, wlInst)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	color.Yellow("Get Invoices Request For %s-%s\nPOST %s\nTime Taken:%s\n", month, year, apiUrl, time.Since(startTime))
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	responseStr := string(body)
	jsonExtract := gjson.Get(responseStr, "result")
	finalJson := fmt.Sprint(`{"result" : ` + jsonExtract.String() + `}`)
	allInvoices := AllInvoiceSummary{}
	err = json.Unmarshal([]byte(finalJson), &allInvoices)
	if err != nil {
		return nil, fmt.Errorf("Error wrapping invoicesList response into JSON")
	}
	return &allInvoices, nil
}

func getSingleInvoice(cookie, wlInst, invoiceNo string) (*FinalInvoices, error) {
	newReqFormData := formData()
	newReqFormData.Set("parameters", getInvoicePurposePemp(invoiceNo))
	startTime := time.Now()
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(newReqFormData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookie, wlInst)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	color.Green("Get Single Invoice PEMP Detail Request For Invoice No:%s\nPOST %s\nTime Taken:%s\n", invoiceNo, apiUrl, time.Since(startTime))
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	responseStr := string(body)
	jsonExtract := gjson.Get(responseStr, "result")
	finalJson := fmt.Sprint(`{"result" : ` + jsonExtract.String() + `}`)
	allEmpInvoices := AllEmpInvoice{}
	err = json.Unmarshal([]byte(finalJson), &allEmpInvoices)
	if err != nil {
		return nil, fmt.Errorf("Error wrapping  InvoiceEmpList response into JSON %v", err)
	}

	newReqFormData.Set("parameters", getInvoicePurposePid(invoiceNo))
	startTime = time.Now()
	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(newReqFormData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookie, wlInst)
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	color.Red("Get Single Invoice PID Detail Request For Invoice No:%s\nPOST %s\nTime Taken:%s\n", invoiceNo, apiUrl, time.Since(startTime))
	responseStr = string(body)
	jsonExtract = gjson.Get(responseStr, "result.0")
	finalJson = jsonExtract.String()
	invoicePid := InvoicePid{}
	err = json.Unmarshal([]byte(finalJson), &invoicePid)
	if err != nil {
		return nil, fmt.Errorf("Error wrapping InvoicePId respinse into json: %v", err)
	}

	return &FinalInvoices{OneInvoicePid: &invoicePid, EmpInvoices: allEmpInvoices.Result}, nil
}
