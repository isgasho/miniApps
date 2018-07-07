package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

//For Beena File
func getCreateInvoiceList(client *http.Client, cookies, wlInst string) error {
	newReqFromData := formData()
	newReqFromData.Set("parameters", getCreateInvoicePurpose())
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(newReqFromData.Encode()))
	if err != nil {
		return err
	}
	req.Header = getHeaders(false, cookies, wlInst)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		return err
	}
	color.Yellow("Get Create Invoice Request \nPOST %s\nTime Taken:%s\n", apiUrl, timeTaken)
	jsonExtract := gjson.Get(*responseStr, "result")
	finalJSON := fmt.Sprintf(`{"result": ` + jsonExtract.String() + `}`)
	allInvoices := AllCreateInvoiceEmpList{}
	err = json.Unmarshal([]byte(finalJSON), &allInvoices)
	if err != nil {
		return fmt.Errorf("Error wrapping invoicesList response to JSON")
	}
	records := make([][]string, 0)
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
	finalFilePath := path.Join(filePath, fmt.Sprintf("./createInvoiceListCurrent.csv"))
	color.Cyan("Creating file %s", finalFilePath)
	file, err := os.Create(finalFilePath)
	if err != nil {
		fmt.Println("Error creating file", err)
		return err
	}
	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	if err != nil {
		fmt.Println("error Writing to the file", err)
		return err
	}
	w.Flush()
	color.Cyan("Done creating file....")
	return nil
}
