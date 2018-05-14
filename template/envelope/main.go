package envelope

import (
	"github.com/jung-kurt/gofpdf"
	"strings"
)

type PdfDoc struct {
	pdf *gofpdf.Fpdf
	ht  float64
}

func NewEnvelope() *PdfDoc {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 241.3, Ht: 104.902}, // 241.3 x 104.902
		OrientationStr: "P",
	})
	pdf.SetMargins(80, 20, 10)
	x := &PdfDoc{pdf: pdf, ht: pdf.PointConvert(14)}
	return x
}

func (pdf *PdfDoc) setBold() {
	pdf.pdf.SetFont("Arial", "B", 16)
	pdf.ht = pdf.pdf.PointConvert(16)
}

func (pdf *PdfDoc) setRegular() {
	pdf.pdf.SetFont("Arial", "B", 14)
	pdf.ht = pdf.pdf.PointConvert(14)
}

func (pdf *PdfDoc) write(str string) {
	pdf.pdf.MultiCell(140, pdf.ht, str, "", "L", false)
}

func (pdf *PdfDoc) NewAddress(address string) {
	pdf.pdf.AddPage()
	pdf.setBold()
	pdf.write("To")
	pdf.setRegular()
	addressStr := strings.Split(address, "\n")
	for _, addrOne := range addressStr {
		if strings.Contains(addrOne, "**") {
			newStr := strings.Replace(addrOne, "*", "", -1)
			pdf.setBold()
			pdf.write(newStr)
			pdf.setRegular()
		} else {
			pdf.write(addrOne)
		}
	}
}

func (pdf *PdfDoc) GenerateFile(filePath string) error {
	err := pdf.pdf.OutputFileAndClose(filePath)
	if err != nil {
		return err
	}
	return nil
}
