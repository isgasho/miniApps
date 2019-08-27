package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func setTableProps(table *document.Table, attribs map[string]string) {
	for key, value := range attribs {
		switch strings.ToLower(key) {
		case "width":
			widthPer, err := strconv.ParseFloat(value, 32)
			if err != nil {
				fmt.Println("not a valid table widht")
			}
			table.Properties().SetWidthPercent(widthPer)
		case "border":
			borders := table.Properties().Borders()
			borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)
		}
	}
}

func setTableRowProps(row *document.Row, attribs map[string]string) {

}

func setRowCellProps(cell *document.Cell, attribs map[string]string) {
	for key, value := range attribs {
		switch strings.ToLower(key) {
		case "alignment":
			switch strings.ToLower(value) {
			case "left":
			case "right":
			case "center":
			}
		}
	}
}

/*
case "div":
				currentState = setupNewState(currentState, tname)
				para := doc.AddParagraph()
				currentState.currentPara = &para
case "p":
				currentState = setupNewState(currentState, tname)
				run := currentState.currentPara.AddRun()
				currentState.currentRun = &run
				if hasAttrib {
					attribs := getAttributes(tokenizer)
					currentState.setParaProps(attribs)
				}
case "table":
				tbl := doc.AddTable()
				tbl.Properties().SetAlignment(wml.ST_JcTableLeft)
				attribs := getAttributes(tokenizer)
				setTableProps(&tbl, attribs)
				p.currentTable = &tbl
			case "tr":
				row := p.currentTable.AddRow()
				p.currentRow = &row
			case "td":
				cell := p.currentRow.AddCell()
				p.currentCell = &cell
				para := p.currentCell.AddParagraph()
				p.currentPara = &para


				case "hyperlink":
				hlink := p.currentPara.AddHyperLink()
				if hasAttrib {
					attribs := getAttributes(tokenizer)
					link := attribs["href"]
					hlink.SetTarget(link)
					run := hlink.AddRun()
					clr := color.FromHex("#0563C1")
					run.Properties().SetColor(clr)
					run.Properties().SetUnderline(wml.ST_UnderlineSingle, clr)
					p.currentRun = &run
				}

				case "b":
				currentState = setupNewState(currentState)

				run := p.currentPara.AddRun()
				run.Properties().SetBold(true)
				p.currentRun = &run
			case "i":
				run := p.currentPara.AddRun()
				run.Properties().SetItalic(true)
				p.currentRun = &run
			case "u":
				run := p.currentPara.AddRun()
				run.Properties().SetUnderline(wml.ST_UnderlineSingle, color.Auto)
				p.currentRun = &run


				switch string(tn) {
			case "b":
				p.currentRun.Properties().SetBold(false)
			case "i":
				p.currentRun.Properties().SetItalic(false)
			case "u":
				p.currentRun.Properties().SetUnderline(wml.ST_UnderlineNone, color.Auto)
			}

*/
