package main

import (
	"bytes"
	"fmt"
	"github.com/devarsh/miniApps/template/csvReader"
	"github.com/devarsh/miniApps/template/mdToPdf"
	"os"
	"text/template"
)

type Mixture struct {
	Quotation *csvReader.Quotation
	Date      string
}

func main() {
	ncm, err := Asset("templates/ncm.md")
	if err != nil {
		panic(err)
	}
	ccm, err := Asset("templates/ccm.md")
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
	tmpl := template.New("NCM").Funcs(funcMap)
	ncmTmpl, err := tmpl.Parse(string(ncm))
	if err != nil {
		panic(err)
	}
	ccmTmpl, err := tmpl.Parse(string(ccm))
	if err != nil {
		panic(err)
	}
	csvRd := csvReader.NewTemplateReader("02-May-2018")
	err = csvRd.ReadCsv("./input.csv")
	if err != nil {
		fmt.Println(err)
	}

	mdPdf := mdToPdf.NewMdtoPdf(12, "Arial")
	mixture := &Mixture{}
	mixture.Date = csvRd.Date
	for csvRd.Next() {
		var b bytes.Buffer
		oneRecord := csvRd.GetRecord()
		mixture.Quotation = oneRecord
		if oneRecord.MachineType == "CCM" {
			ccmTmpl.Execute(&b, mixture)
		} else if oneRecord.MachineType == "NCM" {
			ncmTmpl.Execute(&b, mixture)
		} else {
			continue
		}
		bytes := b.Bytes()
		filename := fmt.Sprintf("./out/Quotation-%s-%s-%s", oneRecord.SrNo, oneRecord.MachineType, oneRecord.Region)
		fs, err := os.Create(filename + ".md")
		if err != nil {
			panic(err)
		}
		fs.Write(bytes)
		mdPdf.NewPdf(bytes, filename, "For any complaints call us on (L)079 26424229/(M)99252 04929/(M)99099 58229")
	}
}
