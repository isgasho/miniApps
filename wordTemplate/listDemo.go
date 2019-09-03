package main

import (
	"fmt"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main2() {
	doc := document.New()
	nd := doc.Numbering.AddDefinition()
	const indentStart = 0
	const indentDelta = 20
	const hangingIndent = 20
	for i := 0; i < 9; i++ {
		lvl := nd.AddLevel()
		lvl.SetAlignment(wml.ST_JcLeft)
		lvl.SetText(fmt.Sprintf("%%%d.", i+1))
		leftindent := int64(i*indentDelta + indentStart)
		lvl.Properties().SetLeftIndent(measurement.Distance(leftindent))
		lvl.Properties().SetHangingIndent(measurement.Distance(hangingIndent))
		switch i {
		case 0, 3, 6:
			lvl.SetFormat(wml.ST_NumberFormatDecimal)
		case 1, 4, 7:
			lvl.SetFormat(wml.ST_NumberFormatLowerLetter)
		case 2, 5, 8:
			lvl.SetFormat(wml.ST_NumberFormatLowerRoman)
		}
	}
	for i := 0; i < 5; i++ {
		para := doc.AddParagraph()
		para.SetNumberingDefinition(nd)
		para.SetNumberingLevel(0)
		para.Properties().SetSpacing(measurement.Distance(2), measurement.Distance(2))
		para.SetStyle("Heading1")
		_ = para.AddRun()
		run := para.AddRun()
		run.AddText(lorem)
	}

	doc.SaveToFile("demo.docx")
}
