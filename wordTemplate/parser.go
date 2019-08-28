package main

import (
	"io"
	"log"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"golang.org/x/net/html"
)

func parser(tokenizer *html.Tokenizer, ancestorState *parserState) {
	if ancestorState == nil || tokenizer == nil {
		return
	}
	currentState := ancestorState
	doc := document.New()
	isDone := false
	for !isDone {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				isDone = true
				break
			}
			log.Fatalf("error tokenizing html: %v", tokenizer.Err())
		case html.StartTagToken:
			tn, hasAttrib := tokenizer.TagName()
			tnameStr := string(tn)
			if tname, ok := WhiteListTags[tnameStr]; ok {
				switch tname {
				case Document:
					currentState = NewParserState(currentState, tname)
					currentState.section = tname
				case Body:
					currentState = NewParserState(currentState, tname)
					currentState.section = tname
				case Center, Left, Right:
					currentState = NewParserState(currentState, tname)
					run := currentState.currentPara.AddRun()
					currentState.currentRun = &run
					switch tname {
					case Left:
						currentState.setAlignmentTab(wml.ST_PTabRelativeToMargin, wml.ST_PTabLeaderNone, wml.ST_PTabAlignmentLeft)
					case Center:
						currentState.setAlignmentTab(wml.ST_PTabRelativeToMargin, wml.ST_PTabLeaderNone, wml.ST_PTabAlignmentCenter)
					case Right:
						currentState.setAlignmentTab(wml.ST_PTabRelativeToMargin, wml.ST_PTabLeaderNone, wml.ST_PTabAlignmentRight)
					}
				default:
					if currentState.section == Document {
						switch tname {
						case PageHeader:
							hdr := doc.AddHeader()
							doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
							hdrPara := hdr.AddParagraph()
							currentState = NewParserState(currentState, tname)
							currentState.currentPara = &hdrPara
							currentState.setHeaderFooterParagraphPropsPstyle("Header")
						case PageFooter:
							ftr := doc.AddFooter()
							doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)
							ftrPara := ftr.AddParagraph()
							currentState = NewParserState(currentState, tname)
							currentState.currentPara = &ftrPara
							currentState.setHeaderFooterParagraphPropsPstyle("Footer")
						case DocProps:
							currentState = NewParserState(currentState, tname)
							currentState.section = DocProps
						}
					} else if currentState.section == DocProps {
						switch tname {
						case Title, Author, Description,
							Category, Version, Application, Company:
							currentState = NewParserState(currentState, tname)
						}
					}
				}
			} else if tname, ok := WhiteListStyleTags[tnameStr]; ok {
				var attribs map[string]string
				if hasAttrib {
					attribs = getAttributes(tokenizer)
				}
				switch tname {
				case Bold, Italic, Caps, SmallCaps,
					StrikeThrough, DoubleStrikeThrough, Outline,
					Shadow, Emboss, Imprint, NoProof,
					SnapToGrid, Vanish, WebHidden, RightToLeft,
					SubScript, SuperScript:
					currentState.currentTextStyle.flags = currentState.currentTextStyle.flags ^ tname
				case Underline:
					currentState.setUnderline(attribs)
				case Emphasis:
					currentState.setEmphasis(attribs)
				case Font:
					currentState.setFontStyles(attribs)
				case TextHighlight:
					currentState.setTextHighlight(attribs)
				case TextEffect:
					currentState.setTextEffect(attribs)
				case TextBorder:
					currentState.setTextBorders(attribs)
				case TextShading:
					currentState.setTextShading(attribs)
				}
			}
			//fmt.Printf("%+v %p <%s>\n", currentState, currentState, tnameStr)
		case html.EndTagToken:
			tn, _ := tokenizer.TagName()
			tnameStr := string(tn)
			//fmt.Printf("%+v %p </%s>\n", currentState, currentState, tnameStr)
			if tname, ok := WhiteListStyleTags[tnameStr]; ok {
				switch tname {
				case Bold, Italic, Caps, SmallCaps,
					StrikeThrough, DoubleStrikeThrough, Outline,
					Shadow, Emboss, Imprint, NoProof,
					SnapToGrid, Vanish, WebHidden, RightToLeft,
					SubScript, SuperScript:
					currentState.currentTextStyle.flags = currentState.currentTextStyle.flags ^ tname
				case Underline:
					currentState.currentTextStyle.underline = nil
				case Emphasis:
					currentState.currentTextStyle.emphasisStyle = wml.ST_EmUnset
				case Font:
					currentState.currentTextStyle.font = nil
				case TextHighlight:
					currentState.currentTextStyle.textHighlight = wml.ST_HighlightColorUnset
				case TextEffect:
					currentState.currentTextStyle.textEffect = wml.ST_TextEffectUnset
				case TextBorder:
					currentState.currentTextStyle.textBorder = nil
				case TextShading:
					currentState.currentTextStyle.textshading = nil
				}
			} else if _, ok := WhiteListTags[tnameStr]; ok {
				currentState = currentState.prev
			}
		case html.SelfClosingTagToken:
			tn, _ := tokenizer.TagName()
			tnameStr := string(tn)
			if tname, ok := WhiteListSelfTags[tnameStr]; ok {
				switch tname {
				case FieldCurrentPage:
					run := currentState.currentPara.AddRun()
					run.AddField(document.FieldCurrentPage)
				case FieldNumberofPages:
					run := currentState.currentPara.AddRun()
					run.AddField(document.FieldNumberOfPages)
				case PageBreak:
					run := currentState.currentPara.AddRun()
					run.AddPageBreak()
				case LineBreak:
					run := currentState.currentPara.AddRun()
					run.AddBreak()
				}
			}
		case html.TextToken:
			txt := string(tokenizer.Text())
			if currentState != nil {
				if currentState.section == DocProps {
					switch currentState.currentTag {
					case Title:
						doc.CoreProperties.SetTitle(txt)
					case Author:
						doc.CoreProperties.SetAuthor(txt)
					case Description:
						doc.CoreProperties.SetDescription(txt)
					case Category:
						doc.CoreProperties.SetCategory(txt)
					case Version:
						doc.AppProperties.SetApplicationVersion(txt)
					case Application:
						doc.AppProperties.SetApplication(txt)
					case Company:
						doc.AppProperties.SetCompany(txt)
					}
				} else if currentState.currentTextStyle != nil && currentState.currentPara != nil {
					run := currentState.currentPara.AddRun()
					if currentState.currentTextStyle.flags != 0 {
						applyRunStyles(&run, currentState.currentTextStyle.flags)
					}
					if currentState.currentTextStyle.underline != nil {
						uline := currentState.currentTextStyle.underline
						run.Properties().SetUnderline(uline.style, uline.color)
					}
					if currentState.currentTextStyle.emphasisStyle != wml.ST_EmUnset {
						em := wml.NewCT_Em()
						em.ValAttr = currentState.currentTextStyle.emphasisStyle
						run.Properties().X().Em = em
					}
					if currentState.currentTextStyle.font != nil {
						applyFontStyles(&run, currentState.currentTextStyle.font)
					}
					if currentState.currentTextStyle.textHighlight != wml.ST_HighlightColorUnset {
						hl := wml.NewCT_Highlight()
						hl.ValAttr = currentState.currentTextStyle.textHighlight
						run.Properties().X().Highlight = hl
					}
					if currentState.currentTextStyle.textEffect != wml.ST_TextEffectUnset {
						te := wml.NewCT_TextEffect()
						te.ValAttr = currentState.currentTextStyle.textEffect
						run.Properties().X().Effect = te
					}
					if currentState.currentTextStyle.textBorder != nil {
						applyTextBorder(&run, currentState.currentTextStyle.textBorder)
					}
					if currentState.currentTextStyle.textshading != nil {
						applyTextShading(&run, currentState.currentTextStyle.textshading)
					}
					run.AddText(txt)
				} else if currentState.currentRun != nil {
					currentState.currentRun.AddText(txt)
				}
			}
		}
	}
	doc.SaveToFile("out.docx")
}
