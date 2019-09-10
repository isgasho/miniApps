package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func createTempDir(currentDir, pattern string) string {
	dir, err := ioutil.TempDir("", "template")
	if err != nil {
		log.Fatal(err)
	}
	fi, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, oneFile := range fi {
		if !oneFile.IsDir() {
			matched, err := filepath.Match(pattern, oneFile.Name())
			if err != nil {
				fmt.Println(err)
			} else if matched {
				inFile, err := os.Open(path.Join(currentDir, oneFile.Name()))
				if err != nil {
					log.Fatal(err)
				}
				defer inFile.Close()
				data, err := ioutil.ReadAll(inFile)
				if err != nil {
					log.Fatal(err)
				}
				length := len(data)
				if data[length-1] == 10 && data[length-2] == 59 {
					data = data[:length-2]
				}
				err = ioutil.WriteFile(filepath.Join(dir, oneFile.Name()), data, 0777)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return dir
}

func main() {
	tmpDir := createTempDir("./", "*.jsx")
	file, err := os.Create("../out/out.jsx")
	if err != nil {
		log.Fatalf("error creating the file %s", err)
	}
	defer file.Close()
	patter := "*.jsx"
	tmpl := template.Must(template.ParseGlob(filepath.Join(tmpDir, patter)))
	err = tmpl.Execute(file, nil)
	if err != nil {
		log.Fatalf("error generating template: %s", err)
	}
}
