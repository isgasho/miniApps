package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/fatih/color"
	"golang.org/x/net/publicsuffix"
)

func getHeaders(init bool, cookieString string) http.Header {
	customHeader := http.Header{}
	customHeader.Add("upgrade-insecure-requests", "1")
	customHeader.Add("user-agent", uas.GetRndUserAgent())
	customHeader.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	customHeader.Add("dnt", "1")
	customHeader.Add("accept-encoding", "gzip, deflate, br")
	customHeader.Add("accept-language", "en-US,en;q=0.9,es;q=0.8,hi;q=0.7")
	if init == false {
		customHeader.Add("content-type", "application/x-www-form-urlencoded")
		customHeader.Add("origin", "https://synergy.wipro.com")
		customHeader.Add("referer", initURL)
	}
	if cookieString != "" {
		customHeader.Add("cookie", cookieString)
	}
	return customHeader
}

func appendCookies(cookies string, cookiesArray []*http.Cookie) string {
	for _, oneCookie := range cookiesArray {
		cookies = cookies + oneCookie.Name + "=" + oneCookie.Value + ";"
	}
	return cookies
}

func makeClient() *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}
	res, proxyAdr := proxy.GetProxy()
	color.Magenta("Setting Up Http Client")
	if res == false {
		color.Magenta("Couldnt setup proxy using local IP")
		client := http.Client{Jar: jar}
		return &client
	}
	color.Magenta("New Proxy Setup :", proxyAdr.ToString())
	host := fmt.Sprintf("%s:%s", proxyAdr.Ip, proxyAdr.Port)
	urlProxy := &url.URL{Host: host}
	client := http.Client{Jar: jar, Transport: &http.Transport{Proxy: http.ProxyURL(urlProxy)}}
	return &client
}
