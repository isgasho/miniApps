package main

import (
	"fmt"
	"github.com/devarsh/miniApps/synergyParser/utils"
)

func main() {
	/*	utils.ProxyCheckingUrl = "https://appstore.wipro.com/worklight/apps/services/preview/ContractPartner/common/0/default/index.html"
		x := utils.NewProxy(10)
		fmt.Println("fetching proxy")
		x.LoadProxies()
		fmt.Println("ranking proxy")
		x.RankProxy()
		x.ListAll()
	*/
	uas := utils.NewRandomUA()
	uas.LoadDummyUserAgents()
	fmt.Println(uas.GetRndUserAgent())
	fmt.Println(uas.GetRndUserAgent())
}
