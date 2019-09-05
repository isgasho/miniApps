package main

import (
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

type ListProps struct {
	level       int
	indentDelta int
	numDefLevel int
}

func (p *parserState) setupOrderedList(attribs map[string]string) {
	var indentDelta = 20
	var hangingIndent = 20
	var align = 11
	var style = 1
	lvl := p.numDef.AddLevel()
	currentLevel := p.currentList.level
	lvl.SetText(fmt.Sprintf("%%%d.", currentLevel+1))
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "indent", "hangIndent", "align", "style":
				num, err := strconv.Atoi(val)
				if err == nil {
					switch key {
					case "indent":
						indentDelta = num
					case "hangIndent":
						hangingIndent = num
					case "align":
						align = num
					case "style":
						style = num
					}
				}
			}
		}
	}
	lvl.Properties().SetLeftIndent(measurement.Distance(int64(currentLevel * indentDelta)))
	lvl.Properties().SetHangingIndent(measurement.Distance(hangingIndent))
	lvl.SetAlignment(wml.ST_Jc(align))
	lvl.SetFormat(wml.ST_NumberFormat(style))
	p.currentList.indentDelta = indentDelta
}

func (p *parserState) setupUnorderedList(attribs map[string]string) {
	var indentDelta = 20
	var hangingIndent = 20
	var align = 11
	var style = ""
	lvl := p.numDef.AddLevel()
	currentLevel := p.currentList.level
	if len(attribs) != 0 {
		for key, val := range attribs {
			switch key {
			case "indent", "hangIndent", "align":
				num, err := strconv.Atoi(val)
				if err == nil {
					switch key {
					case "indent":
						indentDelta = num
					case "hangIndent":
						hangingIndent = num
					case "align":
						align = num
					}
				}
			case "style":
				style = val
			}
		}
	}
	lvl.Properties().SetLeftIndent(measurement.Distance(int64(currentLevel * indentDelta)))
	lvl.Properties().SetHangingIndent(measurement.Distance(hangingIndent))
	lvl.SetAlignment(wml.ST_Jc(align))
	lvl.SetFormat(wml.ST_NumberFormatBullet)
	lvl.RunProperties().SetFontFamily("Symbol")
	lvl.SetText(style)
	p.currentList.indentDelta = indentDelta

}
