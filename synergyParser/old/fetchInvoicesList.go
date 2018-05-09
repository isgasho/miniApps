package main

import (
	"github.com/anaskhan96/soup"
	"strings"
)

type EmployeeDetail struct {
	EmployeeName    string
	EmployeeId      string
	InvoiceMonth    string
	InvoiceYear     string
	PeriodFrom      string
	PeriodTo        string
	WorkingDuration string
	ContractorRate  string
	InvoiceAmount   string
}

type AllEmployeeDetails []EmployeeDetail

//this function will fetch invoices coming from page - invoicesUrl variable
func FetchInvoiceList(respBody string) []string {
	doc := soup.HTMLParse(respBody)
	trs := doc.Find("form", "name", "AgencyInvoice").
		Find("table").
		Find("tbody").
		Find("tbody").FindAll("tr")

	allInvoices := make([]string, len(trs)-1)

	for index, oneTr := range trs {
		if index == 0 {
			continue
		}
		allInvoices[index-1] = ExtractInvoicesNoFromTrs(oneTr)
	}
	return allInvoices

}

func ExtractInvoicesNoFromTrs(parent soup.Root) string {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	tdsList := parent.FindAll("td")
	if len(tdsList) == 8 {
		return strings.TrimSpace(tdsList[2].Find("div").Text())
	}
	return ""
}

func FetchInvoiceDetail(respBody string) *AllEmployeeDetails {
	doc := soup.HTMLParse(respBody)
	trs := doc.Find("form", "name", "AgencyInvoice").
		Find("table").Find("tbody").Find("tbody").FindAll("tr")
	details := make(AllEmployeeDetails, 0)
	for _, oneTr := range trs {
		details = append(details, *FetchInvoiceDetailsEmployee(oneTr))
	}
	return &details
}

func FetchInvoiceDetailsEmployee(parent soup.Root) *EmployeeDetail {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	tdsList := parent.FindAll("td")
	if len(tdsList) != 12 {
		return nil
	}
	data := EmployeeDetail{}
	data.EmployeeName = strings.TrimSpace(tdsList[0].Find("div").Text())
	data.InvoiceMonth = strings.TrimSpace(tdsList[2].Find("div").Text())
	data.PeriodFrom = strings.TrimSpace(tdsList[3].Find("div").Text())
	data.PeriodTo = strings.TrimSpace(tdsList[4].Find("div").Text())
	data.WorkingDuration = strings.TrimSpace(tdsList[5].Find("div").Text())
	data.ContractorRate = strings.TrimSpace(tdsList[6].Find("div").Text())
	data.InvoiceAmount = strings.TrimSpace(tdsList[7].Find("div", "id", "print_con61").Text())
	return &data
}
