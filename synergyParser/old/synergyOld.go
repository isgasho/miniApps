/*
Link:- https://synergy.wipro.com/synergy/PartnerWILogin.jsp
User Type:- Partner Supervisor
Login ID:- 13618
Password:- 5alD_PlbOVu3-
*/

package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

var (
	initUrl          = "https://synergy.wipro.com/synergy/PartnerWILogin.jsp"
	loginUrl         = "https://synergy.wipro.com/synergy/LoginServlet"
	invoicesUrl      = "https://synergy.wipro.com/synergy/CN_AgencyInvoicesView.jsp"
	invoiceDetailUrl = "https://synergy.wipro.com/synergy/CN_AgencyInvoiceSingleView.jsp?hSelectedPartner=%s&hSelectedInvoiceNumber=%s&hSelectedInvoiceStatus=PARKED"
	usernameG        = "13618"
	passwordG        = "5alD_PlbOVu3-"
)

var client *http.Client

func getHeaders(init bool, cookieString string) http.Header {
	customHeader := http.Header{}
	customHeader.Add("upgrade-insecure-requests", "1")
	customHeader.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	customHeader.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	customHeader.Add("dnt", "1")
	customHeader.Add("accept-encoding", "gzip, deflate, br")
	customHeader.Add("accept-language", "en-US,en;q=0.9,es;q=0.8,hi;q=0.7")
	if init == false {
		customHeader.Add("content-type", "application/x-www-form-urlencoded")
		customHeader.Add("origin", "https://synergy.wipro.com")
		customHeader.Add("referer", initUrl)
	}
	if cookieString != "" {
		customHeader.Add("cookie", cookieString)
	}
	return customHeader
}

func appendCookies(cookies string, cookiesArray []*http.Cookie) string {
	for _, oneCookie := range cookiesArray {
		cookies = cookies + oneCookie.Name + "=" + oneCookie.Value + ";"
	}
	return cookies
}

func main() {
	color.Magenta("......Start.....")
	client = initClient()
	cookies, err := loadPage()
	if err != nil {
		panic(err)
	}
	res, err := performLogin(cookies)
	if err != nil {
		panic(err)
	}
	fmt.Println("Logged In:", res, err)
	invoices, err := fetchInvoicesList(cookies)
	if err != nil {
		panic(err)
	}
	err = fetchAllSingleInvoice(invoices, cookies)
	if err != nil {
		panic(err)
	}

}

func initClient() *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}
	proxyUrl, err := url.Parse(fmt.Sprintf("http://%s:%s", "51.15.86.88", "3128"))
	client := http.Client{Jar: jar, Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	return &client
}

func loadPage() (string, error) {
	startTime := time.Now()
	req, err := http.NewRequest("GET", initUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header = getHeaders(true, "")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	color.Cyan("Load Login Page Request\nGET %s\nTime Taken:%s\n", initUrl, time.Since(startTime))
	defer res.Body.Close()
	cookies := appendCookies("", res.Cookies())

	return cookies, nil
}

func performLogin(cookies string) (bool, error) {
	form := url.Values{}
	form.Add("CompanyCode", "WI")
	form.Add("Operation", "PartnerCheckUser")
	form.Add("SelectComp", "0")
	form.Add("UserType", "PS")
	form.Add("UserName", usernameG)
	form.Add("Password", passwordG)

	startTime := time.Now()
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return false, err
	}
	req.Header = getHeaders(false, cookies)
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	color.Cyan("Login Req With Cookie %s\nPOST %s\nTime Taken:%s\n", cookies, loginUrl, time.Since(startTime))
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	responseStr := string(body)
	result := strings.Contains(responseStr, "Authentication Failed")
	return !result, nil
}

func fetchInvoicesList(cookies string) ([]string, error) {
	form := url.Values{}
	form.Add("PageNo", "1")
	form.Add("RecordsPerPage", "30")
	form.Add("selectMonth", "February")
	form.Add("selectYear", "2018")
	form.Add("selectStatus", "NONE")
	form.Add("InvoiceNumber", "NONE")

	startTime := time.Now()
	req, err := http.NewRequest("POST", invoicesUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookies)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	color.Cyan("Fetching All Invoices %s\nPOST %s\nTime Taken:%s\n", cookies, invoicesUrl, time.Since(startTime))
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	responseStr := string(body)
	invoices := FetchInvoiceList(responseStr)
	return invoices, nil
}

func fetchAllSingleInvoice(invoices []string, cookies string) error {
	startTime := time.Now()
	for _, oneInvoice := range invoices {
		url := fmt.Sprintf(invoiceDetailUrl, usernameG, oneInvoice)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		color.Green("Get Single Invoice Detail For Invoice No:%s\nGET %s\nTime Taken:%s\n", oneInvoice, url, time.Since(startTime))
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		responseStr := string(body)
		oneInvoiceEmps := FetchInvoiceDetail(responseStr)
		for _, oneEmp := range *oneInvoiceEmps {
			fmt.Println(oneEmp)
		}
		return nil
	}
	return nil
}
