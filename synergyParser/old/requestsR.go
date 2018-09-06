package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
)

func fetchReimbursmentsList(client *http.Client, cookies string) ([]string, error) {
	form := url.Values{}
	form.Add("PageNo", "1")
	form.Add("Operation", "ViewRecords")
	form.Add("RecordsPerPage", recordsPerPage)
	form.Add("selectMonth", month)
	form.Add("selectYear", year)
	form.Add("selectStatus", "NONE")
	form.Add("InvoiceNumber", "NONE")
	req, err := http.NewRequest("POST", reimbursementURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookies)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		fmt.Println("Error making request")
		return nil, err
	}
	color.Cyan("Fetching All Reimbursement Invoices %s\nPOST %s\nTime Taken:%s\n", cookies, reimbursementURL, *timeTaken)
	if *responseStr == "" {
		return nil, errors.New("Response string is empty")
	}
	invoices := fetchReimbursmentList(*responseStr)
	return invoices, nil
}

func fetchOneRInvoice(client *http.Client, invoiceNo string, cookies string) (OneInvoiceDetail, string, error) {
	if invoiceNo == "" {
		return nil, "", errors.New("No invoice number found to fetch")
	}
	url := fmt.Sprintf(reimbursementDetailURL, usernameG, invoiceNo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, invoiceNo, err
	}
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		fmt.Println("Error Making Request")
		return nil, invoiceNo, err
	}
	color.Green("Get Single Invoice Detail For REIMBURSEMENT Invoice No:%s\nGET %s\nTime Taken:%s\n", invoiceNo, url, *timeTaken)
	oneInvoice := fetchRInvoiceDetail(*responseStr)
	return oneInvoice, invoiceNo, nil
}

func downloadRinvoicetoExcel(client *http.Client, invoice OneInvoiceDetail, invoiceNo string, cookies string) {
	for _, oneEmp := range invoice {
		err := downloadRinvoicetoExcelReq(client, oneEmp, invoiceNo, cookies)
		if err != nil {
			fmt.Printf("Error creating emp invoce file %s", oneEmp.EmployeeID)
		}
	}
}

func downloadRinvoicetoExcelReq(client *http.Client, empDtl *EmployeeDetail, invoiceNo string, cookies string) error {
	form := url.Values{}
	form.Add("CompanyCode", "null")
	form.Add("Operation", "viewContractorAdditionalExpenses")
	form.Add("Value", "ReimbursementInvoice")
	form.Add("endDate", empDtl.PeriodTo)
	form.Add("n_timesheet_year", empDtl.InvoiceYear)
	form.Add("startDate", empDtl.PeriodFrom)
	form.Add("str_contractor_id", empDtl.EmployeeID)
	form.Add("str_timesheet_month", empDtl.InvoiceMonth)
	/*reader := strings.NewReader(form.Encode())
	bytes, _ := ioutil.ReadAll(reader)
	fmt.Println(string(bytes))*/

	req, err := http.NewRequest("POST", reinbursementInvoiceEmpURL, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error creating request", err)
		return err
	}
	req.Header = getHeaders(false, cookies)
	startTime := time.Now()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching request", err)
		return err
	}
	fmt.Printf("GET %s Took %s\n", reinbursementInvoiceEmpURL, time.Since(startTime).String())
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading bindary data from body", err)
		return err
	}
	finalFilePath := path.Join(filePath, "./rexcels", fmt.Sprintf("%s-%s-%s-%s.xls", invoiceNo, empDtl.EmployeeID, empDtl.InvoiceMonth, empDtl.InvoiceYear))
	err = ioutil.WriteFile(finalFilePath, buf, 0666)
	if err != nil {
		fmt.Println("Error writing file", err)
		return err
	}
	color.Cyan("Done creating file %s", finalFilePath)

	return nil
}

func writeRinvoiceOnetoCsv(invoice OneInvoiceDetail, invoiceNo string) error {
	records := make([][]string, 0)
	headerRecord := []string{"Name (EID)", "From", "To", "Days", "Amt"}
	records = append(records, headerRecord)
	for _, oneEmp := range invoice {
		oneRecord := []string{
			oneEmp.EmployeeName + "(" + oneEmp.EmployeeID + ")",
			oneEmp.PeriodFrom,
			oneEmp.PeriodTo,
			oneEmp.WorkingDuration,
			oneEmp.InvoiceAmount,
		}
		records = append(records, oneRecord)
	}
	finalFilePath := path.Join(filePath, "./rinvoices", fmt.Sprintf("Reimbursment-Synergy-%s-%s-%s.csv", invoiceNo, month, year))
	color.Cyan("Creating File %s", finalFilePath)
	err := utils.WriteCsv(records, finalFilePath)
	if err != nil {
		return err
	}
	color.Cyan("Done creating file %s", finalFilePath)
	return nil
}
