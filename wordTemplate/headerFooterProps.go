package main

import (
	"github.com/unidoc/unioffice/schema/soo/wml"
)

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
