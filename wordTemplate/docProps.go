package main

import (
	"strconv"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func setDocBorderProps(doc *document.Document, attribs map[string]string) {
	if len(attribs) != 0 {
		brd := doc.BodySection().X().PgBorders
		if brd != nil {
			for key, val := range attribs {
				switch key {
				case "zorder", "display", "offset":
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "zorder":
							brd.ZOrderAttr = wml.ST_PageBorderZOrder(num)
						case "display":
							brd.DisplayAttr = wml.ST_PageBorderDisplay(num)
						case "offset":
							brd.OffsetFromAttr = wml.ST_PageBorderOffset(num)
						}
					}
				}
			}
		}
	}
}

func setLeftRightDocBorder(doc *document.Document, attribs map[string]string, direction SelfTags) {
	if len(attribs) != 0 {
		brd := doc.BodySection().X().PgBorders
		if brd != nil {
			brdLeftRight := wml.NewCT_PageBorder()
			if direction == BorderLeft {
				brd.Left = brdLeftRight
			} else if direction == BorderRight {
				brd.Right = brdLeftRight
			}
			for key, val := range attribs {
				switch key {
				case "style", "size", "space":
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "style":
							brdLeftRight.ValAttr = wml.ST_Border(num)
						case "size":
							brdLeftRight.SzAttr = unioffice.Uint64(uint64(num))
						case "space":
							brdLeftRight.SpaceAttr = unioffice.Uint64(uint64(num))
						}
					}
				case "color":
					clr, err := wml.ParseUnionST_HexColor(val)
					if err == nil {
						brdLeftRight.ColorAttr = &clr
					}
				case "frame", "shadow":
					onoff, err := wml.ParseUnionST_OnOff(val)
					if err == nil {
						switch key {
						case "frame":
							brdLeftRight.FrameAttr = &onoff
						case "shadow":
							brdLeftRight.ShadowAttr = &onoff
						}
					}
				}
			}
		}
	}
}

func setBottomDocBorder(doc *document.Document, attribs map[string]string) {
	if len(attribs) != 0 {
		brd := doc.BodySection().X().PgBorders
		if brd != nil {
			brdBottom := wml.NewCT_BottomPageBorder()
			brd.Bottom = brdBottom
			for key, val := range attribs {
				switch key {
				case "style", "size", "space":
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "style":
							brdBottom.ValAttr = wml.ST_Border(num)
						case "size":
							brdBottom.SzAttr = unioffice.Uint64(uint64(num))
						case "space":
							brdBottom.SpaceAttr = unioffice.Uint64(uint64(num))
						}
					}
				case "color":
					clr, err := wml.ParseUnionST_HexColor(val)
					if err == nil {
						brdBottom.ColorAttr = &clr
					}
				case "frame", "shadow":
					onoff, err := wml.ParseUnionST_OnOff(val)
					if err == nil {
						switch key {
						case "frame":
							brdBottom.FrameAttr = &onoff
						case "shadow":
							brdBottom.ShadowAttr = &onoff
						}
					}
				}
			}
		}
	}
}

func setTopDocBorder(doc *document.Document, attribs map[string]string) {
	if len(attribs) != 0 {
		brd := doc.BodySection().X().PgBorders
		if brd != nil {
			brdTop := wml.NewCT_TopPageBorder()
			brd.Top = brdTop
			for key, val := range attribs {
				switch key {
				case "style", "size", "space":
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "style":
							brdTop.ValAttr = wml.ST_Border(num)
						case "size":
							brdTop.SzAttr = unioffice.Uint64(uint64(num))
						case "space":
							brdTop.SpaceAttr = unioffice.Uint64(uint64(num))
						}
					}
				case "color":
					clr, err := wml.ParseUnionST_HexColor(val)
					if err == nil {
						brdTop.ColorAttr = &clr
					}
				case "frame", "shadow":
					onoff, err := wml.ParseUnionST_OnOff(val)
					if err == nil {
						switch key {
						case "frame":
							brdTop.FrameAttr = &onoff
						case "shadow":
							brdTop.ShadowAttr = &onoff
						}
					}
				}
			}
		}

	}
}

func setDocBackground(doc *document.Document, attribs map[string]string) {
	if len(attribs) != 0 {
		var background *wml.CT_Background
		if doc.X().Background == nil {
			background = wml.NewCT_Background()
			doc.X().Background = background
		} else {
			background = doc.X().Background
		}
		for key, val := range attribs {
			switch key {
			case "color":
				clr, err := wml.ParseUnionST_HexColor(val)
				if err == nil {
					background.ColorAttr = &clr
				}
			case "theme":
				num, err := strconv.Atoi(val)
				if err == nil {
					background.ThemeColorAttr = wml.ST_ThemeColor(num)
				}
			case "tint":
				background.ThemeTintAttr = &val
			case "shade":
				background.ThemeShadeAttr = &val

			}
		}
	}
}

func setDocPageMargin(doc *document.Document, attribs map[string]string) {
	if len(attribs) != 0 {
		pgMargin := wml.NewCT_PageMar()
		doc.BodySection().X().PgMar = pgMargin
		for key, val := range attribs {
			switch key {
			case "top", "bottom":
				mes, err := wml.ParseUnionST_SignedTwipsMeasure(val)
				if err == nil {
					switch key {
					case "top":
						pgMargin.TopAttr = mes
					case "bottom":
						pgMargin.BottomAttr = mes
					}
				}
			case "left", "right", "header", "footer":
				mes, err := wml.ParseUnionST_TwipsMeasure(val)
				if err == nil {
					switch key {
					case "left":
						pgMargin.LeftAttr = mes
					case "right":
						pgMargin.RightAttr = mes
					case "header":
						pgMargin.HeaderAttr = mes
					case "footer":
						pgMargin.FooterAttr = mes
					case "gutter":
						pgMargin.GutterAttr = mes
					}
				}
			}
		}
	}
}

func setDocPageSize(doc *document.Document, attribs map[string]string) {
	if len(attribs) != 0 {
		pgSize := wml.NewCT_PageSz()
		doc.BodySection().X().PgSz = pgSize
		for key, val := range attribs {
			switch key {
			case "width", "height":
				mes, err := wml.ParseUnionST_TwipsMeasure(val)
				if err == nil {
					switch key {
					case "width":
						pgSize.WAttr = &mes
					case "height":
						pgSize.HAttr = &mes
					}
				}
			case "orientation":
				num, err := strconv.Atoi(val)
				if err == nil {
					pgSize.OrientAttr = wml.ST_PageOrientation(num)
				}
			}
		}
	}
}
