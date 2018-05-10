package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"
)

type TemplateDetails struct {
	Date         string
	SrNo         string
	Region       string
	MachineType  string
	ExpiryDate   string
	Address      string
	BankName     string
	Period       string
	Machines     []*MachineDetails
	PaymentTerms string
	Total        float64
}

type MachineDetails struct {
	SrNo    string
	Model   string
	Rate    string
	Qty     string
	GstInt  float64
	RateInt float64
	QtyInt  float64
}

func main() {
	funcMap := template.FuncMap{
		"totWithTax": func(gst float64, rate float64) string {
			return fmt.Sprintf("%f", rate+gst)
		},
		"totAmt": func(gst float64, rate float64, qty float64) string {
			return fmt.Sprintf("%f", (rate+gst)*qty)
		},
	}
	y := []*MachineDetails{{
		SrNo:  "1",
		Model: "Robomak T.T",
		Rate:  "5000",
		Qty:   "1",
	}, {
		SrNo:  "2",
		Model: "Kores - Kashman II",
		Rate:  "6000",
		Qty:   "2",
	},
		{
			SrNo:  "3",
			Model: "Kores - Kashman II",
			Rate:  "6500",
			Qty:   "2",
		}}
	sum := 0.0
	for _, oneM := range y {
		rate, err := strconv.ParseFloat(oneM.Rate, 64)
		if err != nil {
			panic("error converting rate string to float64")
		}
		qty, err := strconv.ParseFloat(oneM.Qty, 64)
		if err != nil {
			panic("error converting qty string to int")
		}
		oneM.RateInt = rate
		oneM.GstInt = rate * 0.18
		oneM.QtyInt = qty
		sum += (oneM.RateInt + oneM.GstInt) * oneM.QtyInt
	}

	x := TemplateDetails{
		SrNo:         "1",
		Region:       "Ahmedabad",
		MachineType:  "NCM",
		ExpiryDate:   "31-03-2018",
		Period:       "01-04-2018 to 31-03-2019",
		Address:      "The Branch Manager\nState Bank of India\nMandvi Main Branch\nAddress Line 1\nAddress Line 2\nAddress Line 3",
		BankName:     "SBI",
		Date:         "02-04-2018",
		PaymentTerms: "Annually Advance",
		Machines:     y,
	}
	x.Total = sum
	tmpl := template.New("NCM").Funcs(funcMap)
	fs, err := os.Open("./sample3.md")
	if err != nil {
		panic(err)
	}

	fsb, err := ioutil.ReadAll(fs)
	if err != nil {
		panic(err)
	}
	tl, err := tmpl.Parse(string(fsb))
	if err != nil {
		panic(err)
	}
	of, err := os.Create("./sample3.out.md")
	if err != nil {
		panic(err)
	}
	tl.Execute(of, x)
}
