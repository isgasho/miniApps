package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(filepath.Join(outDir, "doc.docx"), data, 0766)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done")

}

func fileOut(data io.Reader, filename string) (io.Reader, error) {
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)
	if _, err := io.Copy(w, data); err != nil {
		return nil, err
	}
	myData, err := ioutil.ReadAll(&buf1)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile(filename, myData, 0766)
	return &buf2, err
}
