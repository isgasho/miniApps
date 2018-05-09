package utils

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type ProxyIp struct {
	ip        string
	port      string
	timeTaken time.Duration
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
		working, time := CheckProxy(oneProxy)
		oneProxy.working = working
		oneProxy.timeTaken = time
	}
  for
	for i := 0; i < len(p.allProxyList); i++ {
		for j := i + 1; j < len(p.allProxyList)-i; j++ {
			if p.allProxyList[i].timeTaken > p.allProxyList[j].timeTaken {
				temp := p.allProxyList[i]
				p.allProxyList[i] = p.allProxyList[j]
				p.allProxyList[j] = temp
			}
		}
	}
}

func CheckProxy(ipadr *ProxyIp) (bool, time.Duration) {
	startTime := time.Now()
	timeout := time.Duration(15 * time.Second)
	host := fmt.Sprintf("%s:%s", ipadr.ip, ipadr.port)
	urlProxy := &url.URL{Host: host}

	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(urlProxy)},
		Timeout:   timeout}
	resp, err := client.Get("http://google.com")
	if err != nil {
		return false, time.Time{}
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true, time.Since(startTime)
	}
	return false, time.Time{}
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
	if tds[6].Text() == "yes" {
		return &ProxyIp{ip: tds[0].Text(), port: tds[1].Text()}
	} else {
		return nil
	}
}
