package main

import (
	bf "gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"os"
)

func main() {
	fp, err := os.Open("./template.md")
	if err != nil {
		panic(err)
	}
	input, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	out := bf.Run(input, bf.WithExtensions(bf.Tables|bf.Footnotes|bf.Strikethrough|bf.FencedCode|bf.LaxHTMLBlocks|bf.HardLineBreak))
	of, err := os.Create("./output.html")
	_, err = of.Write(out)
	if err != nil {
		panic(err)
	}
}
