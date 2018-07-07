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
	purposeLogin          = "login"
	purposeGetInvoices    = "invoices"
	purposeGetOneInvoice  = "oneInvoice"
	initURL               = "https://appstore.wipro.com/worklight/apps/services/preview/ContractPartner/common/0/default/index.html"
	apiUrl                = "https://appstore.wipro.com/worklight/apps/services/api/ContractPartner/common/query"
	usernameG             = ""
	passwordG             = ""
	month                 = ""
	year                  = ""
	filePath              = ""
	reimbursement         = false
	generateCreateInvFile = false
	noOfProxy             = 6
	threadCnt             = 3
	waitThreshold         = 10
)

var proxy *utils.Proxy
var uas *utils.UA

func setupFlags() {
	username := flag.String("u", "13618", "Enter Username")
	password := flag.String("pwd", "acute#258", "Enter Password")
	yearStr := flag.String("y", "2018", "Enter year for which you want the invoice i.e 2017,2018")
	monthInt := flag.Int("m", 1, "Enter Month number for which you want the invoice i.e 1 -January, 2- February")
	reimbursementB := flag.Bool("r", false, "If Need to fetch reimbursement details pass true")
	outFilePath := flag.String("p", "./", "outfile path")
	generateCreateInvFileG := flag.Bool("g", false, "If need to generate create invoice file pass true")
	flag.Parse()
	usernameG = *username
	passwordG = *password
	year = *yearStr
	reimbursement = *reimbursementB
	generateCreateInvFile = *generateCreateInvFileG
	filePath = *outFilePath
	monthStr := time.Month(*monthInt)
	month = monthStr.String()
	fmt.Println(month,year)
	return
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

func getCreateInvoicePurpose() string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\"}}","cinv"]`, usernameG)
}

func getLoginPurpose() string {
	x := fmt.Sprintf(`["{\"inputs\":{\"userId\":\"%s\",\"pwd\":\"%s\",\"userType\":\"PS\"}}","login"]`, usernameG, passwordG)
	fmt.Println(x)
	return x
}

func getReimbursementListPurpose() string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"\",\"strinvoicestatus\":\"null\",\"strinvoicemonth\":\"%s\",\"strinvoiceyear\":\"%s\"}}","creimbmultibasedcri"]`, usernameG, month, year)
}

func getReimbursementInvoicePurpose(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","creimburseviewpemp"]`, usernameG, invoiceNo)
}

func getInvoicesListPurpose() string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"\",\"strinvoicestatus\":\"null\",\"strinvoicemonth\":\"%s\",\"strinvoiceyear\":\"%s\"}}","cinviewmul"]`, usernameG, month, year)
}

func getInvoicePurposePtax(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","cinviewptax"]`, usernameG, invoiceNo)
}

func getInvoicePurposePemp(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","cinviewpemp"]`, usernameG, invoiceNo)
}

func getInvoicePurposePid(invoiceNo string) string {
	return fmt.Sprintf(`["{\"inputs\":{\"nconsfocrecid\":\"%s\",\"strinvoicenumber\":\"%s\"}}","cinviewpid"]`, usernameG, invoiceNo)
}

func main() {
	setupFlags()
	color.Magenta("......Start......\n Press Ctrl+c to stop")
	_init()
	client := makeClient()
	cookie, err := loadPage(client)
	if err != nil {
		fmt.Println("Error Loading the Page: ", err)
		return
	}
	logintoken, err := performLogin1(client, cookie)
	if err != nil {
		fmt.Println("Error getting login token: ", err)
		return
	}
	res, err := performLogin2(client, cookie, logintoken)
	if err != nil {
		fmt.Println("Error Loggin in: ", err)
		return
	}
	if res == false {
		fmt.Println("Invalid Login credentails")
	} else {
		if generateCreateInvFile == true {
			fmt.Println("Generate Invoices")
			getCreateInvoiceList(client, cookie, logintoken)
			return
		}
		if reimbursement == true {
			fmt.Println("Reimbursement Invoices")
			ctx1 := context.Background()
			ctx, cancel := context.WithCancel(ctx1)
			defer cancel()
			invoicesNoChan, err := invoicesChanR(ctx, client, cookie, logintoken)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch1 := makeRequestChanR(ctx, client, cookie, logintoken, invoicesNoChan)
			ch2 := makeRequestChanR(ctx, client, cookie, logintoken, invoicesNoChan)
			ch3 := makeRequestChanR(ctx, client, cookie, logintoken, invoicesNoChan)
			chAll := mergeRequestChanR(ctx, ch1, ch2, ch3)
			split1 := make(chan *RResult)
			split2 := make(chan *RResult)
			var wg sync.WaitGroup
			wg.Add(1)
			go writeAllInvoicesToCsvChanR(ctx, "ReimbursementAllInvoices", &wg, split1)
			for i := 0; i < 2; i++ {
				wg.Add(1)
				go writeInvoiceToCsvChanR(ctx, &wg, split2)
			}
			go duplicateChannelsR(ctx, chAll, split1, split2)
			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt)
			go func() {
				select {
				case <-c:
					cancel()
				}
			}()
			wg.Wait()
		} else {
			fmt.Println("Regular Invoices")
			ctx1 := context.Background()
			ctx, cancel := context.WithCancel(ctx1)
			defer cancel()
			invoicesNoChan, err := invoicesChan(ctx, client, cookie, logintoken)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch1 := makeRequestChan(ctx, client, cookie, logintoken, invoicesNoChan)
			ch2 := makeRequestChan(ctx, client, cookie, logintoken, invoicesNoChan)
			ch3 := makeRequestChan(ctx, client, cookie, logintoken, invoicesNoChan)
			chAll := mergeRequestChan(ctx, ch1, ch2, ch3)
			split1 := make(chan *Result)
			split2 := make(chan *Result)
			var wg sync.WaitGroup
			wg.Add(1)
			go writeAllInvoicesToCsvChan(ctx, "AllInvoices", &wg, split1)
			for i := 0; i < 2; i++ {
				wg.Add(1)
				go writeInvoiceToCsvChan(ctx, &wg, split2)
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
		}
		color.Magenta("......Done.....")
	}
}
