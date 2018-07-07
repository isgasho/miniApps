package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func fetchInvoiceList(client *http.Client, cookies, wlInst string) ([]gjson.Result, error) {
	newReqFormData := formData()
	newReqFormData.Set("parameters", getInvoicesListPurpose())
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(newReqFormData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookies, wlInst)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)

	color.Yellow("Get Invoices Request For %s-%s\nPOST %s\nTime Taken:%s\n", month, year, apiUrl, timeTaken)
	jsonExtract := gjson.Get(*responseStr, "result.#.str_invoice_number")
	if jsonExtract.IsArray() {
		return jsonExtract.Array(), nil
	}
	return nil, fmt.Errorf("could'nt convert json response into Array of Invoice List")
}

func invoicesChan(ctx context.Context, client *http.Client, cookie, logintoken string) (<-chan string, error) {
	invoiceList, err := fetchInvoiceList(client, cookie, logintoken)
	if err != nil {
		return nil, err
	}
	invChan := make(chan string)
	go func() {
		defer close(invChan)
		for _, oneInv := range invoiceList {
			select {
			case invChan <- oneInv.String():
			case <-ctx.Done():
				return
			}
		}
	}()
	return invChan, nil
}

func makeRequestChan(ctx context.Context, client *http.Client, cookies, loginToken string, invoiceChan <-chan string) <-chan *Result {
	resChan := make(chan *Result)
	go func() {
		defer close(resChan)
		for oneInv := range invoiceChan {
			invDtl, err := fetchOneInvoice(client, cookies, loginToken, oneInv)
			select {
			case resChan <- &Result{Err: err, Invoice: invDtl}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return resChan
}

func fetchOneInvoice(client *http.Client, cookie, wlInst, invoiceNo string) (*FinalInvoices, error) {
	if invoiceNo == "" {
		return nil, fmt.Errorf("Error Fetching Employee for empty Invoice")
	}
	newReqFormData := formData()
	newReqFormData.Set("parameters", getInvoicePurposePemp(invoiceNo))
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(newReqFormData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookie, wlInst)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		return nil, fmt.Errorf("Error Fetching PEMP Invoice No:%s", invoiceNo)
	}
	color.Green("Get Single Invoice PEMP Detail Request For Invoice No:%s\nPOST %s\nTime Taken:%s\n", invoiceNo, apiUrl, timeTaken)
	jsonExtract := gjson.Get(*responseStr, "result")
	finalJSON := fmt.Sprint(`{"result" : ` + jsonExtract.String() + `}`)
	allEmpInvoices := AllEmpInvoice{}
	err = json.Unmarshal([]byte(finalJSON), &allEmpInvoices)
	if err != nil {
		return nil, fmt.Errorf("Error wrapping  InvoiceEmpList response into JSON %v", err)
	}
	newReqFormData.Set("parameters", getInvoicePurposePid(invoiceNo))
	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(newReqFormData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = getHeaders(false, cookie, wlInst)
	responseStr, timeTaken, err = utils.RequestMaker(client, req)
	if err != nil {
		return nil, fmt.Errorf("Error Fetching PID Invoice No:%s", invoiceNo)
	}
	color.Red("Get Single Invoice PID Detail Request For Invoice No:%s\nPOST %s\nTime Taken:%s\n", invoiceNo, apiUrl, timeTaken)
	jsonExtract = gjson.Get(*responseStr, "result.0")
	finalJSON = jsonExtract.String()
	invoicePid := InvoicePid{}
	err = json.Unmarshal([]byte(finalJSON), &invoicePid)
	if err != nil {
		return nil, fmt.Errorf("Error wrapping InvoicePId respinse into json: %v", err)
	}
	return &FinalInvoices{OneInvoicePid: &invoicePid, EmpInvoices: allEmpInvoices.Result}, nil
}

func mergeRequestChan(ctx context.Context, cs ...<-chan *Result) <-chan *Result {
	resChan := make(chan *Result)
	var wg sync.WaitGroup
	for _, oneCs := range cs {
		wg.Add(1)
		go func(oneChan <-chan *Result) {
			defer wg.Done()
			for res := range oneChan {
				select {
				case resChan <- res:
				case <-ctx.Done():
					return
				}
			}
		}(oneCs)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	return resChan
}

func duplicateChannels(ctx context.Context, inChannel <-chan *Result, out ...chan *Result) {
	for inChRes := range inChannel {
		for _, oneChan := range out {
			oneChan <- inChRes
			select {
			case <-ctx.Done():
				goto a
			default:
			}
		}
	}
a:
	for _, oneChan := range out {
		close(oneChan)
	}
}

func writeInvoiceToCsvChan(ctx context.Context, wg *sync.WaitGroup, res <-chan *Result) {
	defer wg.Done()
	for oneRes := range res {
		writeInvoiceToCsv(oneRes.Invoice)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
