package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var mainURL string = "https://aa.mts.ru/api/v5"
var startCampain string = "/TeleInfo/Campaign/Start?Phone=9134007542&IdCampaign="
var clearResp string = "/TeleInfo/Respondents/Clean?Phone=9134007542&IdCampaign="
var addResp string = "/TeleInfo/Respondents/Add?Phone=9134007542&IdCampaign="
var stopCampain string = "/TeleInfo/Campaign/Stop?Phone=9134007542&IdCampaign="

type respondent struct {
	Phone string `json:"phone"`
}

var authString string = "Api-key "

func main() {
	v := os.Args[1]
	authString += os.Args[2]
	stopReq(v)
	delNumbers(v)
	addRespondents(v, os.Args[3:])
	startReq(v)
	fmt.Println("done")
	os.Exit(0)
}

func startReq(v string) {
	url := mainURL + startCampain + v
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Host":          []string{"aa.mts.ru"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{authString},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("campain start request fail")
	}
	if res.StatusCode == 200 {
		fmt.Println("success")
	} else {
		fmt.Println("request !OK")
	}
}

func delNumbers(v string) {
	url := mainURL + clearResp + v
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header = http.Header{
		"Host":          []string{"aa.mts.ru"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{authString},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("campain delete respondents request fail")
	}
	if res.StatusCode == 200 {
		fmt.Println("success")
	} else {
		fmt.Println("request !OK")
	}
}

func stopReq(v string) {
	url := mainURL + stopCampain + v
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Host":          []string{"aa.mts.ru"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{authString},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("campain stop request fail")
	}
	if res.StatusCode == 200 {
		fmt.Println("success")
	} else {
		fmt.Println("request !OK")
	}
}

func addRespondents(v string, n []string) {
	r := []respondent{}
	for _, s := range n {
		r = append(r, respondent{s})
	}
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	url := mainURL + addResp + v
	client := http.Client{}
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	req.Header = http.Header{
		"Host":          []string{"aa.mts.ru"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{authString},
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("campain add respondents request fail")
	}
	if res.StatusCode == 200 {
		fmt.Println("success")
	} else {
		fmt.Println("request !OK")
	}
}
