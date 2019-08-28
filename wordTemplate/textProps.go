package main

import (
	"strconv"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

type TextStyles struct {
	flags         StyleTags
	underline     *UnderlineProps
	emphasisStyle wml.ST_Em
	font          *FontProps
	textHighlight wml.ST_HighlightColor
	textEffect    wml.ST_TextEffect
	textBorder    *TextBorderProps
	textshading   *TextShadingProps
}

type TextShadingProps struct {
	style wml.ST_Shd
	color *wml.ST_HexColor
	fill  *wml.ST_HexColor
}

type TextBorderProps struct {
	style  wml.ST_Border
	color  *wml.ST_HexColor
	frame  *sharedTypes.ST_OnOff
	shadow *sharedTypes.ST_OnOff
	size   *uint64
	space  *uint64
}

type UnderlineProps struct {
	style wml.ST_Underline
	color color.Color
}

type FontProps struct {
	family      string
	size        float64
	kern        float64
	charSpacing float64
	color       color.Color
	scale       string
	csize       string
}

func (p *parserState) setTextShading(attribs map[string]string) {
	shading := &TextShadingProps{}
	shading.style = wml.ST_ShdSolid
	shading.color = nil
	shading.fill = nil
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "style":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					shading.style = wml.ST_Shd(num)
				}
			case "color", "fill":
				clr, err := wml.ParseUnionST_HexColor(val)
				if err == nil {
					switch key {
					case "color":
						shading.color = &clr
					case "fill":
						shading.fill = &clr
					}
				}
			}
		}
	}
}

func (p *parserState) setTextBorders(attribs map[string]string) {
	border := &TextBorderProps{}
	border.style = wml.ST_BorderDouble
	border.color = nil
	border.frame = nil
	border.shadow = nil
	border.size = nil
	border.space = nil
	p.currentTextStyle.textBorder = border
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "style", "size", "space":
				num, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					switch key {
					case "style":
						border.style = wml.ST_Border(num)
					case "size":
						border.size = unioffice.Uint64(uint64(num))
					case "space":
						border.space = unioffice.Uint64(uint64(num))
					}
				}
			case "color":
				clr, err := wml.ParseUnionST_HexColor(val)
				if err == nil {
					border.color = &clr
				}
			case "frame", "shadow":
				onoff, err := wml.ParseUnionST_OnOff(val)
				if err == nil {
					switch key {
					case "frame":
						border.frame = &onoff
					case "shadow":
						border.shadow = &onoff
					}
				}
			}
		}
	}

}

func (p *parserState) setFontStyles(attribs map[string]string) {
	font := &FontProps{}
	font.family = ""
	font.scale = ""
	font.size = 0
	font.kern = 0
	font.charSpacing = 0
	font.color = color.Auto
	p.currentTextStyle.font = font
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "scale":
				font.scale = val
			case "family":
				font.family = val
			case "size", "kern", "charSpacing":
				num, err := strconv.ParseFloat(val, 32)
				if err == nil {
					switch key {
					case "size":
						font.size = num
					case "kern":
						font.kern = num
					case "charSpacing":
						font.charSpacing = num
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
	p.currentTextStyle.underline = uline
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
	p.currentTextStyle.emphasisStyle = wml.ST_EmUnderDot
	if len(attribs) != 0 {
		for key, val := range attribs {
			if key == "style" {
				num, err := strconv.Atoi(val)
				if err == nil {
					p.currentTextStyle.emphasisStyle = wml.ST_Em(num)
				}
			}
		}
	}
}

func (p *parserState) setTextHighlight(attribs map[string]string) {
	p.currentTextStyle.textHighlight = wml.ST_HighlightColorYellow
	if len(attribs) != 0 {
		for key, val := range attribs {
			if key == "style" {
				num, err := strconv.Atoi(val)
				if err == nil {
					p.currentTextStyle.textHighlight = wml.ST_HighlightColor(num)
				}
			}
		}
	}
}

func (p *parserState) setTextEffect(attribs map[string]string) {
	p.currentTextStyle.textEffect = wml.ST_TextEffectShimmer
	if len(attribs) != 0 {
		for key, val := range attribs {
			if key == "style" {
				num, err := strconv.Atoi(val)
				if err == nil {
					p.currentTextStyle.textEffect = wml.ST_TextEffect(num)
				}
			}
		}
	}
}

func applyTextShading(runner *document.Run, shadeConfig *TextShadingProps) {
	shading := wml.NewCT_Shd()
	runner.Properties().X().Shd = shading
	if shadeConfig.style != wml.ST_ShdUnset {
		shading.ValAttr = shadeConfig.style
	}
	if shadeConfig.color != nil {
		shading.ColorAttr = shadeConfig.color
	}
	if shadeConfig.fill != nil {
		shading.FillAttr = shadeConfig.fill
	}
}

func applyTextBorder(runner *document.Run, borderConfig *TextBorderProps) {
	border := wml.NewCT_Border()
	runner.Properties().X().Bdr = border
	if borderConfig.style != wml.ST_BorderUnset {
		border.ValAttr = borderConfig.style
	}
	if borderConfig.color != nil {
		border.ColorAttr = borderConfig.color
	}
	if borderConfig.frame != nil {
		border.FrameAttr = borderConfig.frame
	}
	if borderConfig.shadow != nil {
		border.ShadowAttr = borderConfig.shadow
	}
	if borderConfig.space != nil {
		border.SpaceAttr = borderConfig.space
	}
	if borderConfig.size != nil {
		border.SzAttr = borderConfig.size
	}
}

func applyFontStyles(runner *document.Run, fontconfig *FontProps) {
	if fontconfig.family != "" {
		runner.Properties().SetFontFamily(fontconfig.family)
	}
	if fontconfig.size != 0 {
		runner.Properties().SetSize(measurement.Distance(fontconfig.size))
	}
	if fontconfig.color != color.Auto {
		runner.Properties().SetColor(fontconfig.color)
	}
	if fontconfig.kern != 0 {
		runner.Properties().SetKerning(measurement.Distance(fontconfig.kern))
	}
	if fontconfig.charSpacing != 0 {
		runner.Properties().SetCharacterSpacing(measurement.Distance(fontconfig.charSpacing))
	}
	if fontconfig.scale != "" {
		scale := wml.NewCT_TextScale()
		val, err := wml.ParseUnionST_TextScale(fontconfig.scale)
		if err != nil {
			scale.ValAttr = &val
			runner.Properties().X().W = scale
		}
	}
	if fontconfig.csize != "" {
		cszie := wml.NewCT_HpsMeasure()
		val, err := wml.ParseUnionST_HpsMeasure(fontconfig.csize)
		if err != nil {
			cszie.ValAttr = val
			runner.Properties().X().SzCs = cszie
		}
	}
}

func applyRunStyles(runner *document.Run, flags StyleTags) {
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
