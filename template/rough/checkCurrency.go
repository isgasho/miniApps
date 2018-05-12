package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("(\\d+?)(?=(\\d\\d)+(\\d)(?!\\d))(\\.\\d+)")
	re.Find([]byte("999999"))
	fmt.Println(re)

}
