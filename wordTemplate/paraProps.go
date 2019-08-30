package main

import (
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func (p *parserState) setParaProps(attribs map[string]string) {
	currentPara := p.currentPara
	if len(attribs) != 0 {
		for key, value := range attribs {
			switch key {
			case "keepNext", "keepLines", "pageBreakBefore",
				"suppressLineNumbers", "widowControl", "wordWrap",
				"overflowPunct", "topLinePunct", "autoSpaceDE", "autoSpaceDN",
				"rtl", "kinsoku", "adjustRightInd", "snapToGrid":
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
*/

/*
	PStyle *CT_String
	// Text Frame Properties
	FramePr *CT_FramePr
	// Numbering Definition Instance Reference
	NumPr *CT_NumPr
	// Paragraph Borders
	PBdr *CT_PBdr
	// Paragraph Shading
	Shd *CT_Shd
	// Set of Custom Tab Stops
	Tabs *CT_Tabs
	// Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *CT_OnOff
	// Spacing Between Lines and Above/Below Paragraph
	Spacing *CT_Spacing
	// Paragraph Indentation
	Ind *CT_Ind
	// Paragraph Alignment
	Jc *CT_Jc
	// Paragraph Text Flow Direction
	TextDirection *CT_TextDirection
	// Vertical Character Alignment on Line
	TextAlignment *CT_TextAlignment
	// Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	TextboxTightWrap *CT_TextboxTightWrap
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
