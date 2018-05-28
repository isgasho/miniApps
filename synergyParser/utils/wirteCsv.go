package utils

import (
	"encoding/csv"
	"os"
)

func WriteCsv(records [][]string, filePath string) error {
	fs, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fs.Close()
	w := csv.NewWriter(fs)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
