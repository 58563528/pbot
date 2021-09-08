package utils

import "strings"

func GetRedBagStr(cqm string) string {
	res := ""
	if strings.Contains(cqm, "[CQ:redbag,title=") {
		res = strings.Trim(cqm, "[CQ:redbag,title=")
		res = res[:len(res)-1]
	}

	return res
}
