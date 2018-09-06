/*
Link:- https://synergy.wipro.com/synergy/PartnerWILogin.jsp
User Type:- Partner Supervisor
Login ID:- 13618
Password:-
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
)

var (
	initURL                    = "https://synergy.wipro.com/synergy/PartnerWILogin.jsp"
	loginURL                   = "https://synergy.wipro.com/synergy/LoginServlet"
	invoicesURL                = "https://synergy.wipro.com/synergy/CN_AgencyInvoicesView.jsp"
	invoiceDetailURL           = "https://synergy.wipro.com/synergy/CN_AgencyInvoiceSingleView.jsp?hSelectedPartner=%s&hSelectedInvoiceNumber=%s&hSelectedInvoiceStatus=PARKED"
	reimbursementLogin1        = "https://synergy.wipro.com/synergy/LoginServlet?Operation=Reimbursement%20Invoice%20View"
	reimbursementLogin2        = "https://synergy.wipro.com/synergy/Authentication.do?Operation=AuthenticateUser"
	reimbursementURL           = "https://synergy.wipro.com/synergy/AgencyViewAction.do"
	reimbursementDetailURL     = "https://synergy.wipro.com/synergy/AgencyViewAction.do?Operation=InvoiceViewScreen&hSelectedPartner=%s&hSelectedInvoiceNumber=%s&hSelectedInvoiceStatus=PARKED"
	reinbursementInvoiceEmpURL = "https://synergy.wipro.com/synergy/ReimbursementViewServlet"
	usernameG                  = ""
	passwordG                  = ""
	recordsPerPage             = ""
	month                      = ""
	year                       = ""
	filePath                   = ""
	reimbursement              = false
	reimbursementExcel         = false
	noOfProxy                  = 6
	threadCnt                  = 3
	waitThreshold              = 10
	version                    = ""
)

var proxy *utils.Proxy
var uas *utils.UA

func setUpFlags() {
	username := flag.String("u", usernameG, "Enter login username")
	password := flag.String("pwd", passwordG, "Enter Login password")
	yearStr := flag.String("y", "2018", "Year for which you want to fetch invoices")
	monthInt := flag.Int("m", 3, "Month for which you want to fetch invoices i.e 1 - Janurary, 2 February")
	recordsPerPageStr := flag.String("recordsPerPage", "100", "No of invoices to fetch since pagination is not supported we fetch all records in one go. Don't change if you dont know what you're doing it should be a positive number")
	reimbursementB := flag.Bool("r", false, "If Need to fetch reimbursement details pass true")
	reimbursementBE := flag.Bool("re", false, "If Need to fetch reimbursement invoice employee claims pass true")
	outFilePath := flag.String("p", "./", "outfile path")
	flag.Parse()
	usernameG = *username
	passwordG = *password
	year = *yearStr
	recordsPerPage = *recordsPerPageStr
	reimbursement = *reimbursementB
	reimbursementExcel = *reimbursementBE
	filePath = *outFilePath
	monthStr := time.Month(*monthInt)
	month = monthStr.String()
}

func _init() {
	color.Magenta("Inititalizing Proxy")
	proxy = utils.NewProxy(noOfProxy, threadCnt, waitThreshold)
	color.Magenta("Fetching Proxy list")
	proxy.LoadProxies()
	color.Magenta("Ranking Proxies")
	proxy.RankProxy()
	color.Magenta("Initializing User Agents")
	uas = utils.NewRandomUA()
	color.Magenta("Loading User agents")
	uas.LoadDummyUserAgents()
}

func main() {
	fmt.Println("Running Version:", version)
	setUpFlags()
	initDirectory(filePath)
	color.Magenta("......Start......\n Press Ctrl+c to stop")
	_init()
	client := makeClient()
	cookies, err := loadPage(client)
	if err != nil {
		fmt.Println("Error Loading the login page", err)
		return
	}
	res, err := performLogin(client, cookies)
	if err != nil {
		fmt.Println("Error couldnt login in", err)
	} else if res == false {
		fmt.Println("Could not Log In:")
		return
	}
	if reimbursement {
		color.Red("Reimbursement Invoices")
		res, err := performRLoginRequest(client, cookies)
		if err != nil {
			fmt.Println("Error something happened", err)
			return
		} else if res == false {
			fmt.Println("could not login")
			return
		}
		allInvoicesNo, err := fetchReimbursmentsList(client, cookies)
		if err != nil {
			panic(err)
		}
		if len(allInvoicesNo) > 0 {
			ctx1 := context.Background()
			ctx, cancel := context.WithCancel(ctx1)
			defer cancel()
			invoicesNoChan := invoicesChan(ctx, allInvoicesNo)
			ch1 := makeRequestChanR(ctx, client, cookies, invoicesNoChan)
			ch2 := makeRequestChanR(ctx, client, cookies, invoicesNoChan)
			ch3 := makeRequestChanR(ctx, client, cookies, invoicesNoChan)
			ch4 := makeRequestChanR(ctx, client, cookies, invoicesNoChan)
			chAll := mergeRequestChan(ctx, ch1, ch2, ch3, ch4)
			split1 := make(chan *Result)
			split2 := make(chan *Result)
			var split3 chan *Result
			var wg sync.WaitGroup
			wg.Add(1)
			go writeInvoicesAllToCsvChan(ctx, "Reimbursments-AllInvoices", &wg, split1)
			for i := 0; i < 2; i++ {
				wg.Add(1)
				go writeRInvoiceOneToCsvChan(ctx, &wg, split2)
			}
			if reimbursementExcel {
				split3 = make(chan *Result, 5)
				for i := 0; i < 2; i++ {
					wg.Add(1)
					go downloadRInvoiceOneToExcelChan(ctx, &wg, client, cookies, split3)
				}
			}
			if reimbursementExcel {
				go duplicateChannels(ctx, chAll, split1, split2, split3)
			} else {
				go duplicateChannels(ctx, chAll, split1, split2)
			}
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt)
			go func() {
				select {
				case <-c:
					cancel()
				}
			}()
			wg.Wait()
			fmt.Println("Done...")
			return
		}
	} else {
		color.Red("Regular Invoices")
		allInvoicesNo, err := fetchInvoicesList(client, cookies)
		if err != nil {
			panic(err)
		}
		if len(allInvoicesNo) > 0 {
			ctx1 := context.Background()
			ctx, cancel := context.WithCancel(ctx1)
			defer cancel()
			invoicesNoChan := invoicesChan(ctx, allInvoicesNo)
			ch1 := makeRequestChan(ctx, client, cookies, invoicesNoChan)
			ch2 := makeRequestChan(ctx, client, cookies, invoicesNoChan)
			ch3 := makeRequestChan(ctx, client, cookies, invoicesNoChan)
			ch4 := makeRequestChan(ctx, client, cookies, invoicesNoChan)
			chAll := mergeRequestChan(ctx, ch1, ch2, ch3, ch4)
			split1 := make(chan *Result)
			split2 := make(chan *Result)
			var wg sync.WaitGroup
			wg.Add(1)
			go writeInvoicesAllToCsvChan(ctx, "AllInvoices", &wg, split1)
			for i := 0; i < 2; i++ {
				wg.Add(1)
				go writeInvoiceOneToCsvChan(ctx, &wg, split2)
			}
			go duplicateChannels(ctx, chAll, split1, split2)
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt)
			go func() {
				select {
				case <-c:
					cancel()
				}
			}()
			wg.Wait()
			fmt.Println("Done..")
			return
		}
	}
	fmt.Println("Oops empty invoice list")
}
