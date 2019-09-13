package main

import (
	"strconv"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

type TableProps struct {
	tbl             *document.Table
	currentRowProps *TableRowProps
	currentRow      int
	currentCol      int
	spans           []*RowSpan
}

type RowSpan struct {
	startRow     int
	endRow       int
	columnNumber int
}
type TableRowProps struct {
	currentRow *document.Row
	alignment  wml.ST_Jc
	shading    *TableRowShadingProps
	margin     *TableRowCellsMargin
}

type TableRowShadingProps struct {
	style wml.ST_Shd
	color *color.Color
	fill  *color.Color
}

type TableRowCellsMargin struct {
	top    int
	bottom int
	left   int
	right  int
}

func (p *parserState) setTableCellProps(cell *document.Cell, attribs map[string]string) {
	if p.currentTable.currentRowProps.shading != nil {
		shd := p.currentTable.currentRowProps.shading
		cell.Properties().SetShading(shd.style, *shd.color, *shd.fill)
	}
	if p.currentTable.currentRowProps.margin != nil {
		mrg := p.currentTable.currentRowProps.margin
		cell.Properties().Margins().SetBottom(measurement.Distance(mrg.bottom))
		cell.Properties().Margins().SetTop(measurement.Distance(mrg.top))
		cell.Properties().Margins().SetLeft(measurement.Distance(mrg.left))
		cell.Properties().Margins().SetRight(measurement.Distance(mrg.right))
	}
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "width":
				num, err := strconv.ParseFloat(val, 64)
				if err == nil {
					cell.Properties().SetWidthPercent(num)
				}
			case "valign":
				num, err := strconv.Atoi(val)
				if err == nil {
					cell.Properties().SetVerticalAlignment(wml.ST_VerticalJc(num))
				}
			case "align":
				num, err := strconv.Atoi(val)
				if err == nil {
					p.currentPara.Properties().SetAlignment(wml.ST_Jc(num))
				}
			case "colspan":
				num, err := strconv.Atoi(val)
				if err == nil {
					cell.Properties().SetColumnSpan(num)
				}
			case "rowspan":
				span, err := strconv.Atoi(val)
				spanInst := RowSpan{}
				spanInst.startRow = p.currentTable.currentRow
				spanInst.columnNumber = p.currentTable.currentCol
				spanInst.endRow = p.currentTable.currentRow + span - 1
				p.currentTable.spans = append(p.currentTable.spans, &spanInst)
				if err == nil {
					cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)
				}
			}
		}
	}
	for _, one := range p.currentTable.spans {
		if one.columnNumber == p.currentTable.currentCol {
			if one.endRow >= p.currentTable.currentRow {
				cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
			}
		}
	}
}

func (p *parserState) setTableRowCellMargin(attribs map[string]string) {
	if p.currentTable.currentRowProps != nil {
		if len(attribs) != 0 {
			cellMarginInst := &TableRowCellsMargin{}
			p.currentTable.currentRowProps.margin = cellMarginInst
			for key, val := range attribs {
				switch key {
				case "top", "bottom", "left", "right", "all":
					num, err := strconv.Atoi(val)
					if err == nil {
						switch key {
						case "all":
							cellMarginInst.top = num
							cellMarginInst.bottom = num
							cellMarginInst.left = num
							cellMarginInst.right = num
						case "top":
							cellMarginInst.top = num
						case "bottom":
							cellMarginInst.bottom = num
						case "left":
							cellMarginInst.left = num
						case "right":
							cellMarginInst.right = num
						}
					}
				}
			}
		}
	}
}

func (p *parserState) setTableRowShading(attribs map[string]string) {
	if p.currentTable.currentRowProps != nil {
		rowShdInst := TableRowShadingProps{}
		p.currentTable.currentRowProps.shading = &rowShdInst
		rowShdInst.style = wml.ST_ShdSolid
		rowShdInst.color = &color.LightBlue
		rowShdInst.fill = &color.Black
		if len(attribs) != 0 {
			for key, val := range attribs {
				switch key {
				case "style":
					num, err := strconv.Atoi(val)
					if err == nil {
						rowShdInst.style = wml.ST_Shd(num)
					}
				case "color", "fill":
					clr := color.FromHex(val)
					switch key {
					case "color":
						rowShdInst.color = &clr
					case "fill":
						rowShdInst.fill = &clr
					}
				}
			}
		}
	}
}

func (p *parserState) setTableRowProps(attribs map[string]string) {
	if p.currentTable.currentRowProps != nil {
		rowProps := p.currentTable.currentRowProps
		rowProps.alignment = wml.ST_JcLeft
		if len(attribs) != 0 {
			rowInst := rowProps.currentRow
			for key, val := range attribs {
				switch key {
				case "height":
					strs := strings.Split(val, ",")
					if len(strs) == 2 {
						heightStr := strs[0]
						styleStr := strs[1]
						height, err1 := strconv.Atoi(heightStr)
						style, err2 := strconv.Atoi(styleStr)
						if err1 != nil && err2 != nil {
							rowInst.Properties().SetHeight(measurement.Distance(height), wml.ST_HeightRule(style))
						}
					}
				}
			}
		}
	}
}

func (p *parserState) setTableBorder(attribs map[string]string, direction SelfTags) {
	if p.currentTable.tbl.Properties().X().TblBorders != nil {
		tblBrd := p.currentTable.tbl.Properties().X().TblBorders
		brd := wml.NewCT_Border()
		clr, err := wml.ParseUnionST_HexColor("#000000")
		if err == nil {
			brd.ColorAttr = &clr
		}
		brd.SzAttr = unioffice.Uint64(uint64((measurement.Point * 1) / measurement.Point * 8))
		brd.ValAttr = wml.ST_BorderSingle
		if len(attribs) != 0 {
			for key, val := range attribs {
				switch key {
				case "color":
					clr, err := wml.ParseUnionST_HexColor(val)
					if err == nil {
						brd.ColorAttr = &clr
					}
				case "size":
					num, err := strconv.Atoi(val)
					if err == nil {
						brd.SzAttr = unioffice.Uint64(uint64((measurement.Point * num) / measurement.Point * 8))
					}
				case "style":
					num, err := strconv.Atoi(val)
					if err == nil {
						brd.ValAttr = wml.ST_Border(num)
					}
				}
			}
		}
		switch direction {
		case BorderAll:
			tblBrd.Top = brd
			tblBrd.Bottom = brd
			tblBrd.Left = brd
			tblBrd.Right = brd
			tblBrd.InsideH = brd
			tblBrd.InsideV = brd
		case BorderBottom:
			tblBrd.Bottom = brd
		case BorderTop:
			tblBrd.Top = brd
		case BorderLeft:
			tblBrd.Left = brd
		case BorderRight:
			tblBrd.Right = brd
		case BorderInsideHorizontal:
			tblBrd.InsideH = brd
		case BorderInsideVertical:
			tblBrd.InsideV = brd
		}
	}
}

func (p *parserState) setTableProps(attribs map[string]string) {
	currTbl := p.currentTable.tbl
	currTbl.Properties().SetLayout(wml.ST_TblLayoutTypeFixed)
	currTbl.Properties().SetWidth(90)
	currTbl.Properties().SetAlignment(wml.ST_JcTableLeft)
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "align", "layout":
				num, err := strconv.Atoi(val)
				if err == nil {
					switch key {
					case "align":
						currTbl.Properties().SetAlignment(wml.ST_JcTable(num))
					case "layout":
						currTbl.Properties().SetLayout(wml.ST_TblLayoutType(num))
					}
				}
			case "width", "cellspacing":
				num, err := strconv.ParseFloat(val, 64)
				if err == nil {
					switch key {
					case "width":
						currTbl.Properties().SetWidthPercent(num)
					case "cellspacing":
						currTbl.Properties().SetCellSpacingPercent(num)
					}
				}
			case "indent":
				width := wml.NewCT_TblWidth()
				currTbl.Properties().X().TblInd = width
				w, err := wml.ParseUnionST_MeasurementOrPercent(val)
				if err == nil {
					width.WAttr = &w
					width.TypeAttr = wml.ST_TblWidthDxa
				}
			}
		}
	}
}
