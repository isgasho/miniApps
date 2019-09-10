package main

import (
	"time"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func main2() {
	doc := document.New()
	doc.Settings.X().TrackRevisions = wml.NewCT_OnOff()
	para := doc.AddParagraph()
	doc.AddParagraph()
	run := para.AddRun()
	run.AddText("I was working at")
	AddOldValue(&para, "Yahoo", "00475D3C", 0)
	AddNewValue(&para, "Google", "00475D3C", 1)
	doc.SaveToFile("demo.docx")
}
func AddNewValue(p *document.Paragraph, newVal string, revNum string, id int64) {
	exprops := p.Properties().X()
	pprChange := wml.NewCT_PPrChange()
	exprops.PPrChange = pprChange
	pprChange.AuthorAttr = "Devarsh"
	t := time.Now()
	pprChange.DateAttr = &t
	pprChange.IdAttr = id
	run := p.AddRun()
	run.X().RsidDelAttr = &revNum
	ic := wml.NewEG_RunInnerContent()
	run.X().EG_RunInnerContent = append(run.X().EG_RunInnerContent, ic)
	ic.InstrText = wml.NewCT_Text()
	ic.InstrText.Content = newVal
}
func AddOldValue(p *document.Paragraph, oldVal string, revNum string, id int64) {
	exprops := p.Properties().X()
	pprChange := wml.NewCT_PPrChange()
	exprops.PPrChange = pprChange
	pprChange.AuthorAttr = "Devarsh"
	t := time.Now()
	pprChange.DateAttr = &t
	pprChange.IdAttr = id
	run := p.AddRun()
	run.X().RsidDelAttr = &revNum
	ic := wml.NewEG_RunInnerContent()
	run.X().EG_RunInnerContent = append(run.X().EG_RunInnerContent, ic)
	ic.DelText = wml.NewCT_Text()
	ic.DelText.Content = oldVal
}
