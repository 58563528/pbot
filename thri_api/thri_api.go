package thri_api

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

const (
	HistoryTodayURL = "https://api.qzone.work/api/today.history"
)

type HistoryTodayData struct {
	Code int                   `json:"code"`
	List []HistoryTodayContent `json:"list"`
}

type HistoryTodayContent struct {
	Year  string `json:"year"`
	Link  string `json:"link"`
	Title string `json:"title"`
}

func HistoryToday() string {
	res, err := http.Get(HistoryTodayURL)
	if err != nil {
		return "历史上的今天接口错误"
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "历史上的今天接口错误"
	}

	bodyStr := string(body)
	codeResult := gjson.Get(bodyStr, "code")
	if codeResult.Int() != 10000 {
		return "历史上的今天接口错误"
	}

	year := gjson.Get(bodyStr, "data.list.0.year")
	thing := gjson.Get(bodyStr, "data.list.0.title")
	link := gjson.Get(bodyStr, "data.list.0.link")

	dataStr := fmt.Sprintf("时间：%s\n"+
		"链接：%s\n"+
		"事件：%s\n", year.String(), link.String(), thing.String())
	return dataStr
}
