package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	initConstMap()
	tokenizer, err := readTemplateFile("./tmpl/demo.jsx")
	if err != nil {
		log.Fatal(err)
	}
	ancestorState := parserState{}
	ancestorState.prev = nil
	ancestorState.section = 0
	ancestorState.currentTag = 0
	ancestorState.currentTextStyle = &TextStyles{}
	parser(tokenizer, &ancestorState)
}

func readTemplateFile(filename string) (*html.Tokenizer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening the template file: %s", filename)
	}
	tokenizer := html.NewTokenizer(file)
	return tokenizer, nil
}
