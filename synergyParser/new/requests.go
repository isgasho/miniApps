package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func loadPage() (string, error) {
	startTime := time.Now()
	req, err := http.NewRequest("GET", initUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header = getHeaders(true, "", "")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	color.Cyan("Load Login Page Request\nGET %s\nTime Taken:%s\n", initUrl, time.Since(startTime))
	defer res.Body.Close()
	cookies := appendCookies("", res.Cookies())

	return cookies, nil
}

func performLogin1(cookies string) (string, error) {
	reqFormData := formData()
	reqFormData.Set("parameters", getLoginPurpose())
	startTime := time.Now()
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(reqFormData.Encode()))
	if err != nil {
		return "", err
	}
	req.Header = getHeaders(false, cookies, "")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	color.Cyan("First Login Req With Cookie %s\nPOST %s\nTime Taken:%s\n", cookies, apiUrl, time.Since(startTime))
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	responseStr := string(body)
	wlIst := gjson.Get(responseStr, "challenges.wl_antiXSRFRealm.WL-Instance-Id")
	if wlIst.String() == "" {
		return "", fmt.Errorf("Could retrive WL-Instance-Id")
	}

	return wlIst.String(), nil
}

func performLogin2(cookies string, wlInst string) (bool, error) {
	reqFormData := formData()
	reqFormData.Set("parameters", getLoginPurpose())
	startTime := time.Now()
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(reqFormData.Encode()))
	if err != nil {
		return false, err
	}
	req.Header = getHeaders(false, cookies, wlInst)
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	color.Cyan("Second Login Req With Cookies: %s & WL-InstToken: %s\nPOST %s\nTime Taken:%s\n", cookies, wlInst, apiUrl, time.Since(startTime))
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	responseStr := string(body)
	success := gjson.Get(responseStr, "result.authFlag")

	if success.String() == "failure" {
		return false, fmt.Errorf("Invalid Login Credentials")
	} else if success.String() == "success" {
		return true, nil
	}
	return false, fmt.Errorf("Something went wrong while making request contact devarsh")
}
