package main

import (
	"strconv"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func (p *parserState) applyParaAlignment(attribs map[string]string) {
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "style":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					p.currentPara.Properties().SetAlignment(wml.ST_Jc(num))
				}
			}
		}
	}
}

func (p *parserState) applyParaSpacing(attribs map[string]string) {
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "after", "before":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					switch key {
					case "after":
						p.currentPara.Properties().Spacing().SetAfter(measurement.Distance(num))
					case "before":
						p.currentPara.Properties().Spacing().SetBefore(measurement.Distance(num))
					}
				}
			case "autoafter", "autobefore":
				if val == "true" || val == "1" || val == "on" {
					switch key {
					case "autoAfter":
						p.currentPara.Properties().Spacing().SetAfterAuto(true)

					case "autoBefore":
						p.currentPara.Properties().Spacing().SetBeforeAuto(true)
					}

				}
			case "linespacing":
				vals := strings.Split(val, ",")
				if len(vals) == 2 {
					lheight := vals[0]
					lstyle := vals[1]
					lheightNum, err1 := strconv.ParseInt(lheight, 10, 64)
					lstyleNum, err2 := strconv.ParseInt(lstyle, 10, 64)
					if err1 == nil && err2 == nil {
						p.currentPara.Properties().Spacing().SetLineSpacing(measurement.Distance(lheightNum), wml.ST_LineSpacingRule(lstyleNum))
					}
				}
			}
		}
	}
}

func (p *parserState) applyParaIndent(attribs map[string]string) {
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "start", "end", "hang", "first":
				{
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "start":
							p.currentPara.Properties().SetStartIndent(measurement.Distance(num))
						case "end":
							p.currentPara.Properties().SetEndIndent(measurement.Distance(num))
						case "hang":
							p.currentPara.Properties().SetHangingIndent(measurement.Distance(num))
						case "first":
							p.currentPara.Properties().SetFirstLineIndent(measurement.Distance(num))
						}
					}
				}
			}
		}
	}
}

func (p *parserState) applyParaFrame(attribs map[string]string) {
	if len(attribs) != 0 {
		exProps := p.currentPara.Properties().X()
		frame := wml.NewCT_FramePr()
		exProps.FramePr = frame
		for key, val := range attribs {
			switch key {
			case "dropCap", "lines", "wrap",
				"hAnchor", "vAnchor", "xAlign", "yAlign",
				"hRule":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					switch key {
					case "dropCap":
						frame.DropCapAttr = wml.ST_DropCap(num)
					case "lines":
						frame.LinesAttr = &num
					case "wrap":
						frame.WrapAttr = wml.ST_Wrap(num)
					case "hAnchor":
						frame.HAnchorAttr = wml.ST_HAnchor(num)
					case "vAnchor":
						frame.VAnchorAttr = wml.ST_VAnchor(num)
					case "xAlign":
						frame.XAlignAttr = sharedTypes.ST_XAlign(num)
					case "yAlign":
						frame.YAlignAttr = sharedTypes.ST_YAlign(num)
					case "hRule":
						frame.HRuleAttr = wml.ST_HeightRule(num)
					}
				}
			case "height", "width", "vpad", "hpad":
				mes, err := wml.ParseUnionST_TwipsMeasure(val)
				if err != nil {
					switch key {
					case "height":
						frame.HAttr = &mes
					case "width":
						frame.WAttr = &mes
					case "hSpace":
						frame.HSpaceAttr = &mes
					case "vSpace":
						frame.VSpaceAttr = &mes
					}
				}
			case "x", "y":
				mes, err := wml.ParseUnionST_SignedTwipsMeasure(val)
				if err != nil {
					switch key {
					case "x":
						frame.XAttr = &mes
					case "y":
						frame.YAttr = &mes
					}
				}
			}
		}
	}
}

func (p *parserState) applyParaTextProps(attribs map[string]string) {
	if len(attribs) != 0 {
		for key, val := range attribs {

			switch key {
			case "align", "direction":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					exProps := p.currentPara.Properties().X()
					switch key {
					case "align":
						val := wml.NewCT_TextAlignment()
						val.ValAttr = wml.ST_TextAlignment(num)
						exProps.TextAlignment = val
					case "direction":
						val := wml.NewCT_TextDirection()
						val.ValAttr = wml.ST_TextDirection(num)
						exProps.TextDirection = val
					}
				}

			}
		}
	}
}

func (p *parserState) applyParaShading(attribs map[string]string) {
	if len(attribs) != 0 {
		exProps := p.currentPara.Properties().X()
		newShd := wml.NewCT_Shd()
		exProps.Shd = newShd
		for key, val := range attribs {
			switch key {
			case "style":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					newShd.ValAttr = wml.ST_Shd(num)
				}
			case "color", "fill":
				clr, err := wml.ParseUnionST_HexColor(val)
				if err == nil {
					switch key {
					case "color":
						newShd.ColorAttr = &clr
					case "fill":
						newShd.FillAttr = &clr
					}
				}
			}
		}
	}
}

func (p *parserState) applyParaTextBoxTightWrap(attribs map[string]string) {
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "style":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					wrap := wml.NewCT_TextboxTightWrap()
					wrap.ValAttr = wml.ST_TextboxTightWrap(num)
					p.currentPara.Properties().X().TextboxTightWrap = wrap
				}
			}
		}
	}
}

func (p *parserState) applyParaBorder(attribs map[string]string, direction SelfTags) {
	if len(attribs) != 0 {
		side := p.currentPara.Properties().X().PBdr
		if side != nil {
			currBorder := wml.NewCT_Border()
			for key, val := range attribs {
				switch key {
				case "style", "size", "space":
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "style":
							currBorder.ValAttr = wml.ST_Border(num)
						case "size":
							currBorder.SzAttr = unioffice.Uint64(uint64(num))
						case "space":
							currBorder.SpaceAttr = unioffice.Uint64(uint64(num))
						}
					}
				case "color":
					clr, err := wml.ParseUnionST_HexColor(val)
					if err == nil {
						currBorder.ColorAttr = &clr
					}
				case "frame", "shadow":
					onoff, err := wml.ParseUnionST_OnOff(val)
					if err == nil {
						switch key {
						case "frame":
							currBorder.FrameAttr = &onoff
						case "shadow":
							currBorder.ShadowAttr = &onoff
						}
					}
				}
			}
			switch direction {
			case BorderLeft:
				side.Left = currBorder
			case BorderRight:
				side.Right = currBorder
			case BorderTop:
				side.Top = currBorder
			case BorderBottom:
				side.Bottom = currBorder
			}
		}
	}
}

func (p *parserState) setParaProps(attribs map[string]string) {
	currentPara := p.currentPara
	if len(attribs) != 0 {
		for key, value := range attribs {
			switch key {
			case "keepNext", "keepLines", "pageBreakBefore",
				"suppressLineNumbers", "widowControl", "wordWrap",
				"overflowPunct", "topLinePunct", "autoSpaceDE", "autoSpaceDN",
				"rtl", "kinsoku", "adjustRightInd", "snapToGrid",
				"contextualSpacing", "mirrorIndents", "suppressOverlap",
				"suppressAutoHyphens":
				if value == "true" || value == "on" || value == "1" || value == "" {
					onOff := wml.NewCT_OnOff()
					exprop := currentPara.Properties().X()
					switch key {
					case "keepNext":
						exprop.KeepNext = onOff
					case "keepLines":
						exprop.KeepLines = onOff
					case "pageBreakBefore":
						exprop.PageBreakBefore = onOff
					case "suppressLineNumbers":
						exprop.SuppressLineNumbers = onOff
					case "widowControl":
						exprop.WidowControl = onOff
					case "wordWrap":
						exprop.WordWrap = onOff
					case "overflowPunct":
						exprop.OverflowPunct = onOff
					case "topLinePunct":
						exprop.TopLinePunct = onOff
					case "autoSpaceDE":
						exprop.AutoSpaceDE = onOff
					case "autoSpaceDN":
						exprop.AutoSpaceDN = onOff
					case "rtl":
						exprop.Bidi = onOff
					case "kinsoku":
						exprop.Kinsoku = onOff
					case "adjustRightInd":
						exprop.AdjustRightInd = onOff
					case "snapToGrid":
						exprop.SnapToGrid = onOff
					case "contextualSpacing":
						exprop.ContextualSpacing = onOff
					case "mirrorIndents":
						exprop.MirrorIndents = onOff
					case "suppressOverlap":
						exprop.SuppressOverlap = onOff
					case "suppressAutoHyphens":
						exprop.SuppressAutoHyphens = onOff
					}
				}
			}

		}
	}
}

/*
// Keep Paragraph With Next Paragraph
	KeepNext *CT_OnOff
	// Keep All Lines On One Page
	KeepLines *CT_OnOff
	// Start Paragraph on Next Page
	PageBreakBefore *CT_OnOff
	// Suppress Line Numbers for Paragraph
	SuppressLineNumbers *CT_OnOff
	// Allow First/Last Line to Display on a Separate Page
	WidowControl *CT_OnOff
	// Allow Line Breaking At Character Level
	WordWrap *CT_OnOff
	// Allow Punctuation to Extend Past Text Extents
	OverflowPunct *CT_OnOff
	// Compress Punctuation at Start of a Line
	TopLinePunct *CT_OnOff
	// Automatically Adjust Spacing of Latin and East Asian Text
	AutoSpaceDE *CT_OnOff
	// Automatically Adjust Spacing of East Asian Text and Numbers
	AutoSpaceDN *CT_OnOff
	// Right to Left Paragraph Layout
	Bidi *CT_OnOff //rtl
	// Use East Asian Typography Rules for First and Last Character per Line
	Kinsoku *CT_OnOff
	// Automatically Adjust Right Indent When Using Document Grid
	AdjustRightInd *CT_OnOff
	// Use Document Grid Settings for Inter-Line Paragraph Spacing
	SnapToGrid *CT_OnOff
	// Ignore Spacing Above and Below When Using Identical Styles
	ContextualSpacing *CT_OnOff
	// Use Left/Right Indents as Inside/Outside Indents
	MirrorIndents *CT_OnOff
	// Prevent Text Frames From Overlapping
	SuppressOverlap *CT_OnOff
	// Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *CT_OnOff

	// Paragraph Borders
	PBdr *CT_PBdr
	// Paragraph Shading
	Shd *CT_Shd

	// Spacing Between Lines and Above/Below Paragraph
	Spacing *CT_Spacing
	// Paragraph Alignment
	Jc *CT_Jc
	// Paragraph Text Flow Direction
	TextDirection *CT_TextDirection
	// Vertical Character Alignment on Line
	TextAlignment *CT_TextAlignment
	FramePr *CT_FramePr
	// Numbering Definition Instance Reference
	// Paragraph Indentation
	Ind *CT_Ind
	// Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	TextboxTightWrap *CT_TextboxTightWrap

*/

/*
	PStyle *CT_String
	// Text Frame Properties

	NumPr *CT_NumPr
	// Set of Custom Tab Stops
	Tabs *CT_Tabs


	// Associated Outline Level
	OutlineLvl *CT_DecimalNumber
	// Associated HTML div ID
	DivId *CT_DecimalNumber
	// Paragraph Conditional Formatting
	CnfStyle  *CT_Cnf
	RPr       *CT_ParaRPr
	SectPr    *CT_SectPr
	PPrChange *CT_PPrChange
*/
