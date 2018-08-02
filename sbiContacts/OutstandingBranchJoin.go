package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"

	"bytes"

	"github.com/devarsh/miniApps/sbiContacts/assets"
	"github.com/fatih/color"
)

func main() {
	inputFile := flag.String("f", "./Bills.csv", "input file name")
	skipHeader := flag.Bool("h", true, "Skip header Line")
	flag.Parse()
	db, err := assets.Asset("../_assets/sbiAllBranchDtl.csv")
	if err != nil {
		panic(err)
	}
	readCsv(*inputFile, db, *skipHeader)
}

func mapFile(data []byte) map[string][]string {
	mapper := make(map[string][]string)
	reader := bytes.NewReader(data)
	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, oneRecord := range records {
		mapper[oneRecord[0]] = oneRecord
	}
	return mapper
}

func readCsv(filePath string, database []byte, skipHeader bool) error {
	records := make([][]string, 0)
	regex := regexp.MustCompile(`\d+`)
	color.Cyan("Reading csv file: %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	r := csv.NewReader(bufio.NewReader(file))
	mapper := mapFile(database)
	oneRecord := []string{"Date", "RefNo", "Party_Name", "Branch_Code", "Pending_Amt", "BranchCode", "BranchName", "City", "Pincode", "District", "State", "StdCode", "PhoneNumber", "EmailId", "IFSC_Code", "BranchCircle", "Address-1", "Address-2", "Address-3"}
	records = append(records, oneRecord)
	if skipHeader {
		_, err := r.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		brCd := fmt.Sprintf("%05s", regex.Find([]byte(record[2])))
		if val, ok := mapper[brCd]; ok {
			oneRecord := []string{record[0], record[1], record[2], brCd, record[3], val[0], val[1], val[2], val[3], val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11], val[12], val[13]}
			records = append(records, oneRecord)
		} else {
			oneRecord := []string{record[0], record[1], record[2], "", record[3], "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
			records = append(records, oneRecord)
		}
	}
	err = writeCSV("./", "VRPLOutstandingContacts", &records)
	if err != nil {
		return err
	}
	return nil
}

func writeCSV(outFilePath string, name string, data *[][]string) error {
	filePath := path.Join(outFilePath, fmt.Sprintf("%s.csv", name))
	color.Cyan("Writing data to a csv file: %s", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	records := *data
	err = w.WriteAll(records)
	if err != nil {
		panic(err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	color.Cyan("Done creating file %s", filePath)
	return nil
}
