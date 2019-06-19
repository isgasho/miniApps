package mdToPdf

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/devarsh/miniApps/template/assets"
	"github.com/jung-kurt/gofpdf"
	bf "github.com/russross/blackfriday/v2"
)

type TableItem struct {
	Text      string
	Direction string
	Bold      bool
	Height    float64
	List      [][]byte
	Render    bool
}

type MdtoPdf struct {
	ht                    float64
	orderedListLevelCount map[int]int
	orderedListLevelType  map[int]bool
	tableData             map[int][]*TableItem
	paraDepth             int
	fontSize              float64
	fontStyle             string
	pdf                   *gofpdf.Fpdf
	liLevel               int
	tableEnabled          bool
	tableRowIndex         int
	tableColumnIndex      int
}

func NewMdtoPdf(fontSize float64, fontStyle string) *MdtoPdf {

	newInst := &MdtoPdf{}
	newInst.orderedListLevelCount = make(map[int]int)
	newInst.orderedListLevelType = make(map[int]bool)
	newInst.tableData = make(map[int][]*TableItem)
	newInst.fontSize = fontSize
	newInst.fontStyle = fontStyle
	return newInst
}

func (md *MdtoPdf) NewPdfWithHeader(fileBytes []byte, outfileName string, footer string) error {
	md.resetInternals()
	md.pdf = gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 210, Ht: 297},
		OrientationStr: "P",
	})
	md.ht = md.pdf.PointConvert(md.fontSize)
	md.pdf.SetFont(md.fontStyle, "", md.fontSize)
	md.pdf.SetFooterFunc(func() {
		md.pdf.SetY(-15)
		md.pdf.SetFont(md.fontStyle, "I", md.fontSize)
		md.pdf.CellFormat(0, md.ht, footer, "", 0, "C", false, 0, "")
	})
	md.pdf.AddPage()
	renderer := bf.New(bf.WithExtensions(bf.Tables))
	nodes := renderer.Parse(fileBytes)
	nodes.Walk(md.customWalker(true))
	err := md.pdf.OutputFileAndClose(outfileName + ".pdf")
	if err != nil {
		return err
	}
	return nil
}

func (md *MdtoPdf) NewPdf(fileBytes []byte, outfileName string, footer string) error {
	md.resetInternals()
	md.pdf = gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 210, Ht: 297},
		OrientationStr: "P",
	})
	md.ht = md.pdf.PointConvert(md.fontSize)
	md.pdf.SetMargins(20, 75, 20)
	md.pdf.SetFont(md.fontStyle, "", md.fontSize)
	md.pdf.SetFooterFunc(func() {
		md.pdf.SetY(-15)
		md.pdf.SetFont(md.fontStyle, "I", md.fontSize)
		md.pdf.CellFormat(0, md.ht, footer, "", 0, "C", false, 0, "")
	})
	md.pdf.AddPage()

	renderer := bf.New(bf.WithExtensions(bf.Tables))
	nodes := renderer.Parse(fileBytes)
	nodes.Walk(md.customWalker(false))
	if md.pdf.Err() {
		return md.pdf.Error()
	}
	err := md.pdf.OutputFileAndClose(outfileName + ".pdf")
	if err != nil {
		return err
	}
	return nil
}

func (md *MdtoPdf) resetInternals() {
	for key, _ := range md.orderedListLevelCount {
		delete(md.orderedListLevelCount, key)
	}
	for key, _ := range md.orderedListLevelType {
		delete(md.orderedListLevelType, key)
	}
	for key, _ := range md.tableData {
		delete(md.tableData, key)
	}
}

func (md *MdtoPdf) customWalker(printHeader bool) bf.NodeVisitor {
	return func(node *bf.Node, entering bool) bf.WalkStatus {
		switch node.Type {
		case bf.Strong:
			if entering {
				if md.tableEnabled {
					md.tableData[md.tableRowIndex][md.tableColumnIndex].Bold = true
					break
				}
				md.pdf.SetFont(md.fontStyle, "B", md.fontSize)
			} else {
				md.pdf.SetFont(md.fontStyle, "", md.fontSize)
			}
		case bf.Paragraph:
			if md.liLevel > 0 {
				if entering {
					md.pdf.Ln(md.ht / 2)
					break
				}
			}
			md.pdf.Ln(md.ht)
		case bf.HTMLSpan:
			if string(node.Literal) == "<br/>" {
				md.pdf.Ln(md.ht)
			}
		case bf.Document:
			break
		case bf.HorizontalRule:
			if printHeader != true {
				md.pdf.SetMargins(20, 40, 20)
			}
			md.pdf.AddPage()
		case bf.Text:
			if md.liLevel > 0 {
				if md.orderedListLevelType[md.liLevel] {
					spacer := repeatString("     ", md.liLevel-1)
					if string(node.Literal) != "" {
						md.pdf.Write(md.ht, fmt.Sprintf("%s%d. %s", spacer, md.orderedListLevelCount[md.liLevel], string(node.Literal)))
					}
				} else {
					spacer := repeatString("       ", md.liLevel-1)
					md.pdf.Write(md.ht, fmt.Sprintf("%s", spacer))
					md.drawBullet(md.pdf.GetX()+md.ht/2, md.pdf.GetY()+md.ht/2+md.ht/4, md.ht/3)
					md.pdf.Write(md.ht, fmt.Sprintf(" %s", string(node.Literal)))
				}
				break
			}
			if md.tableEnabled {
				md.tableData[md.tableRowIndex][md.tableColumnIndex].Text = string(node.Literal)
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
				md.tableData[md.tableRowIndex][md.tableColumnIndex].Direction = alignment
				break
			}
			md.pdf.Write(md.ht, string(node.Literal))
		case bf.List:
			if entering {
				md.liLevel++
				if node.ListFlags&bf.ListTypeOrdered != 0 {
					md.orderedListLevelType[md.liLevel] = true
				} else {
					md.orderedListLevelType[md.liLevel] = false
				}
				md.orderedListLevelCount[md.liLevel] = 0
			} else {
				md.orderedListLevelCount[md.liLevel] = 0
				md.liLevel--
			}
		case bf.Item:
			if entering {
				md.orderedListLevelCount[md.liLevel] = md.orderedListLevelCount[md.liLevel] + 1
			}
		case bf.Table:
			if entering {
				md.tableEnabled = true
			} else {
				md.tableEnabled = false
				md.renderTable()
			}
			md.tableRowIndex = -1
		case bf.TableRow:
			if entering {
				md.tableRowIndex++
			}
			md.tableColumnIndex = -1
		case bf.TableCell:
			if entering {
				md.tableColumnIndex++
				md.tableData[md.tableRowIndex] = append(md.tableData[md.tableRowIndex], &TableItem{})
			}
		case bf.Link:
			if printHeader {
				if entering {
					var path, params string
					data := strings.Split(strings.TrimSpace(string(node.LinkData.Destination)), " ")
					if len(data) == 2 {
						path = strings.TrimSpace(data[0])
						params = strings.TrimSpace(data[1])
					}
					paramsArr := strings.Split(params, ",")
					if len(paramsArr) == 3 {
						x, err := strconv.ParseFloat(paramsArr[0], 64)
						if err != nil {
							fmt.Println("Invalid Integer", paramsArr[0])
						}
						w, err := strconv.ParseFloat(paramsArr[1], 64)
						if err != nil {
							fmt.Println("Invalid Integer", paramsArr[1])
						}
						h, err := strconv.ParseFloat(paramsArr[2], 64)
						if err != nil {
							fmt.Println("Invalid Integer", paramsArr[2])
						}
						imgBuf, err := assets.Asset(path)
						if err != nil {
							panic(err)
						}
						reader := bytes.NewReader(imgBuf)
						md.pdf.RegisterImageReader(path, "PNG", reader)
						md.pdf.ImageOptions(path, x, md.pdf.GetY(), w, h, true, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")
					}
				}
			}
		default:
			break
		}
		return bf.GoToNext
	}
}

func (md *MdtoPdf) renderTable() {
	md.pdf.Ln(md.ht)
	marginH := 20.0
	lineHt := 5.5
	cellGap := 2.0

	ColWdArray := make([]float64, 0)
	AlignArray := make([]string, 0)
	if allWidths, ok := md.tableData[1]; ok {
		for _, oneWidth := range allWidths {
			strs := strings.Split(oneWidth.Text, ",")
			if len(strs) != 2 {
				panic("Need to Pass tabe width and alignment" + oneWidth.Text)
			}
			f, err := strconv.ParseFloat(strs[0], 64)
			if err != nil {
				fmt.Println(err)
				panic("size not defined cannot move ahead with tables:")
			}
			ColWdArray = append(ColWdArray, f)
			AlignArray = append(AlignArray, strs[1])
		}
		delete(md.tableData, 1) //because 2nd row ie. 1st index is width size
		y := md.pdf.GetY()
		sortedKeys := md.getSortedKeys()
		for _, itr := range sortedKeys {
			val := md.tableData[itr]
			if itr == 0 {
				md.pdf.SetFont(md.fontStyle, "B", md.fontSize)
			} else {
				md.pdf.SetFont(md.fontStyle, "", md.fontSize)
			}
			maxHt := lineHt
			for key, val2 := range val {
				newStr := strings.Split(val2.Text, `\n`)
				str2 := strings.Join(newStr, "\n")
				val2.List = md.pdf.SplitLines([]byte(str2), ColWdArray[key]-cellGap-cellGap)
				val2.Height = float64(len(val2.List)) * lineHt
				if val2.Height > maxHt {
					maxHt = val2.Height
				}
			}
			x := marginH
			for key, val2 := range val {
				md.pdf.Rect(x, y, ColWdArray[key], maxHt+cellGap+cellGap, "D")
				cellY := y + cellGap //+ (maxHt-val2.Height)/2 //if you need text vertically center
				for _, oneVal := range val2.List {
					md.pdf.SetXY(x+cellGap, cellY)
					if key == 0 {
						md.pdf.CellFormat(ColWdArray[key]-cellGap-cellGap, lineHt, string(oneVal), "", 0, "L", false, 0, "")
					} else {
						md.pdf.CellFormat(ColWdArray[key]-cellGap-cellGap, lineHt, string(oneVal), "", 0, AlignArray[key], false, 0, "")
					}
					cellY += lineHt
				}
				x += ColWdArray[key]
			}
			y += maxHt + cellGap + cellGap
		}
		md.pdf.Ln(md.ht)
		for key, _ := range md.tableData {
			delete(md.tableData, key)
		}
	} else {
		panic("Cannot render table no column size found on Row No 2 index 1 found")
	}
}

func (md *MdtoPdf) drawBullet(x, y, size float64) {
	rs := size / 2
	r, g, b := 0, 0, 0
	md.pdf.SetFillColor(r, g, b)
	md.pdf.Circle(x-size*2, y-rs, rs, "F")
}

func (md *MdtoPdf) getSortedKeys() []int {
	keys := make([]int, 0)
	for key, _ := range md.tableData {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func repeatString(inputString string, times int) string {
	var buffer bytes.Buffer
	for i := 0; i < times; i++ {
		buffer.WriteString(inputString)
	}
	return buffer.String()
}
