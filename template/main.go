package main

import (
	"bytes"
	"fmt"
	"github.com/devarsh/miniApps/template/csvReader"
	"github.com/devarsh/miniApps/template/mdToPdf"
	"io/ioutil"
	"os"
	"text/template"
)

type Mixture struct {
	Quotation *csvReader.Quotation
	Date      string
	MonthYear string
}

func main() {
	fs, err := os.Open("./sample3.md")
	if err != nil {
		panic(err)
	}
	fsb, err := ioutil.ReadAll(fs)
	if err != nil {
		panic(err)
	}
	funcMap := template.FuncMap{
		"machineNames": func(mc []*csvReader.MachineDetails) string {
			mcStr := "Model -"
			for _, oneMc := range mc {
				mcStr = mcStr + oneMc.Model + ","
			}
			return mcStr
		},
	}
	tmpl, err := template.New("NCM").Funcs(funcMap).Parse(string(fsb))
	if err != nil {
		panic(err)
	}

	csvRd := csvReader.NewTemplateReader("0518", "02-05-2018")
	err = csvRd.ReadCsv("./input.csv")
	if err != nil {
		fmt.Println(err)
	}

	mdPdf := mdToPdf.NewMdtoPdf(12)
	mixture := &Mixture{}
	mixture.Date = csvRd.Date
	mixture.MonthYear = csvRd.MonthYear
	for csvRd.Next() {
		var b bytes.Buffer
		oneRecord := csvRd.GetRecord()
		mixture.Quotation = oneRecord
		tmpl.Execute(&b, mixture)
		bytes := b.Bytes()
		fs, err := os.Create(fmt.Sprintf("./out/Quotation-%s-%s.md", oneRecord.SrNo, oneRecord.BankName))
		if err != nil {
			panic(err)
		}
		fs.Write(bytes)

		mdPdf.NewPdf(bytes, fmt.Sprintf("./out/Quotation-%s-%s", oneRecord.SrNo, oneRecord.BankName), "For Any Complains Call us on So & So numbers")
	}
}
