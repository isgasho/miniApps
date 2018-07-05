package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/fatih/color"
)

func invoicesChan(ctx context.Context, invoices []string) <-chan string {
	invChan := make(chan string)
	go func() {
		defer close(invChan)
		for _, oneInv := range invoices {
			select {
			case invChan <- oneInv:
			case <-ctx.Done():
				return
			}
		}
	}()
	return invChan
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

func writeInvoicesAllToCsvChan(ctx context.Context, fileName string, wg *sync.WaitGroup, result <-chan *Result) error {
	defer wg.Done()
	finalFilePath := path.Join(filePath, fmt.Sprintf("%s-%s-%s.csv", fileName, month, year))
	color.Cyan("Creating File for all Invoices: %s", finalFilePath)
	fs, err := os.Create(finalFilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fs.Close()
	w := csv.NewWriter(fs)
	headerRecord := []string{
		"InvoiceNumber", "EmployeeName", "EmployeeId", "InvoiceMonth",
		"InvoiceYear", "PeriodFrom", "PeriodTo", "WorkingDuration",
		"ContractorRate", "InvoiceAmount",
	}
	w.Write(headerRecord)
	for oneRes := range result {
		records, err := writeInvoicesToCsvConso(oneRes.InvDetail, oneRes.InvoiceNo)
		if err != nil {
			fmt.Println("Error writing invoice No:", oneRes.InvoiceNo, " error: ", err)
			continue
		}
		w.WriteAll(records)
		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
	color.Cyan("Done creating file....")
	return nil
}
