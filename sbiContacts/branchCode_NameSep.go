package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"path"
	"regexp"
)

func main() {
	readCsv("./Master.csv")
}

func readCsv(filePath string) error {
	records := make([][]string, 10)
	records1 := make([][]string, 10)
	regex := regexp.MustCompile(`\d+`)
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	r := csv.NewReader(bufio.NewReader(file))
	color.Cyan("Reading csv file: %s", filePath)
	oneRecord := []string{"BranchCode", "BranchName"}
	records = append(records, oneRecord)
	oneRecord = []string{"BranchCode"}
	records1 = append(records1, oneRecord)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		dt := fmt.Sprintf("%05s", regex.Find([]byte(record[0])))
		oneRecord := []string{dt, record[0]}
		records = append(records, oneRecord)
		oneRecord = []string{dt}
		records1 = append(records1, oneRecord)
	}
	err = writeCSV("./", "VRPLSBIBRCODE", &records)
	if err != nil {
		return err
	}
	err = writeCSV("./", "VRPLSBIBRCODE1", &records1)
	if err != nil {
		return err
	}
	return nil
}

func writeCSV(outFilePath string, name string, data *[][]string) error {
	filePath := path.Join(outFilePath, fmt.Sprintf("VebsbiBranchCodes_%s.csv", name))
	color.Cyan("Writing data to a csv file: %s", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	records := *data
	for _, record := range records {
		if len(record) == 0 {
			continue
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	color.Cyan("Done creating file %s", filePath)
	return nil
}
