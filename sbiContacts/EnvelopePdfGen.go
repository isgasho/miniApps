package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io"
	"os"
	"time"
)

type BranchDetails struct {
	BranchCode string
	BranchName string
	Pincode    string
	Address    []string
}

type AllBranchDetails struct {
	Result []BranchDetails
}

type FuncStr func(string)
type FuncOnly func()

func main() {
	stTime := time.Now()
	fmt.Printf("Start generating file\n")
	pdf, write, setBold, setRegular := getPdfDocRead()
	allBranches, err := readCsv("./sbiBranchDetail-0-0.csv")
	if err != nil {
		fmt.Println("CSv Reading Error: ", err)
		return
	}
	for _, branch := range allBranches.Result {
		pdf.AddPage()
		setBold()
		write("To,")
		write("The Branch Manager")
		write("State Bank of India (" + branch.BranchCode + ")")
		setRegular()
		write(branch.BranchName)
		for i := 0; i < len(branch.Address); i++ {
			write(branch.Address[i])
		}
		write("Pin - " + branch.Pincode)
	}
	err = pdf.OutputFileAndClose("./hello.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Done Generating Pdf in : %v\n", time.Since(stTime))
}

func readCsv(filePath string) (*AllBranchDetails, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	branchArray := AllBranchDetails{}
	r := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		branch := BranchDetails{}
		branch.BranchCode = record[0]
		branch.BranchName = record[1]
		branch.Pincode = record[3]
		for i := 11; i < len(record); i++ {
			branch.Address = append(branch.Address, record[i])
		}
		branchArray.Result = append(branchArray.Result, branch)
	}
	return &branchArray, nil
}

func getPdfDocRead() (*gofpdf.Fpdf, FuncStr, FuncOnly, FuncOnly) {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 241.3, Ht: 104.902}, // 241.3 x 104.902
		OrientationStr: "P",
	})
	var ht float64
	pdf.SetMargins(80, 20, 10)
	setBold := func() {
		pdf.SetFont("Arial", "B", 16)
		ht = pdf.PointConvert(16)
	}
	setRegular := func() {
		pdf.SetFont("Arial", "", 14)
		ht = pdf.PointConvert(14)
	}
	write := func(str string) {
		pdf.MultiCell(140, ht, str, "", "L", false)
	}
	return pdf, write, setBold, setRegular
}
