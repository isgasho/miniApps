package main
/*
import (
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
)

func init() {
	soup.SetDebug(false)
}

fetchTimeSheetList(respBody string) []*EmployeeAttendance {
	if respBody == "" {
		return nil
	}
	doc := soup.HTMLParser(respBody)
	trs := doc.Find("form","name","TimeSheet").
	Find("table","class","TbPageBorder").
	Find("table","class","TbStyle").
	Find("tbody").Find("tbody").FindAll("tr")
	details := make([]EmployeeAttendance,0)
	for _, oneTr := range trs {
		oneTrd := fetchTimeSheetDay(oneTr)
		if oneTrd != nil {
			details := append(details, oneTrd)
		}
	}
	return details

}

fetchTimeSheetDay(root soup.Parent) {
	tdsList := parent.FindAll("td")
	data := EmployeeAttendance{}
	if len(tdsList) == 9 {
		data.AttendaceDate = strings.TrimSpace(tdsList[0].text())
		data.SwipeInDate = strings.TrimSpace(tdsList[1].text())
		data.InTime = strings.TrimSpace(tdsList[2].text())
		data.SwipeOutDate = strings.TrimSpace(tdsList[3].text())
		data.TimeOut = strings.TrimSpace(tdsList[4].text())
		data.TimeFromAttenSys = strings.TrimSpace(tdsList[5].text())
		data.LunchTime = strings.TrimSpace(tdsList[6].text())
		data.ActualTimeWorked = strings.TrimSpace(tdsList[7].text())
		data.Remarks = strings.TrimSpace(tdsList[8].text())
	}
	if len(tdsList) == 2 {
		data.AttendaceDate = strings.TrimSpace(tdsList[0])
	}
}
*/