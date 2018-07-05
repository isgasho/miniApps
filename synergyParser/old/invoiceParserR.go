package main

import (
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
)

func init() {
	soup.SetDebug(false)
}

func fetchReimbursmentList(respBody string) []string {
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
		allInvoices[index-1] = extractRinvoiceNoFromTrs(oneTr)
	}
	return allInvoices

}

func extractRinvoiceNoFromTrs(parent soup.Root) string {
	tdsList := parent.FindAll("td")
	if len(tdsList) == 7 {
		return strings.TrimSpace(tdsList[2].Find("div").Text())
	}
	return ""
}

func fetchRInvoiceDetail(respBody string) OneInvoiceDetail {
	if respBody == "" {
		return nil
	}
	doc := soup.HTMLParse(respBody)
	trs := doc.Find("form", "name", "AgencyInvoice").
		Find("table").Find("tbody").Find("tbody").FindAll("tr")
	details := make(OneInvoiceDetail, 0)
	for _, oneTr := range trs {
		oneTrd := fetchRInvoiceDetailsEmployee(oneTr)
		if oneTrd != nil {
			details = append(details, oneTrd)
		}
	}
	return details
}

func fetchRInvoiceDetailsEmployee(parent soup.Root) *EmployeeDetail {
	tdsList := parent.FindAll("td")
	if len(tdsList) != 7 {
		return nil
	}
	braces := regexp.MustCompile("[()]")
	numbersOnly := regexp.MustCompile("[^0-9]")
	data := &EmployeeDetail{}
	data.EmployeeName = strings.TrimSpace(tdsList[0].Find("div").Text())
	vals := braces.Split(data.EmployeeName, -1)
	if len(vals) >= 2 {
		data.EmployeeID = vals[1]
		data.EmployeeName = vals[0]
	}
	data.PeriodFrom = strings.TrimSpace(tdsList[3].Find("div").Text())
	data.PeriodTo = strings.TrimSpace(tdsList[4].Find("div").Text())
	data.InvoiceAmount = numbersOnly.ReplaceAllLiteralString(strings.TrimSpace(tdsList[5].Find("div").Text()), "")
	return data
}

func extractAllText(r soup.Root) string {
	val := ""
	k := r.Pointer.FirstChild
checkNode:
	if k == nil {
		return val
	}
	if k.Type == html.TextNode {
		val = val + strings.TrimSpace(k.Data) + "~"
		k = k.NextSibling
		goto checkNode
	} else if k != nil {
		k = k.NextSibling
		goto checkNode
	}
	return val
}
