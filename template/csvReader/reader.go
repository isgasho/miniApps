package csvReader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Quotation struct {
	SrNo          string
	RefNo         string
	Region        string
	MachineType   string
	BankName      string
	Address       string
	ExpiryDate    string
	Period        string
	Machines      []*MachineDetails
	PaymentTerms  string
	Total         float64
	QuotationType string
}

type MachineDetails struct {
	SrNo         string
	Model        string
	Rate         string
	Qty          string
	TotalWithTax float64
	Gst          float64
	Total        float64
}

type TemplateReader struct {
	allQuotation map[string]*Quotation
	keys         []string
	cursor       int
	Date         string
}

func NewTemplateReader(date string) *TemplateReader {
	inst := &TemplateReader{}
	inst.allQuotation = make(map[string]*Quotation)
	inst.Date = date
	inst.cursor = -1
	inst.keys = make([]string, 0)
	return inst
}

func (t *TemplateReader) ReadCsv(filepath string, ignoreFirstLine bool) error {
	fp, err := os.Open(filepath)
	if err != nil {
		return err
	}
	//Reading Pass1 i.e Read First Line on Quoatation
	cnt := 1
	r := csv.NewReader(fp)
	if ignoreFirstLine == true {
		_, err := r.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if record[0] != "" && record[1] != "" {
			oneQuoatation := &Quotation{}
			oneQuoatation.Machines = make([]*MachineDetails, 0)
			oneQuoatation.SrNo = record[0]
			oneQuoatation.Region = record[1]
			oneQuoatation.MachineType = record[2]
			oneQuoatation.RefNo = record[3]
			oneQuoatation.Address = strings.TrimSpace(record[4])
			oneQuoatation.ExpiryDate = record[5]
			oneQuoatation.Period = record[6]
			oneQuoatation.PaymentTerms = record[7]
			oneQuoatation.Total = 0.0
			oneQuoatation.QuotationType = strings.ToUpper(strings.TrimSpace(record[11]))
			oneMachine := &MachineDetails{}
			oneMachine.SrNo = record[0]
			oneMachine.Model = record[8]
			oneMachine.Rate = record[9]
			oneMachine.Qty = record[10]
			/*Calculate GST & Total & Append total to grand Total*/
			rate, err := strconv.ParseFloat(oneMachine.Rate, 64)
			if err != nil {
				fmt.Print("Error Line No -", cnt, oneMachine.Rate, oneMachine.SrNo)
				return err
			}
			qty, err := strconv.ParseFloat(oneMachine.Qty, 64)
			if err != nil {
				fmt.Print("Error Line No -", cnt, oneMachine.Qty, oneMachine.SrNo)
				return err
			}
			oneMachine.Gst = rate * 0.18
			oneMachine.TotalWithTax = rate + oneMachine.Gst
			oneMachine.Total = oneMachine.TotalWithTax * qty
			oneQuoatation.Total += oneMachine.Total

			strs := strings.Split(strings.TrimSpace(oneQuoatation.Address), "\n")
			for index, _ := range strs {
				strs[index] = strings.Replace(strs[index], "~", "", -1)
				if index == 0 || index == 1 {
					strs[index] = fmt.Sprintf("**%s**", strings.TrimSpace(strs[index]))
				} else {
					strs[index] = strings.TrimSpace(strs[index])
				}
			}
			oneQuoatation.Address = strings.Join(strs, "\n")
			/*Region Modify*/
			if strings.ToUpper(oneQuoatation.Region) == "BARODA" {
				oneQuoatation.Region = "BRD"
			} else if strings.ToUpper(oneQuoatation.Region) == "AHMEDABAD" {
				oneQuoatation.Region = "AHD"
			} else {
				oneQuoatation.Region = "XYZ"
			}

			/*append one machine*/
			oneQuoatation.Machines = append(oneQuoatation.Machines, oneMachine)
			if _, ok := t.allQuotation[oneQuoatation.SrNo]; ok {
				fmt.Printf("Skipping record because duplicate record for SrNo: %s\n", oneQuoatation.SrNo)
			} else {
				t.allQuotation[oneQuoatation.SrNo] = oneQuoatation
				t.keys = append(t.keys, oneQuoatation.SrNo)
			}
		}
		cnt++
	}
	//Reading Pass 2 i.e Read Additional Lines of the Quotation for M/c Details
	fp.Seek(0, 0)
	r = csv.NewReader(fp)
	cnt = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if record[0] != "" && record[1] == "" {
			oneMachine := &MachineDetails{}
			oneMachine.SrNo = record[0]
			oneMachine.Model = record[8]
			oneMachine.Rate = record[9]
			oneMachine.Qty = record[10]
			/*Calculate GST & Total & Append total to grand Total*/
			rate, err := strconv.ParseFloat(oneMachine.Rate, 64)
			if err != nil {
				fmt.Print("Error Line No -", cnt, oneMachine.Rate, oneMachine.SrNo)
				return err
			}
			qty, err := strconv.ParseFloat(oneMachine.Qty, 64)
			if err != nil {
				fmt.Print("Error Line No -", cnt, oneMachine.Qty, oneMachine.SrNo)
				return err
			}
			oneMachine.Gst = rate * 0.18
			oneMachine.TotalWithTax = rate + oneMachine.Gst
			oneMachine.Total = oneMachine.TotalWithTax * qty

			t.allQuotation[oneMachine.SrNo].Total += oneMachine.Total
			t.allQuotation[oneMachine.SrNo].Machines = append(t.allQuotation[oneMachine.SrNo].Machines, oneMachine)
		}
		cnt++
	}
	sort.Strings(t.keys)
	return nil
}

func (t *TemplateReader) ViewRecords() {
	for _, val := range t.allQuotation {
		fmt.Println(val)
		for _, val := range val.Machines {
			fmt.Println("--------------------------------", val)
		}
	}
}

func (t *TemplateReader) Next() bool {
	if t.cursor < len(t.allQuotation)-1 {
		t.cursor++
		return true
	}
	return false
}

func (t *TemplateReader) GetCursorPosition() int {
	return t.cursor
}

func (t *TemplateReader) GetRecord() *Quotation {
	val := t.allQuotation[t.keys[t.cursor]]
	return val

}
