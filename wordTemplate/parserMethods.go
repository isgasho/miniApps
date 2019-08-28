package main

import (
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

type parserState struct {
	currentPara      *document.Paragraph
	currentRun       *document.Run
	prev             *parserState
	section          Tags
	currentTag       Tags
	currentTextStyle *TextStyles
}

func NewParserState(currentState *parserState, tagName Tags) *parserState {
	newState := &parserState{}
	newState.prev = currentState
	newState.currentPara = currentState.currentPara
	newState.currentRun = currentState.currentRun
	newState.section = currentState.section
	newState.currentTag = tagName
	newState.currentTextStyle = currentState.currentTextStyle
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
