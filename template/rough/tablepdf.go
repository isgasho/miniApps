package main

import (
	"github.com/jung-kurt/gofpdf"
	//"strings"
)

type TableItem struct {
	Text      string
	Direction string
	Bold      bool
	Height    float64
	List      [][]byte
}

var tableData map[int][]*TableItem

func populateDate() {
	tableData[0] = append(tableData[0], &TableItem{})
	tableData[0][0].Bold = true
	tableData[0][0].Direction = "R"
	tableData[0][0].Text = "Name"
	tableData[0] = append(tableData[0], &TableItem{})
	tableData[0][1].Bold = true
	tableData[0][1].Direction = "R"
	tableData[0][1].Text = "Addresss"

	tableData[1] = append(tableData[1], &TableItem{})
	tableData[1][0].Bold = true
	tableData[1][0].Direction = "R"
	tableData[1][0].Text = "Devarsh Shah"
	tableData[1] = append(tableData[1], &TableItem{})
	tableData[1][1].Bold = true
	tableData[1][1].Direction = "R"
	tableData[1][1].Text = "B-802 Retreat Tower \n Opp Shyamal Voltas \n Shyamal cross Road \n Satellite"

	tableData[2] = append(tableData[2], &TableItem{})
	tableData[2][0].Bold = true
	tableData[2][0].Direction = "R"
	tableData[2][0].Text = "Harsh Patel"
	tableData[2] = append(tableData[2], &TableItem{})
	tableData[2][1].Bold = true
	tableData[2][1].Direction = "R"
	tableData[2][1].Text = "A-202, Harihar Park \n Behind Haayat"

}

func main() {
	colWd := 60.0
	marginH := 15.0
	lineHt := 5.5
	cellGap := 2.0
	tableData = make(map[int][]*TableItem)
	populateDate()
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 14)
	pdf.SetMargins(marginH, 15, marginH)
	pdf.AddPage()
	pdf.SetTextColor(24, 24, 24)
	pdf.SetFillColor(255, 255, 255)
	y := pdf.GetY()

	for _, val := range tableData {
		maxHt := lineHt
		for _, val2 := range val {
			val2.List = pdf.SplitLines([]byte(val2.Text), colWd-cellGap-cellGap)
			val2.Height = float64(len(val2.List)) * lineHt
			if val2.Height > maxHt {
				maxHt = val2.Height
			}
		}
		x := marginH
		for _, val2 := range val {
			pdf.Rect(x, y, colWd, maxHt+cellGap+cellGap, "D")
			cellY := y + cellGap + (maxHt-val2.Height)/2
			for _, oneVal := range val2.List {
				pdf.SetXY(x+cellGap, cellY)
				pdf.CellFormat(colWd-cellGap-cellGap, lineHt, string(oneVal), "", 0, val2.Direction, false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap + cellGap
	}
	err := pdf.OutputFileAndClose("./out.pdf")
	if err != nil {
		panic(err)
	}
}

func renderTable() {

}
