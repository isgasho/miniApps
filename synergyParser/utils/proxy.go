package utils

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ProxyCheckingUrl = "https://www.google.com"
)

type ProxyIp struct {
	Ip        string
	Port      string
	TimeTaken time.Duration
	working   bool
}

type Proxy struct {
	allProxyList []*ProxyIp
	limit        int
	proxyCursor  int
}

func NewProxy(limit int) *Proxy {
	p := Proxy{limit: limit}
	return &p
}

func (p *Proxy) GetProxy() *ProxyIp {
	p.proxyCursor++
	if p.proxyCursor > len(p.allProxyList)-1 {
		p.proxyCursor = 0
	}
	return p.allProxyList[p.proxyCursor]
}

func (p *Proxy) RankProxy() {
	for _, oneProxy := range p.allProxyList {
		working, time := oneProxy.CheckProxy()
		oneProxy.working = working
		oneProxy.TimeTaken = time
	}
	newAllProxyList := make([]*ProxyIp, 0)
	for _, oneProxy := range p.allProxyList {
		if oneProxy.working {
			newAllProxyList = append(newAllProxyList, oneProxy)
		}
	}
	p.allProxyList = newAllProxyList
	for i := 0; i < len(p.allProxyList); i++ {
		for j := i + 1; j < len(p.allProxyList)-i; j++ {
			if p.allProxyList[i].TimeTaken > p.allProxyList[j].TimeTaken {
				temp := p.allProxyList[i]
				p.allProxyList[i] = p.allProxyList[j]
				p.allProxyList[j] = temp
			}
		}
	}
}

func (p *Proxy) ListAll() {
	for _, oneProxy := range p.allProxyList {
		fmt.Println(oneProxy)
	}
}

func (ipadr *ProxyIp) CheckProxy() (bool, time.Duration) {
	startTime := time.Now()
	timeout := time.Duration(15 * time.Second)
	host := fmt.Sprintf("%s:%s", ipadr.Ip, ipadr.Port)
	urlProxy := &url.URL{Host: host}

	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(urlProxy)},
		Timeout:   timeout}
	resp, err := client.Get(ProxyCheckingUrl)
	if err != nil {
		return false, time.Duration(0)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true, time.Since(startTime)
	}
	return false, time.Duration(0)
}

func (ipadr *ProxyIp) ToString() string {
	return fmt.Sprintf("HostName: %s:%s, Time Taken: %v, Working %t", ipadr.Ip, ipadr.Port, ipadr.TimeTaken, ipadr.working)
}

func (p *Proxy) LoadProxies() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://free-proxy-list.net/", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	respBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	doc := soup.HTMLParse(string(respBuf))

	trs := doc.Find("table", "id", "proxylisttable").Find("tbody").FindAll("tr")
	proxyIpList := make([]*ProxyIp, 0)
	itr := 0
	for _, oneTr := range trs {
		res := getProxyRow(oneTr)
		if res != nil {
			itr++
			proxyIpList = append(proxyIpList, res)
		}
		if itr == p.limit {
			break
		}
	}
	p.allProxyList = proxyIpList
	p.proxyCursor = -1
}

func getProxyRow(node soup.Root) *ProxyIp {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	tds := node.FindAll("td")
	if strings.ToLower(tds[6].Text()) == "yes" && strings.ToLower(tds[4].Text()) == "anonymous" {
		return &ProxyIp{Ip: tds[0].Text(), Port: tds[1].Text()}
	} else {
		return nil
	}
}
