package http

import (
	"net/url"
	"strings"
)

func EncodeUrlParams(params map[string]string) string {
	var sb strings.Builder
	for key, value := range params {
		sb.WriteString(url.QueryEscape(key)) //url参数转义
		sb.WriteString("=")
		sb.WriteString(url.QueryEscape(value))
		sb.WriteString("&")
	}

	if sb.Len() > 1 {
		return sb.String()[:sb.Len()-1] //去除末尾的&
	}

	return ""
}

func ParseUrlParams(rawQuery string) map[string]string {
	params := make(map[string]string, 10)
	args := strings.SplitSeq(rawQuery, "&")
	for ele := range args {
		arr := strings.Split(ele, "=")
		if len(arr) == 2 {
			key, _ := url.QueryUnescape(arr[0])
			value, _ := url.QueryUnescape(arr[1])
			params[key] = value
		}
	}

	return params
}
