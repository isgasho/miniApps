package parser

import (
	"strconv"
	"strings"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func setGlobalHeadingSize(sizes *GlobalConsts, attribs map[string]string) {
	if len(attribs) != 0 {
		for key, val := range attribs {
			num, err := strconv.ParseFloat(val, 64)
			if err == nil {
				switch key {
				case "h1":
					sizes.h1 = num
				case "h2":
					sizes.h2 = num
				case "h3":
					sizes.h3 = num
				case "h4":
					sizes.h4 = num
				case "h5":
					sizes.h5 = num
				case "h6":
					sizes.h6 = num
				}
			}
		}
	}
}

func setGlobalFontStyle(doc *document.Document, attribs map[string]string) {
	docRunDef := doc.Styles.X().DocDefaults.RPrDefault.RPr
	runProps := document.NewRunProperties(docRunDef) //NewRunProperites is a method declared in library to get around the problem of private variable initialization
	if docRunDef != nil {
		if len(attribs) != 0 {
			for key, val := range attribs {
				switch key {
				case "family":
					runProps.SetFontFamily(val)
				case "size", "kern", "charSpacing":
					num, err := strconv.ParseFloat(val, 32)
					if err == nil {
						switch key {
						case "size":
							runProps.SetSize(measurement.Distance(num))
						case "kern":
							runProps.SetKerning(measurement.Distance(num))
						case "charSpacing":
							runProps.SetCharacterSpacing(measurement.Distance(num))
						}
					}
				case "color":
					clr := color.FromHex(val)
					runProps.SetColor(clr)
				}
			}
		}
	}
}
func setGlobalParaSpacing(doc *document.Document, attribs map[string]string) {
	docParaDef := doc.Styles.X().DocDefaults.PPrDefault
	if docParaDef != nil {
		para := wml.NewCT_PPrGeneral()
		docParaDef.PPr = para
		para.Spacing = wml.NewCT_Spacing()
		paraSpacing := document.NewParagraphSpacing(para.Spacing)
		if len(attribs) != 0 {
			for key, val := range attribs {
				switch key {
				case "after", "before":
					num, err := strconv.ParseInt(val, 10, 64)
					if err == nil {
						switch key {
						case "after":
							paraSpacing.SetAfter(measurement.Distance(num))
						case "before":
							paraSpacing.SetBefore(measurement.Distance(num))
						}
					}
				case "autoAfter", "autoBefore":
					if val == "true" || val == "1" || val == "on" {
						switch key {
						case "autoAfter":
							paraSpacing.SetAfterAuto(true)

						case "autoBefore":
							paraSpacing.SetBeforeAuto(true)
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
							paraSpacing.SetLineSpacing(measurement.Distance(lheightNum), wml.ST_LineSpacingRule(lstyleNum))
						}
					}
				}
			}
		}
	}

}
