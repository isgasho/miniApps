package main

import (
	"strconv"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

type parserState struct {
	currentPara  *document.Paragraph
	currentRun   *document.Run
	prev         *parserState
	section      Tags
	currentTag   Tags
	currentStyle *Styles
}

type Styles struct {
	flags         StyleTags
	underline     *UnderlineProps
	emphasisStyle wml.ST_Em
	font          *FontProps
}

type UnderlineProps struct {
	style wml.ST_Underline
	color color.Color
}

type FontProps struct {
	family string
	size   float64
	kern   float64
	color  color.Color
}

func NewParserState(currentState *parserState, tagName Tags) *parserState {
	newState := &parserState{}
	newState.prev = currentState
	newState.currentPara = currentState.currentPara
	newState.currentRun = currentState.currentRun
	newState.section = currentState.section
	newState.currentTag = tagName
	newState.currentStyle = currentState.currentStyle
	return newState
}

func (p *parserState) setHeaderFooterParagraphPropsPstyle(value string) {
	paraProps := p.currentPara.X()
	paraProps.PPr = wml.NewCT_PPr()
	paraProps.PPr.PStyle = wml.NewCT_String()
	paraProps.PPr.PStyle.ValAttr = value
}

func (p *parserState) setAlignmentTab(relativeTo wml.ST_PTabRelativeTo, leader wml.ST_PTabLeader, alignment wml.ST_PTabAlignment) {
	ic := wml.NewEG_RunInnerContent()
	ic.Ptab = wml.NewCT_PTab()
	ic.Ptab.RelativeToAttr = relativeTo
	ic.Ptab.LeaderAttr = leader
	ic.Ptab.AlignmentAttr = alignment
	runProps := p.currentRun.X()
	runProps.EG_RunInnerContent = append(runProps.EG_RunInnerContent, ic)
}

func (p *parserState) setFont(attribs map[string]string) {
	font := &FontProps{}
	font.family = ""
	font.size = 0
	font.kern = 0
	font.color = color.Auto
	p.currentStyle.font = font
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "family":
				font.family = val
			case "size", "kern":
				num, err := strconv.ParseFloat(val, 32)
				if err == nil {
					if key == "size" {
						font.size = num
					} else if key == "kern" {
						font.kern = num
					}
				}
			case "color":
				font.color = color.FromHex(val)
			}
		}
	}
}

func (p *parserState) setUnderline(attribs map[string]string) {
	uline := &UnderlineProps{}
	uline.color = color.Auto
	uline.style = wml.ST_UnderlineSingle
	p.currentStyle.underline = uline
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "style":
				num, err := strconv.Atoi(val)
				if err == nil {
					uline.style = wml.ST_Underline(num)
				}
			case "color":
				uline.color = color.FromHex(val)
			}
		}
	}
}

func (p *parserState) setEmphasis(attribs map[string]string) {
	p.currentStyle.emphasisStyle = wml.ST_EmUnderDot
	if len(attribs) != 0 {
		for key, val := range attribs {
			if key == "style" {
				num, err := strconv.Atoi(val)
				if err == nil {
					p.currentStyle.emphasisStyle = wml.ST_Em(num)
				}
			}
		}
	}
}

func setRunStyles(runner *document.Run, flags StyleTags) {
	if (flags & Bold) != 0 {
		runner.Properties().SetBold(true)
	}
	if (flags & Italic) != 0 {
		runner.Properties().SetItalic(true)
	}
	if (flags & Caps) != 0 {
		runner.Properties().SetAllCaps(true)
	}
	if (flags & SmallCaps) != 0 {
		runner.Properties().SetSmallCaps(true)
	}
	if (flags & StrikeThrough) != 0 {
		runner.Properties().SetStrikeThrough(true)
	}
	if (flags & DoubleStrikeThrough) != 0 {
		runner.Properties().SetDoubleStrikeThrough(true)
	}
	if (flags & Outline) != 0 {
		runner.Properties().SetOutline(true)
	}
	if (flags & Shadow) != 0 {
		runner.Properties().SetShadow(true)
	}
	if (flags & Emboss) != 0 {
		runner.Properties().SetEmboss(true)
	}
	if (flags & Imprint) != 0 {
		runner.Properties().SetImprint(true)
	}
	if (flags & NoProof) != 0 {
		runner.Properties().X().NoProof = wml.NewCT_OnOff()
	}
	if (flags & SnapToGrid) != 0 {
		runner.Properties().X().SnapToGrid = wml.NewCT_OnOff()
	}
	if (flags & Vanish) != 0 {
		runner.Properties().X().Vanish = wml.NewCT_OnOff()
	}
	if (flags & WebHidden) != 0 {
		runner.Properties().X().WebHidden = wml.NewCT_OnOff()
	}
	if (flags & RightToLeft) != 0 {
		runner.Properties().X().Rtl = wml.NewCT_OnOff()
	}
	if (flags & SuperScript) != 0 {
		runner.Properties().SetVerticalAlignment(sharedTypes.ST_VerticalAlignRunSuperscript)
	}
	if (flags & SubScript) != 0 {
		runner.Properties().SetVerticalAlignment(sharedTypes.ST_VerticalAlignRunSubscript)
	}

}
