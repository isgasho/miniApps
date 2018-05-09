package main

import (
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	bf "gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
)

var (
	paraDepth        = 0
	fontSize         = 12.0
	pdf              *gofpdf.Fpdf
	liLevel          = 0
	tableEnabled     = false
	tableRowIndex    = -1
	tableColumnIndex = -1
)

type TableItem struct {
	Text      string
	Direction string
	Bold      bool
	Height    float64
	List      [][]byte
}

var ht float64
var orderedListLevelCount map[int]int
var orderedListLevelType map[int]bool
var tableData map[int][]*TableItem

func repeatString(inputString string, times int) string {
	var buffer bytes.Buffer
	for i := 0; i < times; i++ {
		buffer.WriteString(inputString)
	}
	return buffer.String()
}

func drawBullet(doc *gofpdf.Fpdf, x, y, size float64) {
	rs := size / 2
	r, g, b := 0, 0, 0
	doc.SetFillColor(r, g, b)
	doc.Circle(x-size*2, y-rs, rs, "F")
}

func main() {
	fileBytes, err := ioutil.ReadFile("./sample.md")
	if err != nil {
		panic(err)
	}
	orderedListLevelCount = make(map[int]int)
	orderedListLevelType = make(map[int]bool)
	tableData = make(map[int][]*TableItem)
	pdf = gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 210, Ht: 297},
		OrientationStr: "P",
	})
	pdf.SetMargins(20, 75, 20)
	pdf.SetFont("Arial", "", fontSize)
	ht = pdf.PointConvert(fontSize)
	pdf.AddPage()
	x := bf.New(bf.WithExtensions(bf.Tables))
	nodes := x.Parse(fileBytes)
	nodes.Walk(walker)
	err = pdf.OutputFileAndClose("./quotation1.pdf")
	if err != nil {
		fmt.Println(err)
	}

}

func walker(node *bf.Node, entering bool) bf.WalkStatus {
	fmt.Println(node, entering)
	switch node.Type {
	case bf.Strong:
		if entering {
			if tableEnabled {
				tableData[tableRowIndex][tableColumnIndex].Bold = true
				break
			}
			pdf.SetFont("Arial", "B", fontSize)
		} else {
			pdf.SetFont("Arial", "", fontSize)
		}
	case bf.Paragraph:
		if liLevel > 0 {
			if entering {
				pdf.Ln(ht / 2)
				break
			}
		}
		pdf.Ln(ht)
	case bf.HTMLSpan:
		if string(node.Literal) == "<br/>" {
			pdf.Ln(ht)
		}
	case bf.Document:
		break
	case bf.HorizontalRule:
		pdf.SetMargins(20, 30, 20)
		pdf.AddPage()
	case bf.Text:
		if liLevel > 0 {
			if orderedListLevelType[liLevel] {
				spacer := repeatString("     ", liLevel-1)
				if string(node.Literal) != "" {
					pdf.Write(ht, fmt.Sprintf("%s%d. %s", spacer, orderedListLevelCount[liLevel], string(node.Literal)))
				}
			} else {
				spacer := repeatString("       ", liLevel-1)
				pdf.Write(ht, fmt.Sprintf("%s", spacer))
				drawBullet(pdf, pdf.GetX()+ht/2, pdf.GetY()+ht/2+ht/4, ht/3)
				pdf.Write(ht, fmt.Sprintf(" %s", string(node.Literal)))
			}
			break
		}
		if tableEnabled {
			tableData[tableRowIndex][tableColumnIndex].Text = string(node.Literal)
			var alignment string
			switch node.Align {
			case bf.TableAlignmentCenter:
				alignment = "C"
			case bf.TableAlignmentLeft:
				alignment = "L"
			case bf.TableAlignmentRight:
				alignment = "R"
			default:
				alignment = "L"
			}
			tableData[tableRowIndex][tableColumnIndex].Direction = alignment
			break
		}
		pdf.Write(ht, string(node.Literal))
	case bf.List:
		if entering {
			liLevel++
			if node.ListFlags&bf.ListTypeOrdered != 0 {
				orderedListLevelType[liLevel] = true
			} else {
				orderedListLevelType[liLevel] = false
			}
			orderedListLevelCount[liLevel] = 0
		} else {
			orderedListLevelCount[liLevel] = 0
			liLevel--
		}
	case bf.Item:
		if entering {
			orderedListLevelCount[liLevel] = orderedListLevelCount[liLevel] + 1
		}
	case bf.Table:
		if entering {
			tableEnabled = true
		} else {
			tableEnabled = false
			renderTable()
		}
		tableRowIndex = -1
	case bf.TableRow:
		if entering {
			tableRowIndex++
		}
		tableColumnIndex = -1
	case bf.TableCell:
		if entering {
			tableColumnIndex++
			tableData[tableRowIndex] = append(tableData[tableRowIndex], &TableItem{})
		}
	default:
		break
	}
	return bf.GoToNext
}

func renderTable() {
	pdf.Ln(ht)
	colWd := 60.0
	marginH := 15.0
	lineHt := 5.5
	cellGap := 2.0
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
}
