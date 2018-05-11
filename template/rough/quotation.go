package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 210, Ht: 297},
		OrientationStr: "P",
	})
	fontSize := 14.0
	pdf.SetMargins(10, 75, 20)
	pdf.SetFont("Arial", "B", fontSize)
	swidth := pdf.GetStringWidth("April 02, 2018")
	ht := pdf.PointConvert(fontSize)
	pdf.AddPage()

	pdf.CellFormat(20, ht, "VRPL:KNSB:REN:NCM:AMC:0418:001", "", 0, "L", false, 0, "")
	pdf.CellFormat(210-20-swidth, ht, "April 02, 2018", "", 1, "R", false, 0, "")
	pdf.Ln(ht)
	pdf.MultiCell(170, ht, "To,", "", "L", false)
	pdf.MultiCell(170, ht, "The Branch Manager", "", "L", false)
	pdf.MultiCell(170, ht, "The Kukarwada Nagarik Sahakari Bank Ltd", "", "L", false)
	pdf.SetFont("Arial", "", fontSize)
	pdf.MultiCell(170, ht, "Head Office", "", "L", false)
	pdf.MultiCell(170, ht, "Kukarwada", "", "L", false)
	pdf.MultiCell(170, ht, "[N.Guj]", "", "L", false)
	pdf.MultiCell(170, ht, "Pin - 3800001", "", "L", false)
	pdf.Ln(ht)

	pdf.MultiCell(170, ht, "Dear Sir,", "", "L", false)
	pdf.Ln(ht)
	pdf.SetFont("Arial", "B", fontSize)
	pdf.MultiCell(170, ht, "Sub: Renewal offer for Comphrensive Annual Maintenance Contract for your Robomak: Floor Model - Note Counting Machine.", "", "L", false)
	pdf.Ln(ht)

	pdf.SetFont("Arial", "", fontSize)
	pdf.Write(ht, "This is in connection,\n we have enclosed herewith the ")
	pdf.SetFont("Arial", "B", fontSize)
	pdf.Write(ht, "Note Counting Machines- AMC Period")
	pdf.SetFont("Arial", "", fontSize)
	pdf.Ln(ht)
	pdf.Image("./download.jpeg", pdf.GetX(), pdf.GetY(), 170, 50, true, "", 0, "")
	pdf.Ln(ht)

	var data [][]string
	data = append(data, []string{"Name", "Age", "Date of Birth"})
	data = append(data, []string{"Devarsh Shah\nnew line\n is awesome", "27", "30-jan-1990"})
	data = append(data, []string{"Dvija Shah", "18", "15-oct-1997"})

	str := "devarsh m shah w"
	fmt.Println(pdf.GetStringWidth(str))
	x, y := pdf.GetXY()
	ptSize, unitSize := pdf.GetFontSize()
	marginSize := pdf.GetCellMargin()

	fmt.Println("LineHeight", ht)
	fmt.Println("MarginSize:", marginSize)
	fmt.Println("FONT: point Size:", ptSize, "  UnitSize:", unitSize)
	fmt.Println("X & Y:", x, y)
	pdf.MultiCell(40, ht, str, "LRTB", "L", false)
	x, y = pdf.GetXY()
	fmt.Println("X & Y:", x, y)
	err := pdf.OutputFileAndClose("./quotation.pdf")
	if err != nil {
		fmt.Println(err)
	}

}
