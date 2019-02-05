//acute$258
package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/fatih/color"
)

func writeInvoiceToCsvR(invoice []RInvoiceEmp) {
	if invoice == nil {
		fmt.Println("Invoice is not found")
	}
	records := make([][]string, 0)
	headerRecord := []string{"Name (EID)", "From", "To", "Amt"}
	records = append(records, headerRecord)
	for _, oneInvoiceDetail := range invoice {
		oneRecord := []string{
			oneInvoiceDetail.EmployeeName + " (" + oneInvoiceDetail.EmployeeID + ")",
			oneInvoiceDetail.FromDt,
			oneInvoiceDetail.ToDt,
			fmt.Sprintf("%f", oneInvoiceDetail.InvoiceAmount),
		}
		records = append(records, oneRecord)
	}
	finalFilePath := path.Join(filePath, "./rinvoices", fmt.Sprintf("./%s-%s-%s.csv", invoice[0].InvoiceNo, month, year))
	color.Cyan("Creating file %s", finalFilePath)
	file, err := os.Create(finalFilePath)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	if err != nil {
		fmt.Println("error writing to file", finalFilePath, err)
		return
	}
	w.Flush()
	color.Cyan("Done creating file %s", finalFilePath)
}

func writeInvoicesToCsvConsoR(data []RInvoiceEmp) ([][]string, error) {
	if data == nil {
		return nil, fmt.Errorf("Invoice empty cannot generate csv")
	}
	var records [][]string
	for _, oneInvoiceDetail := range data {
		oneRecord := []string{
			oneInvoiceDetail.InvoiceNo,
			oneInvoiceDetail.InvoiceYear,
			oneInvoiceDetail.InvoiceMonth,
			fmt.Sprintf("%f", oneInvoiceDetail.InvoiceAmount),
			oneInvoiceDetail.EmployeeName,
			oneInvoiceDetail.EmployeeID,
			oneInvoiceDetail.FromDt,
			oneInvoiceDetail.ToDt,
		}
		records = append(records, oneRecord)
	}
	return records, nil
}

func writeAllInvoicesToCsvChanR(ctx context.Context, fileName string, wg *sync.WaitGroup, result <-chan *RResult) error {
	defer wg.Done()
	finalFilePath := path.Join(filePath, fmt.Sprintf("%s-%s-%s.csv", fileName, month, year))
	color.Cyan("Creating File for all Invoices: %s", finalFilePath)
	fs, err := os.Create(finalFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fs.Close()
	w := csv.NewWriter(fs)
	headerRecord := []string{
		"InvoiceNo", "InvoiceYear", "InvoiceMonth",
		"InvoiceAmount", "EmployeeName", "EmployeeId",
		"FromDt", "ToDt"}
	w.Write(headerRecord)
	for oneRes := range result {
		records, err := writeInvoicesToCsvConsoR(oneRes.Invoice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		w.WriteAll(records)
		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
	color.Cyan("Done creating file....")
	return nil
}
