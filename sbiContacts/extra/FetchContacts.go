package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/fatih/color"
	"os"
	"path"
	"strings"
	"time"
)

type BranchDetail struct {
	BranchCode   string
	BranchName   string
	Address      []string
	City         string
	Pincode      string
	District     string
	State        string
	StdCode      string
	PhoneNumber  string
	EmailId      string
	IFSCCode     string
	BranchCircle string
}

func main() {
	stRange := flag.Int("StRange", 0, "Starting Range for branch code")
	enRange := flag.Int("EnRange", 0, "End Range for brach code")
	masterFile := flag.String("masterFile", "./data.txt", "Master file location containg")
	outFilePath := flag.String("outFile", "./", "Csv out file destination path")
	flag.Parse()
	brList := make(map[string]*BranchDetail)
	data, err := getBranchesList(*masterFile, *stRange, *enRange)
	if err != nil {
		fmt.Println(err)
	}
	data1 := *data
	for i := 0; i < len(data1); i++ {
		detail, err := ExtractBranchDetail(data1[i])
		if err != nil {
			fmt.Println(err)
		}
		brList[data1[i]] = detail
	}
	generateCsvFile(&brList, *outFilePath, *stRange, *enRange)
}

func generateCsvFile(BranchDetails *map[string]*BranchDetail, outFilePath string, stRange, enRange int) {
	records := make([][]string, 10)
	oneRecord := []string{
		"BranchCode", "BranchName", "City", "Pincode", "District", "State", "StdCode",
		"PhoneNumber", "EmailId", "IFSC_Code", "BranchCircle", "Address-1", "Address-2",
		"Address-3", "Address-4", "Address-5",
	}
	records = append(records, oneRecord)
	for _, val := range *BranchDetails {
		oneRecord := []string{
			val.BranchCode,
			val.BranchName,
			val.City,
			val.Pincode,
			val.District,
			val.State,
			val.StdCode,
			val.PhoneNumber,
			val.EmailId,
			val.IFSCCode,
			val.BranchCircle,
		}
		for i := 0; i < len(val.Address); i++ {
			oneRecord = append(oneRecord, val.Address[i])
		}
		records = append(records, oneRecord)
	}
	filePath := path.Join(outFilePath, fmt.Sprintf("sbiBranchDetail-%d-%d.csv", stRange, enRange))
	color.Cyan("Writing data to a csv file: %s", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	w := csv.NewWriter(file)
	for _, record := range records {
		if len(record) == 0 {
			continue
		}
		if err := w.Write(record); err != nil {
			fmt.Println("oops error writing to file", err)
			return
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("error writing to file", err)
		return
	}
	color.Cyan("Done creating file %s", filePath)
}

func getBranchesList(filename string, stRange, endRange int) (*[]string, error) {
	if stRange <= 0 {
		stRange = 1
	}
	if endRange <= 0 {
		endRange = 99999
	}
	color.Yellow("Started Reading Branch codes from %s file for the range %d to %d", filename, stRange, endRange)
	stTime := time.Now()
	var branchesList []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 1
	for scanner.Scan() {
		if count >= stRange && count <= endRange {
			branchesList = append(branchesList, scanner.Text())
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	color.Yellow("Finished Reading Braches List\nTime Taken: %s\n", time.Since(stTime))
	return &branchesList, nil
}

func ExtractBranchDetail(branchCode string) (*BranchDetail, error) {
	color.Cyan("Extracting details for branch: %s", branchCode)
	stTime := time.Now()
	resp, err := soup.Get("https://www.sbi.co.in/corporate/branchlocatorfinal.htm?bcode=" + branchCode)
	if err != nil {
		return nil, err
	}
	doc := soup.HTMLParse(resp)
	trs := doc.Find("form").Find("table").Find("table").
		Find("table").Find("table").Find("tbody").
		FindAll("tr")
	isChild := false
	prevKey := ""
	branchDetail := BranchDetail{}
	for _, oneTr := range trs {
		key, val := ExtractDataFromTd(oneTr, prevKey, &isChild)
		prevKey = key

		switch strings.TrimSpace(key) {
		case "Branch/Office Name":
			branchDetail.BranchName = strings.TrimSpace(val)
		case "Branch Code":
			branchDetail.BranchCode = strings.TrimSpace(val)
		case "Address":
			branchDetail.Address = append(branchDetail.Address, strings.TrimSpace(val))
		case "City":
			branchDetail.City = strings.TrimSpace(val)
		case "PIN Code":
			branchDetail.Pincode = strings.TrimSpace(val)
		case "District":
			branchDetail.District = strings.TrimSpace(val)
		case "State":
			branchDetail.State = strings.TrimSpace(val)
		case "STD Code":
			branchDetail.StdCode = strings.TrimSpace(val)
		case "Phone Number":
			branchDetail.PhoneNumber = strings.TrimSpace(val)
		case "Email ID":
			branchDetail.EmailId = strings.TrimSpace(val)
		case "IFS Code":
			branchDetail.IFSCCode = strings.TrimSpace(val)
		case "Branch Circle":
			branchDetail.BranchCircle = strings.TrimSpace(val)
		default:
		}
	}
	color.Cyan("Time Taken: %s\n", time.Since(stTime))
	return &branchDetail, nil
}

func ExtractDataFromTd(parent soup.Root, parentKey string, isChild *bool) (string, string) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	tdsList := parent.FindAll("td")
	if len(tdsList) == 3 {
		*isChild = true
		value := ""
		value = tdsList[2].Text()
		return tdsList[0].Text(), value
	}
	if len(tdsList) == 1 && *isChild == true {
		return parentKey, tdsList[0].Text()
	}
	if len(tdsList) == 1 && *isChild == false {
		res := strings.Split(tdsList[0].Text(), ":")
		if len(res) == 2 {
			return res[0], res[1]
		} else {
			return res[0], ""
		}
	}
	return "", ""
}
