package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devarsh/miniApps/wordTemplate/minify"
	"github.com/devarsh/miniApps/wordTemplate/parser"
	"github.com/devarsh/miniApps/wordTemplate/templategen"
)

func main() {
	outDir := "./out"
	err := os.MkdirAll(outDir, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err := templategen.GenerateTemplate("./tmpl", "*.tmpl", "00-main.tmpl", getTemplateFns(), initData())
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err = fileOut(reader, filepath.Join(outDir, "out.jsx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err = minify.MinifyCustomXml(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err = fileOut(reader, filepath.Join(outDir, "out.min.jsx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err = parser.ParseFile(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err = fileOut(reader, filepath.Join(outDir, "out.docx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	Unzip(reader, filepath.Join(outDir, "outUnzip"))
	fmt.Println("Done")
}
