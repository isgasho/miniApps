package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

func DownloadExcel(client *http.Client, invoice []RInvoiceEmp) {
	if invoice == nil {
		fmt.Println("Invoice is not found")
	}
	for _, oneInv := range invoice {
		query := getReimbursementInvoiceEmpExcel(oneInv.EmployeeID, oneInv.InvoiceMonth, oneInv.InvoiceYear, oneInv.FromDt, oneInv.ToDt)
		finalFilePath := path.Join(filePath, "./rexcels", fmt.Sprintf("./%s-%s-%s-%s.xls", oneInv.InvoiceNo, oneInv.EmployeeID, oneInv.InvoiceMonth, oneInv.InvoiceYear))
		DownloadFile(client, finalFilePath, query)
	}
}

func DownloadFile(client *http.Client, filePathl string, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request", err)
		return
	}
	startTime := time.Now()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching request", err)
		return
	}
	fmt.Printf("GET %s Took %s\n", url, time.Since(startTime).String())
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading bindary data from body", err)
		return
	}
	err = ioutil.WriteFile(filePathl, buf, 0666)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	fmt.Println("done creating file", filePathl)
}
