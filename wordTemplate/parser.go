package main

import (
	"io"
	"log"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
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
				switch tname {
				case Bold, Italic, Caps, SmallCaps,
					StrikeThrough, DoubleStrikeThrough, Outline,
					Shadow, Emboss, Imprint, NoProof,
					SnapToGrid, Vanish, WebHidden, RightToLeft,
					SubScript, SuperScript:
					currentState.currentStyle.flags = currentState.currentStyle.flags ^ tname
				case Underline:
					var attribs map[string]string
					if hasAttrib {
						attribs = getAttributes(tokenizer)
					}
					currentState.setUnderline(attribs)
				case Emphasis:
					var attribs map[string]string
					if hasAttrib {
						attribs = getAttributes(tokenizer)
					}
					currentState.setEmphasis(attribs)
				case Font:
					var attribs map[string]string
					if hasAttrib {
						attribs = getAttributes(tokenizer)
					}
					currentState.setFont(attribs)
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
					currentState.currentStyle.flags = currentState.currentStyle.flags ^ tname
				case Underline:
					currentState.currentStyle.underline = nil
				case Emphasis:
					currentState.currentStyle.emphasisStyle = wml.ST_EmUnset
				case Font:
					currentState.currentStyle.font = nil
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
				} else if currentState.currentStyle != nil && currentState.currentPara != nil {
					run := currentState.currentPara.AddRun()
					if currentState.currentStyle.flags != 0 {
						setRunStyles(&run, currentState.currentStyle.flags)
					}
					if currentState.currentStyle.underline != nil {
						uline := currentState.currentStyle.underline
						run.Properties().SetUnderline(uline.style, uline.color)
					}
					if currentState.currentStyle.emphasisStyle != wml.ST_EmUnset {
						em := wml.NewCT_Em()
						em.ValAttr = currentState.currentStyle.emphasisStyle
						run.Properties().X().Em = em
					}
					if currentState.currentStyle.font != nil {
						fontconfig := currentState.currentStyle.font
						if fontconfig.family != "" {
							run.Properties().SetFontFamily(fontconfig.family)
						}
						if fontconfig.size != 0 {
							run.Properties().SetSize(measurement.Distance(fontconfig.size))
						}
						if fontconfig.color != color.Auto {
							run.Properties().SetColor(fontconfig.color)
						}
						if fontconfig.kern != 0 {
							run.Properties().SetKerning(measurement.Distance(fontconfig.kern))
						}

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
