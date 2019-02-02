package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyokomi/emoji"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"path"
	"strings"
)

var (
	inputFile         = "./validation.csv"
	paymentFile       = "./payment.csv"
	skipHeader        = true
	scanDir           = "./"
	outputPath        = "./out"
	sheetName         = "ASheet2"
	EmployeeName      = 0
	EmployeeId        = 1
	Month             = 8
	Year              = 9
	ClaimedAmount     = 13
	EligibleAmount    = 17
	ExcessAmount      = 18
	InvoiceNo         = 20
	Payment_EmpID     = 1
	Payment_SynergyID = 2
	AcuteSynergyIDMap = map[string]string{}
)

func flags() {
	sinput := flag.String("v", "./validation.csv", "Validation file for Reimbursment")
	spayment := flag.String("p", "./payment.csv", "Payment file for Reimbursment")
	sscan := flag.String("s", "./", "Directory to scan for excel files")
	sout := flag.String("o", "./out", "Directory where to put completed files")
	sskip := flag.Bool("skip", true, "Weather to skip first line from input file")
	sName := flag.String("sname", "ASheet2", "Sheet Name to be used for generation")
	flag.Parse()
	inputFile = *sinput
	skipHeader = *sskip
	scanDir = *sscan
	outputPath = *sout
	paymentFile = *spayment
	sheetName = *sName
}

type Result struct {
	FileName      string
	Result        bool
	FailureReason string
	OuputFileName string
}

func main() {
	flags()
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		os.Mkdir(outputPath, 0755)
	}
	ReadPayment()
	out := ReadValidation()
	allResults := make([]*Result, 0)
	for oneLine := range out {
		res := OpenAndEditExcel(oneLine)
		allResults = append(allResults, res)
	}
	PrintOutput(allResults)

}

func PrintOutput(AllResults []*Result) {
	file, err := os.Create("Program_output.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error Generating Report", err)
	}
	w := io.MultiWriter(file, os.Stdout)
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"FileName", "Status", "Failure_Reason", "Output_File_Name"})
	for _, v := range AllResults {
		if v.Result == true {
			table.Append([]string{v.FileName, emoji.Sprint(":white_check_mark:"), v.FailureReason, v.OuputFileName})
		} else {
			table.Append([]string{v.FileName, emoji.Sprint(":x:"), v.FailureReason, v.OuputFileName})
		}
	}
	table.Render()
}

func OpenAndEditExcel(line []string) *Result {
	filename := fmt.Sprintf("%s-%s-%s-%s.xlsx", strings.TrimSpace(line[InvoiceNo]), strings.TrimSpace(line[EmployeeId]), strings.TrimSpace(line[Month]), strings.TrimSpace(line[Year]))
	xlsx, err := excelize.OpenFile(path.Join(scanDir, fmt.Sprintf("./%s", filename)))
	if err != nil {
		if os.IsNotExist(err) {
			return &Result{FileName: filename, Result: false, FailureReason: "Not Found"}
		}
		return &Result{FileName: filename, Result: false, FailureReason: err.Error()}

	}
	if _, ok := AcuteSynergyIDMap[line[EmployeeId]]; !ok {
		return &Result{FileName: filename, Result: false, FailureReason: "Acute EmployeeID NotFound"}
	}
	_ = xlsx.NewSheet(sheetName)
	AcuteId := AcuteSynergyIDMap[line[EmployeeId]]
	heading := map[string]string{"A1": "Employee Name", "B1": "Acute ID", "C1": "Employee Id", "D1": "Month", "E1": "Year", "F1": "Claimed Amount", "G1": "Eligible Amount", "H1": "Excess Amount"}
	excelLine := map[string]string{"A2": line[EmployeeName], "B2": AcuteId, "C2": line[EmployeeId], "D2": line[Month], "E2": line[Year], "F2": line[ClaimedAmount], "G2": line[EligibleAmount], "H2": line[ExcessAmount]}
	for k, v := range heading {
		xlsx.SetCellValue(sheetName, k, v)
	}
	for k, v := range excelLine {
		xlsx.SetCellValue(sheetName, k, v)
	}
	outputFile := path.Join(outputPath, fmt.Sprintf("./%s-%s-%s-%s.xlsx", strings.TrimSpace(line[EmployeeId]), strings.TrimSpace(line[InvoiceNo]), strings.TrimSpace(line[Month]), strings.TrimSpace(line[Year])))
	err = xlsx.SaveAs(outputFile)
	if err != nil {
		return &Result{FileName: filename, Result: false, FailureReason: err.Error()}
	}
	return &Result{FileName: filename, Result: true, OuputFileName: outputFile}
}

func ReadPayment() {
	f, err := os.OpenFile(paymentFile, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := csv.NewReader(bufio.NewReader(f))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {

			fmt.Print(err)
			break
		}
		AcuteSynergyIDMap[line[Payment_SynergyID]] = line[Payment_EmpID]
	}
}

func ReadValidation() <-chan []string {
	ch := make(chan []string)
	f, err := os.OpenFile(inputFile, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	reader := csv.NewReader(bufio.NewReader(f))
	if skipHeader {
		_, err := reader.Read()
		if err != nil {
			fmt.Println(err)
		}
	}
	go func() {
		for {
			line, err := reader.Read()
			if err == io.EOF {
				close(ch)
				break
			} else if err != nil {
				close(ch)
				fmt.Print(err)
				break
			}
			ch <- line
		}
	}()
	return ch
}
