package main

import (
	"bytes"
	"fmt"
	"github.com/devarsh/miniApps/template/csvReader"
	"github.com/devarsh/miniApps/template/envelope"
	"github.com/devarsh/miniApps/template/mdToPdf"
	//"io/ioutil"
	"flag"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Mixture struct {
	Quotation *csvReader.Quotation
	Date      string
}

func initDirectory(dirPath string) {
	subdirs := []string{"./md", "./pdf"}
	for _, oneDir := range subdirs {
		pather := path.Join(dirPath, oneDir)
		if _, err := os.Stat(pather); os.IsNotExist(err) {
			err := os.MkdirAll(pather, 0755)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	inputFile := flag.String("inputFile", "./input.csv", "Specify input CSV file to be used")
	outputDirectory := flag.String("outputDir", "./out", "Output Director where files will be generated")
	ignoreFirstLine := flag.Bool("ignoreFirstLine", true, "Ignore the Header Line of CSV input file")
	quotatonDate := flag.String("qdate", "01-01-2000", "Quotation Date for generation")
	flag.Parse()
	initDirectory(*outputDirectory)
	ncm, err := Asset("templates/COMP/REN/ncm.md")
	if err != nil {
		panic(err)
	}
	ccm, err := Asset("templates/COMP/REN/ccm.md")
	if err != nil {
		panic(err)
	}
	ccmNew, err := Asset("templates/COMP/NEW/ccm.md")
	if err != nil {
		panic(err)
	}
	ncmNew, err := Asset("templates/COMP/NEW/ncm.md")
	if err != nil {
		panic(err)
	}

	nNcm, err := Asset("templates/NONCOMP/REN/ncm.md")
	if err != nil {
		panic(err)
	}
	nCcm, err := Asset("templates/NONCOMP/REN/ccm.md")
	if err != nil {
		panic(err)
	}
	nCcmNew, err := Asset("templates/NONCOMP/NEW/ccm.md")
	if err != nil {
		panic(err)
	}
	nNcmNew, err := Asset("templates/NONCOMP/NEW/ncm.md")
	if err != nil {
		panic(err)
	}
	funcMap := template.FuncMap{
		"machineNames": func(mc []*csvReader.MachineDetails) string {
			mcStr := " Model: "
			mclen := len(mc)
			for index, oneMc := range mc {
				if index == mclen-1 {
					mcStr = mcStr + oneMc.Model + " "
				} else {
					mcStr = mcStr + oneMc.Model + ","
				}
			}
			return mcStr
		},
		"indianCurr": func(amount string) string {
			return IndianCurrComma(amount)
		},
		"indianCurrF": func(amount float64) string {
			return IndianCurrComma(fmt.Sprintf("%.2f", amount))
		},
		"datesFmt": func(dtStr string) string {
			return datesFmt(dtStr)
		},
	}
	/*f1, err := os.Open("./templates/ncm.md")
	if err != nil {
		panic(err)
	}
	f2, err := os.Open("./templates/ccm.md")
	if err != nil {
		panic(err)
	}
	f3, err := os.Open("./templates/ccmNew.md")
	if err != nil {
		panic(err)
	}
	f4, err := os.Open("./templates/ncmNew.md")
	if err != nil {
		panic(err)
	}
	ncm, err := ioutil.ReadAll(f1)
	if err != nil {
		panic(err)
	}
	ccm, err := ioutil.ReadAll(f2)
	if err != nil {
		panic(err)
	}
	ccmNew, err := ioutil.ReadAll(f3)
	if err != nil {
		panic(err)
	}
	ncmNew, err := ioutil.ReadAll(f4)
	if err != nil {
		panic(err)
	}*/

	ncmTmpl, err := template.New("NCM").Funcs(funcMap).Parse(string(ncm))
	if err != nil {
		panic(err)
	}
	ccmTmpl, err := template.New("CCM").Funcs(funcMap).Parse(string(ccm))
	if err != nil {
		panic(err)
	}
	ncmTmplNew, err := template.New("NCMNew").Funcs(funcMap).Parse(string(ncmNew))
	if err != nil {
		panic(err)
	}
	ccmTmplNew, err := template.New("CCMNew").Funcs(funcMap).Parse(string(ccmNew))
	if err != nil {
		panic(err)
	}

	nNcmTmpl, err := template.New("NCM").Funcs(funcMap).Parse(string(nNcm))
	if err != nil {
		panic(err)
	}
	nCcmTmpl, err := template.New("CCM").Funcs(funcMap).Parse(string(nCcm))
	if err != nil {
		panic(err)
	}
	nNcmTmplNew, err := template.New("NCMNew").Funcs(funcMap).Parse(string(nNcmNew))
	if err != nil {
		panic(err)
	}
	nCcmTmplNew, err := template.New("CCMNew").Funcs(funcMap).Parse(string(nCcmNew))
	if err != nil {
		panic(err)
	}

	csvRd := csvReader.NewTemplateReader(*quotatonDate)
	err = csvRd.ReadCsv(*inputFile, *ignoreFirstLine)
	if err != nil {
		fmt.Println(err)
	}
	envelopeGen := envelope.NewEnvelope()
	mdPdf := mdToPdf.NewMdtoPdf(13, "Arial")
	mixture := &Mixture{}
	mixture.Date = csvRd.Date
	for csvRd.Next() {
		var b bytes.Buffer
		oneRecord := csvRd.GetRecord()
		mixture.Quotation = oneRecord
		if oneRecord.QuotationType == "REN" {
			if oneRecord.MachineType == "CCM" {
				ccmTmpl.Execute(&b, mixture)
			} else if oneRecord.MachineType == "NCM" {
				ncmTmpl.Execute(&b, mixture)
			} else {
				continue
			}
		} else if oneRecord.QuotationType == "NEW" {
			if oneRecord.MachineType == "CCM" {
				ccmTmplNew.Execute(&b, mixture)
			} else if oneRecord.MachineType == "NCM" {
				ncmTmplNew.Execute(&b, mixture)
			} else {
				continue
			}
		} else if oneRecord.QuotationType == "RENNON" {
			if oneRecord.MachineType == "CCM" {
				nCcmTmpl.Execute(&b, mixture)
			} else if oneRecord.MachineType == "NCM" {
				nNcmTmpl.Execute(&b, mixture)
			}
		} else if oneRecord.QuotationType == "NEWNON" {
			if oneRecord.MachineType == "CCM" {
				nCcmTmplNew.Execute(&b, mixture)
			} else if oneRecord.MachineType == "NCM" {
				nNcmTmplNew.Execute(&b, mixture)
			}
		} else {
			continue
		}
		bytes := b.Bytes()
		filename := fmt.Sprintf("./Qoutation-%s-%s-%s-%s", oneRecord.SrNo, oneRecord.MachineType, oneRecord.Region, oneRecord.QuotationType)
		filename1 := path.Join(*outputDirectory, "./md", filename)
		filename2 := path.Join(*outputDirectory, "./pdf", filename)

		fs, err := os.Create(filename1 + ".md")
		if err != nil {
			panic(err)
		}
		fs.Write(bytes)
		mdPdf.NewPdf(bytes, filename2, "For any complaints call us on 079-26424229 / 99252 04929 / 99099 58229")
		envelopeGen.NewAddress(oneRecord.Address)
	}
	filename3 := path.Join(*outputDirectory, "./envelope.pdf")
	envelopeGen.GenerateFile(filename3)
}

func IndianCurrComma(amount string) string {
	amounts := strings.Split(amount, ".")
	threeCount := false
	firstPart := amounts[0]
	var secondPart int64
	if len(amounts) > 1 {
		second1Part, err := strconv.ParseInt(amounts[1], 10, 64)
		if err != nil {
			panic(err)
		}
		secondPart = second1Part
	}
	if len(firstPart) <= 3 {
		if len(amounts) < 2 && secondPart == 0 {
			return firstPart + "/-"
		} else {
			return firstPart + "." + amounts[1]
		}
	}

	bytes := []byte(firstPart)
	bytesout := make([]byte, 0)
	commaByte := []byte(",")
	count := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		count++
		if count > 2 && threeCount == true {
			count = 1
			bytesout = append(bytesout, commaByte[0])
		}
		if count > 3 && threeCount == false {
			threeCount = true
			count = 1
			bytesout = append(bytesout, commaByte[0])
		}
		bytesout = append(bytesout, bytes[i])

	}
	for i := len(bytesout)/2 - 1; i >= 0; i-- {
		opp := len(bytesout) - 1 - i
		bytesout[i], bytesout[opp] = bytesout[opp], bytesout[i]
	}
	if len(amounts) < 2 || secondPart == 0 {
		return string(bytesout) + "/-"
	} else {
		return string(bytesout) + "." + amounts[1]
	}
}

func datesFmt(dtStr string) string {
	regex := regexp.MustCompile("(?i)to")
	dates := regex.Split(dtStr, -1)
	if len(dates) == 2 {
		return fmt.Sprintf("%s to %s", dateFormat(dates[0]), dateFormat(dates[1]))
	}
	return dtStr
}

func dateFormat(dtStr string) string {
	if dtStr == "" {
		return "DD-MM-YYYY"
	}
	regex := regexp.MustCompile("[-/]")
	date := regex.Split(strings.TrimSpace((dtStr)), -1)
	if len(date) == 3 {
		dd := date[0]
		mm := date[1]
		yy := date[2]
		yyInt, err := strconv.ParseInt(yy, 10, 64)
		if err != nil {
			panic(err)
		}
		mmInt, err := strconv.ParseInt(mm, 10, 64)
		if err != nil {
			panic(err)
		}
		if yyInt <= 999 {
			yyInt += 2000
		}
		if mmInt > 0 || mmInt < 12 {
			mmStr := time.Month(mmInt)
			mmBytes := []byte(mmStr.String())
			mmStrMin := string(mmBytes[0:3])
			return fmt.Sprintf("%02s-%s-%d", dd, mmStrMin, yyInt)
		}
	}
	return "DD-MM-YYYY"
}
