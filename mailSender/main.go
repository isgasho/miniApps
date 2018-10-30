package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/smtp"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jordan-wright/email"
)

type SmtpAddress struct {
	Address string
	Port    string
}

func (adr *SmtpAddress) getHostAddress() string {
	return fmt.Sprintf("%s:%s", adr.Address, adr.Port)
}

func main() {
	myEmail := flag.String("u", "demo@xyz.com", "Provide the email")
	password := flag.String("pwd", "gctuexoiyawsyypr", "Provide password")
	path := flag.String("p", "./", "Provide file path")
	claim := flag.Bool("c", false, "Claim Yes or NO")
	ccMe := flag.Bool("ccMe", false, "CC me by default")
	to := flag.String("to", "xyz@xyxz.com", "to csv emails")
	cc := flag.String("cc", "", "cc csv emails")
	extensionRegex := flag.String("r", `\.(pdf)`, "Regex for file extensions")
	flag.Parse()
	address := SmtpAddress{Address: "smtp.mail.yahoo.com", Port: "465"}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.mail.yahoo.com",
	}
	auth := smtp.PlainAuth("", *myEmail, *password, address.Address)
	filesChan := FileList(*path, *extensionRegex)
	tos := strings.Split(*to, ",")
	ccs := strings.Split(*cc, ",")
	if *ccMe {
		ccs = append(ccs, *myEmail)
	}
	tos = RemoveDuplicatesFromSlice(tos)
	ccs = RemoveDuplicatesFromSlice(ccs)
	if checkToCcSimilarEmail(tos, ccs) {
		return
	}
	fmt.Println(tos, ccs)
	for oneChan := range filesChan {
		newEmail := email.NewEmail()
		newEmail.From = *myEmail
		newEmail.To = tos
		newEmail.Cc = ccs
		if *claim {
			newEmail.Subject = fmt.Sprintf("Synergy Claim Id:%s", oneChan.Name)
		} else {
			newEmail.Subject = fmt.Sprintf("Synergy Payout Invoice:%s", oneChan.Name)
		}
		stringBefore := `<p>Hi Team,</p>
		<br/>
		Share VP No of below Invoice
		<br/>
		<table style="border: 1px solid black">
			<tr style="border: 1px solid black">
			<td style="border: 1px solid black">%s</td>
			<td style="border: 1px solid black; width:80px"></td>
			</tr>
		</table>
		`
		newEmail.HTML = []byte(fmt.Sprintf(stringBefore, oneChan.Name))
		fmt.Println(oneChan.FullPath)
		newEmail.AttachFile(oneChan.FullPath)
		fmt.Printf("Sending Mail for Invoice %s\n", oneChan.Name)
		err := newEmail.SendWithTLS(address.getHostAddress(), auth, tlsConfig)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type FileDetails struct {
	FullPath string
	Name     string
}

func FileList(root string, extRegex string) <-chan *FileDetails {
	ch := make(chan *FileDetails)
	pattern := regexp.MustCompile(extRegex)
	go func() {
		defer close(ch)
		err := filepath.Walk(root, func(myPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			bytesName := []byte(info.Name())
			loc := pattern.FindIndex(bytesName)
			if len(loc) == 2 {
				newFile := &FileDetails{}
				newFile.FullPath = myPath
				newFile.Name = fmt.Sprintf("%s", bytesName[0:loc[0]])
				ch <- newFile
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error occured", err)
			close(ch)
		}
	}()
	return ch
}

//Removes duplicate & blanks
func RemoveDuplicatesFromSlice(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; !ok {
			m[item] = true
		}
	}
	var result []string
	for item, _ := range m {
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}

func checkToCcSimilarEmail(s1 []string, s2 []string) bool {
	m1 := make(map[string]bool)
	for _, item := range s1 {
		m1[item] = true
	}
	for _, item := range s2 {
		if _, ok := m1[item]; ok {
			fmt.Println("duplicate email in to & cc: ", item)
			return true
		}
	}
	return false
}
