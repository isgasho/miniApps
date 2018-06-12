package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/fatih/color"
	"golang.org/x/net/publicsuffix"
)

func getHeaders(init bool, cookieString, wlInstance string) http.Header {
	customHeader := http.Header{}
	customHeader.Add("accept", "text/javascript, text/html, application/xml, text/xml, */*")
	customHeader.Add("dnt", "1")

	customHeader.Add("user-agent", "")
	if init == true {
		customHeader.Add("accept-language", "en-US,en;q=0.9,es;q=0.8,hi;q=0.7")
		customHeader.Add("upgrade-insecure-requests", "1")
	}
	if init == false {
		customHeader.Add("accept-language", "en-GB")
		customHeader.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
		customHeader.Add("origin", "https://appstore.wipro.com")
		customHeader.Add("referer", initUrl)
		customHeader.Add("x-requested-with", "XMLHttpRequest")
		customHeader.Add("x-wl-app-details", `{"applicationDetails":{"platformVersion":"6.3.0.0","nativeVersion":""}}`)
		customHeader.Add("x-wl-clientlog-env", "common")
		customHeader.Add("x-wl-clientlog-osversion", "UNKNOWN")
		customHeader.Add("x-wl-clientlog-appversion", "1.0")
		customHeader.Add("x-wl-app-version", "1.0")
		customHeader.Add("x-wl-clientlog-model", "UNKNOWN")
		customHeader.Add("x-wl-clientlog-appname", "ContractPartner")
		if cookieString != "" {
			customHeader.Add("cookie", cookieString)
		}
		if wlInstance != "" {
			customHeader.Add("wl-instance-id", wlInstance)
		}
	}
	return customHeader
}

func formData() url.Values {
	form := url.Values{}
	form.Add("adapter", "PartnerWebAdapter")
	form.Add("isAjaxRequest", "true")
	form.Add("procedure", "invokeService")
	form.Add("parameters", "")
	return form
}

func appendCookies(cookies string, cookiesArray []*http.Cookie) string {
	for _, oneCookie := range cookiesArray {
		cookies = cookies + oneCookie.Name + "=" + oneCookie.Value + ";"
	}
	return cookies
}

func initClient() *http.Client {
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
