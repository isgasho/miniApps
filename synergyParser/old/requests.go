package main

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
)

func loadPage(client *http.Client) (string, error) {
	startTime := time.Now()
	req, err := http.NewRequest("GET", initURL, nil)
	if err != nil {
		return "", err
	}
	req.Header = getHeaders(true, "")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	color.Cyan("Load Login Page Request\nGET %s\nTime Taken:%s\n", initURL, time.Since(startTime))
	defer res.Body.Close()
	cookies := appendCookies("", res.Cookies())
	return cookies, nil
}

func performLogin(client *http.Client, cookies string) (bool, error) {
	//setting up form values for login
	form := url.Values{}
	form.Add("CompanyCode", "WI")
	form.Add("Operation", "PartnerCheckUser")
	form.Add("SelectComp", "0")
	form.Add("UserType", "PS")
	form.Add("UserName", usernameG)
	form.Add("Password", passwordG)
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(form.Encode()))
	if err != nil {
		return false, err
	}
	req.Header = getHeaders(false, cookies)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		return false, err
	}
	color.Cyan("Login Req With Cookie %s\nPOST %s\nTime Taken:%s\n", cookies, loginURL, *timeTaken)
	result := strings.Contains(*responseStr, "Authentication Failed")
	return !result, nil
}

func performRLoginRequest(client *http.Client, cookies string) (bool, error) {
	//setting up form values for login
	req, err := http.NewRequest("GET", reimbursementLogin1, nil)
	if err != nil {
		return false, err
	}
	req.Header = getHeaders(false, cookies)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		return false, err
	}
	color.Cyan("Login Req-1 for Reimbursment With Cookie %s\nGET %s\nTime Taken:%s\n", cookies, reimbursementLogin1, *timeTaken)
	result := strings.Contains(*responseStr, "Failure Message")
	if result != true {
		req, err = http.NewRequest("GET", reimbursementLogin2, nil)
		if err != nil {
			return false, err
		}
		req.Header = getHeaders(false, cookies)
		responseStr, timeTaken, err = utils.RequestMaker(client, req)
		if err != nil {
			return false, err
		}
		color.Cyan("Login Req-2 for Reimbursment With Cookie %s\nGET %s\nTime Taken:%s\n", cookies, reimbursementLogin2, *timeTaken)
	}
	return !result, nil
}

func writeInvoicesToCsvConso(invoice OneInvoiceDetail, invoiceNo string) ([][]string, error) {
	records := make([][]string, 0)
	for _, oneEmp := range invoice {
		if oneEmp == nil {
			return nil, errors.New("empty Employees dtl found invoice detail")
		}
		oneRecord := []string{
			invoiceNo,
			oneEmp.EmployeeName,
			oneEmp.EmployeeID,
			oneEmp.InvoiceMonth,
			oneEmp.InvoiceYear,
			oneEmp.PeriodFrom,
			oneEmp.PeriodTo,
			oneEmp.WorkingDuration,
			oneEmp.ContractorRate,
			oneEmp.InvoiceAmount,
		}
		records = append(records, oneRecord)
	}
	return records, nil
}
