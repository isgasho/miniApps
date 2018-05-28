package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

func RequestMaker(client *http.Client, req *http.Request) (*string, *time.Duration, error) {
	startTime := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	endTime := time.Since(startTime)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	responseStr := string(body)
	return &responseStr, &endTime, nil
}
