package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/devarsh/miniApps/synergyParser/utils"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func loadPage(client *http.Client) (string, error) {
	startTime := time.Now()
	req, err := http.NewRequest("GET", initURL, nil)
	if err != nil {
		return "", err
	}
	req.Header = getHeaders(true, "", "")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	color.Cyan("Load Login Page Request\nGET %s\nTime Taken:%s\n", initURL, time.Since(startTime))
	defer res.Body.Close()
	cookies := appendCookies("", res.Cookies())
	return cookies, nil
}

func performLogin1(client *http.Client, cookies string) (string, error) {
	reqFormData := formData()
	reqFormData.Set("parameters", getLoginPurpose())
	fmt.Println(getLoginPurpose())
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(reqFormData.Encode()))
	if err != nil {
		return "", err
	}
	req.Header = getHeaders(false, cookies, "")
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		return "", err
	}
	color.Cyan("First Login Req With Cookie %s\nPOST %s\nTime Taken:%s\n", cookies, apiUrl, timeTaken)
	wlIst := gjson.Get(*responseStr, "challenges.wl_antiXSRFRealm.WL-Instance-Id")
	if wlIst.String() == "" {
		return "", fmt.Errorf("Could retrive WL-Instance-Id")
	}
	return wlIst.String(), nil
}

func performLogin2(client *http.Client, cookies string, wlInst string) (bool, error) {
	reqFormData := formData()
	reqFormData.Set("parameters", getLoginPurpose())
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(reqFormData.Encode()))
	if err != nil {
		return false, err
	}
	req.Header = getHeaders(false, cookies, wlInst)
	responseStr, timeTaken, err := utils.RequestMaker(client, req)
	if err != nil {
		return false, err
	}
	color.Cyan("Second Login Req With Cookies: %s & WL-InstToken: %s\nPOST %s\nTime Taken:%s\n", cookies, wlInst, apiUrl, timeTaken)
	success := gjson.Get(*responseStr, "result.authFlag")
	if success.String() == "failure" {
		return false, fmt.Errorf("Invalid Login Credentials")
	} else if success.String() == "success" {
		return true, nil
	}
	return false, fmt.Errorf("Something went wrong while making request contact devarsh")
}
