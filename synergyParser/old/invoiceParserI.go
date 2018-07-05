package main

import (
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

func init() {
	soup.SetDebug(false)
}

func fetchInvoiceList(respBody string) []string {
	if respBody == "" {
		return []string{""}
	}
	doc := soup.HTMLParse(respBody)
	title := doc.Find("title")
	if strings.TrimSpace(title.Text()) == "Failure Message" {
		return []string{""}
	}
	trs := doc.Find("form", "name", "AgencyInvoice").
		Find("table").
		Find("tbody").
		Find("tbody").FindAll("tr")

	allInvoices := make([]string, len(trs)-1)
	for index, oneTr := range trs {
		if index == 0 {
			continue
		}
		allInvoices[index-1] = extractInvoicesNoFromTrs(oneTr)
	}
	return allInvoices
}

func extractInvoicesNoFromTrs(parent soup.Root) string {
	tdsList := parent.FindAll("td")
	if len(tdsList) == 8 {
		return strings.TrimSpace(tdsList[2].Find("div").Text())
	}
	return ""
}

func fetchInvoiceDetail(respBody string) OneInvoiceDetail {
	if respBody == "" {
		return nil
	}
	doc := soup.HTMLParse(respBody)
	trs := doc.Find("form", "name", "AgencyInvoice").
		Find("table").Find("tbody").Find("tbody").FindAll("tr")
	details := make(OneInvoiceDetail, 0)
	for _, oneTr := range trs {
		oneTrd := fetchInvoiceDetailsEmployee(oneTr)
		if oneTrd != nil {
			details = append(details, oneTrd)
		}
	}
	return details
}

func fetchInvoiceDetailsEmployee(parent soup.Root) *EmployeeDetail {
	tdsList := parent.FindAll("td")
	if len(tdsList) != 12 {
		return nil
	}
	braces := regexp.MustCompile("[()]")
	data := &EmployeeDetail{}
	data.EmployeeName = strings.TrimSpace(extractAllText(tdsList[0].Find("div")))
	empData := strings.Split(data.EmployeeName, "~")
	if len(empData) == 4 {
		data.EmployeeName = empData[0]
		data.EmployeeID = braces.ReplaceAllString(empData[1], "")
	}
	data.InvoiceMonth = strings.TrimSpace(extractAllText(tdsList[2].Find("div")))
	empData = strings.Split(data.InvoiceMonth, "~")
	if len(empData) == 3 {
		data.InvoiceMonth = empData[0]
		data.InvoiceYear = braces.ReplaceAllString(empData[1], "")
	}
	data.PeriodFrom = strings.TrimSpace(tdsList[3].Find("div").Text())
	data.PeriodTo = strings.TrimSpace(tdsList[4].Find("div").Text())
	data.WorkingDuration = strings.TrimSpace(tdsList[5].Find("div").Text())
	data.ContractorRate = strings.TrimSpace(tdsList[6].Find("div").Text())
	data.InvoiceAmount = strings.TrimSpace(tdsList[7].Find("div", "style", "display:block").Text())
	return data
}
