package parser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/net/html"
)

func initParser() *parserState {
	ancestorState := parserState{}
	ancestorState.prev = nil
	ancestorState.section = 0
	ancestorState.currentTag = 0
	ancestorState.currentTextStyle = &TextStyles{}
	return &ancestorState
}

func ParseFile(reader io.Reader) (io.Reader, error) {
	initConstMap()
	tokenizer := html.NewTokenizer(reader)
	ancestorState := initParser()
	reader, err := parse(tokenizer, ancestorState)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func ParseFile2(filename string) (io.Reader, error) {
	initConstMap()
	tokenizer, err := readTemplateFileAndTokenize(filename)
	if err != nil {
		return nil, err
	}
	ancestorState := initParser()
	reader, err := parse(tokenizer, ancestorState)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func ParseFile3(filename string, outPath string, outFileName string) error {
	initConstMap()
	tokenizer, err := readTemplateFileAndTokenize(filename)
	if err != nil {
		return err
	}
	ancestorState := initParser()
	reader, err := parse(tokenizer, ancestorState)
	if err != nil {
		return err
	}
	err = os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	ioutil.WriteFile(fmt.Sprintf("%s.docx", path.Join(outPath, outFileName)), data, 0644)
	return nil
}

func readTemplateFileAndTokenize(filename string) (*html.Tokenizer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening the template file: %s", filename)
	}
	tokenizer := html.NewTokenizer(file)
	return tokenizer, nil
}
