package main

import (
	"time"

	"fmt"

	"github.com/devarsh/miniApps/synergyParser/utils"
)

func main() {
	x := utils.NewProxy(10, 8, 2)
	x.LoadProxies()
	st := time.Now()
	x.RankProxy()
	x.ListAll()
	fmt.Println("Program Completed in", time.Since(st))
}
