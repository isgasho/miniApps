package main

import (
	"strconv"
	"strings"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func setInlineImage(doc *document.Document, runner *document.Run, attribs map[string]string) {
	if len(attribs) != 0 {
		source := attribs["src"]
		img, err := common.ImageFromFile(source)
		if err == nil {
			imgRef, err := doc.AddImage(img)
			if err == nil {
				runner.AddDrawingInline(imgRef)
			}
		}
	}
}

func setAnchoredImage(doc *document.Document, runner *document.Run, attribs map[string]string) {
	if len(attribs) != 0 {
		source := attribs["src"]
		img, err := common.ImageFromFile(source)
		if err == nil {
			imgRef, err := doc.AddImage(img)
			if err == nil {
				anchor, err := runner.AddDrawingAnchored(imgRef)
				if err == nil {
					for key, val := range attribs {
						switch key {
						case "noWrap":
							if val == "on" || val == "true" || val == "1" {
								anchor.SetTextWrapNone()
							}
						case "name":
							anchor.SetName(val)
						case "hAlign", "vAlign", "yOffset", "xOffset", "wrap":
							num, err := strconv.Atoi(val)
							if err == nil {
								switch key {
								case "hAlign":
									anchor.SetHAlignment(wml.WdST_AlignH(num))
								case "vAlign":
									anchor.SetVAlignment(wml.WdST_AlignV(num))
								case "xOffset":
									anchor.SetXOffset(measurement.Distance(num) * measurement.Inch)
								case "yOffset":
									anchor.SetYOffset(measurement.Distance(num) * measurement.Inch)
								case "wrap":
									anchor.SetTextWrapSquare(wml.WdST_WrapText(num))
								}
							}
						case "size", "origin":
							strs := strings.Split(val, ",")
							if len(strs) == 2 {
								first, err1 := strconv.Atoi(strs[0])
								second, err2 := strconv.Atoi(strs[1])
								if err1 == nil && err2 == nil {
									switch key {
									case "size":
										anchor.SetSize(measurement.Distance(first)*measurement.Inch, measurement.Distance(second)*measurement.Inch)
									case "origin":
										anchor.SetOrigin(wml.WdST_RelFromH(first), wml.WdST_RelFromV(second))
									}

								}
							}
						}
					}
				}
			}
		}
	}
}
