package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {

	file, err := os.Create("useragent.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	client := http.Client{}
	resp, err := client.Get("https://gist.githubusercontent.com/enginnr/ed572cf5c324ad04ff2e/raw/805b881fbe564531699a5c60b3e155110a124daf/useragentswitcher.xml")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	doc := soup.HTMLParse(string(respBytes))
	uagents := doc.FindAll("useragent")

	for _, oneUa := range uagents {
		x := oneUa.Attrs()
		bufferedWriter.WriteString(x["useragent"] + "\n")
	}
	err = bufferedWriter.Flush()
	if err != nil {
		panic(err)
	}
}
