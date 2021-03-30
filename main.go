package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var cookieStr string
var successNumber int

type reportMessage struct {
	HealthCondition             string `json:"healthCondition"`
	TodayMorningTemperature     string `json:"todayMorningTemperature"`
	YesterdayEveningTemperature string `json:"yesterdayEveningTemperature"`
	YesterdayMiddayTemperature  string `json:"yesterdayMiddayTemperature"`
	Location                    string `json:"location"`
}

func init() {
	flag.StringVar(&cookieStr, "cookie", "", "")
}

func main() {
	flag.Parse()
	cookies := strings.Split(cookieStr, "#")
	for i, cookie := range cookies {
		log.Printf("-----------------------\n")
		log.Printf("准备为第%d位学生上报\n", i+1)
		time.Sleep(3 * time.Second)
		checkReport(cookie, i+1)
	}
	log.Printf("---------------上报完成--------------\n")
	log.Printf("成功上报%d位同学的体温\n", successNumber)
}

func reportFault(id int) {
	log.Printf("第%d位同学上报失败\n", id)
}

// check if it had report
func checkReport(cookie string, id int) {
	client := &http.Client{}
	url := "https://jzsz.uestc.edu.cn/wxvacation/checkRegisterNew"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("NewRequest error: %v\n", err)
		reportFault(id)
		return
	}
	request.Header.Add("content-type", "application/json")
	request.Header.Add("User-Agent", "Mozilla / 5.0(Linux; Android 9; G8441 Build / 47.2.A .6 .30; wv) AppleWebKit / 537.36(KHTML, like Gecko) Version / 4.0 Chrome / 66.0 .3359 .158 Mobile Safari / 537.36 MicroMessenger / 7.0 .13 .1640(0x27000D39) Process / appbrand3 NetType / WIFI Language / zh_CN ABI / arm64 WeChat / arm64")
	//request.Header.Add("Accept-Encoding", "gzip, compress, br, deflate")
	request.Header.Add("Content-Length", "2")
	request.Header.Add("encode", "false")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("x-tag", "flyio")
	request.Header.Add("charset", "utf-8")
	request.Header.Add("cookie", cookie)
	request.Header.Add("Referer", "https://servicewechat.com/wx521c0c16b77041a0/28/page-frame.html")
	response, err := client.Do(request)
	if err != nil {
		log.Printf("client.Do(request) error: %v\n", err)
		reportFault(id)
		return
	}
	defer response.Body.Close()
	context, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll(response.Body) error: %v\n", err)
		reportFault(id)
		return
	}
	responseData := make(map[string]interface{})
	err = json.Unmarshal(context, &responseData)
	if err != nil {
		log.Printf("json.Unmarshal(context, response) error: %v\n", err)
		reportFault(id)
		return
	}
	data := responseData["data"]
	v, ok := data.(map[string]interface{})
	if !ok {
		log.Printf("type assert error\n")
		reportFault(id)
		return
	}
	isCheck, ok := (v["appliedTimes"]).(float64)
	if !ok {
		log.Printf("type assert error\n")
		reportFault(id)
		return
	}
	if isCheck == 0 {
		log.Printf("正在为第%d位学生上报\n", id)
		DoReport(cookie, id)
	} else if isCheck == 1 {
		log.Printf("第%d位同学已经上报过了\n", id)
		reportFault(id)
		return
	} else {
		log.Printf("response data has been changed\n")
		reportFault(id)
		return
	}

}

func DoReport(cookie string, id int) {
	url := "https://jzsz.uestc.edu.cn/wxvacation/monitorRegisterForReturned"
	oneReportMessage := reportMessage{
		HealthCondition:             "正常",
		TodayMorningTemperature:     "36°C~36.5°C",
		YesterdayEveningTemperature: "36°C~36.5°C",
		YesterdayMiddayTemperature:  "36°C~36.5°C",
		Location:                    "四川省成都市郫都区银杏大道",
	}
	jsons, err := json.Marshal(oneReportMessage)
	if err != nil {
		log.Printf("json.Marshal error, err: %v\n", err)
		reportFault(id)
		return
	}
	result := string(jsons)
	jsoninfo := strings.NewReader(result)

	request, _ := http.NewRequest("POST", url, jsoninfo)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("User-Agent", "Mozilla / 5.0(Linux; Android 10; LIO-AN00 Build/HUAWEILIO-AN00; wv) AppleWebKit / 537.36(KHTML, like Gecko) Version / 4.0 Chrome / 66.0 .3359 .158 Mobile Safari / 537.36 MicroMessenger / 7.0 .13 .1640(0x27000D39) Process / appbrand3 NetType / WIFI Language / zh_CN ABI / arm64 WeChat / arm64")
	//request.Header.Add("Accept-Encoding","gzip, compress, br, deflate")
	request.Header.Add("Content-Length", "220")
	request.Header.Add("encode", "false")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("x-tag", "flyio")
	request.Header.Add("charset", "utf-8")
	request.Header.Add("cookie", cookie)
	request.Header.Add("Referer", "https://servicewechat.com/wx521c0c16b77041a0/28/page-frame.html")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("request error, err: %v\n", err)
	}
	defer response.Body.Close()
	context, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll(response.Body) error: %v\n", err)
		reportFault(id)
		return
	}
	responseData := make(map[string]interface{})
	err = json.Unmarshal(context, &responseData)
	if err != nil {
		log.Printf("json.Unmarshal(context, response) error: %v\n", err)
		reportFault(id)
		return
	}
	v, ok := responseData["status"]
	if !ok {
		log.Printf("responseData data form had been changed, error: %v\n", err)
		reportFault(id)
		return
	}
	status := v.(bool)
	if status == false {
		log.Printf("status is false, student has report\n")
		reportFault(id)
		return
	}

	log.Printf("第%d位同学签到成功\n", id)
	successNumber++
}
