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
	initURL          = "https://synergy.wipro.com/synergy/PartnerWILogin.jsp"
	loginURL         = "https://synergy.wipro.com/synergy/LoginServlet"
	invoicesURL      = "https://synergy.wipro.com/synergy/CN_AgencyInvoicesView.jsp"
	invoiceDetailURL = "https://synergy.wipro.com/synergy/CN_AgencyInvoiceSingleView.jsp?hSelectedPartner=%s&hSelectedInvoiceNumber=%s&hSelectedInvoiceStatus=PARKED"
	usernameG        = ""
	passwordG        = ""
	month            = ""
	year             = ""
	filePath         = ""
	noOfProxy        = 6
	threadCnt        = 3
	waitThreshold    = 10
)

var proxy *utils.Proxy
var uas *utils.UA

func setUpFlags() {
	username := flag.String("username", "13618", "Enter login username")
	password := flag.String("password", "5alD_PlbOVu3-", "Enter Login password")
	yearStr := flag.String("year", "2018", "Year for which you want to fetch invoices")
	monthInt := flag.Int("month", 3, "Moth for which you want to fetch invoices i.e 1 - Janurary, 2 February")
	outFilePath := flag.String("path", "./", "outfile path")
	flag.Parse()
	usernameG = *username
	passwordG = *password
	year = *yearStr
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
	setUpFlags()
	color.Magenta("......Start.....\n Press Ctrl+c to stop")

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
	}
	fmt.Println("Logged In:", res, err)
	allInvoicesNo, err := fetchInvoicesList(client, cookies)
	if err != nil {
		panic(err)
	}
	fmt.Println(allInvoicesNo)
	if len(allInvoicesNo) > 0 {
		ctx1 := context.Background()
		ctx, cancel := context.WithCancel(ctx1)
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
		go writeInvoicesAllToCsvChan(ctx, &wg, split1)
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
	fmt.Println("Oops empty invoice list")
}
