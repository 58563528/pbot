package thri_api

import (
	"io/ioutil"
	"net/http"
)

const (
	PovertyPornURL = "https://du.shadiao.app/api.php"
)

func PovertyPorn() string {
	res, err := http.Get(PovertyPornURL)
	if err != nil {
		return "毒鸡汤接口错误"
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "毒鸡汤接口错误"
	}

	bodyStr := string(body)
	return bodyStr
}
