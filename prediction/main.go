package main

import (
	"fmt"
	"prediction/mdata"
	"prediction/utils"
	"strings"
)

func main() {

	cleaned := "aaaa,"
	cleaned = strings.TrimSuffix(cleaned, ",")
	fmt.Println(cleaned)

	var (
		path       = fmt.Sprintf("%s%s", "http://datafeed2.tysondata.com:8080", "/datashare/boardcastNearbyMatchSync")
		curTime    = "20221007224747"
		authToken  = utils.Md5V(fmt.Sprintf("%s%s%s%s", utils.Md5V("bob"), "724ff289285f0bd6199142ae08044664", "315ef8a6f535434a860c57feafaf8626", curTime))
		curUrl     = fmt.Sprintf("%s?sourcetype=%d&code=%s&t=%s&auth_token=%s", path, 11, "315ef8a6f535434a860c57feafaf8626", curTime, authToken)
		httpResult = ""
		result     []*mdata.BKvideosData
	)
	httpResult = utils.GetUrl(curUrl)
	err := mdata.Cjson.Unmarshal([]byte(httpResult), &result)

	if err != nil || len(result) < 1 {
		fmt.Println(httpResult, err)
		return
	}
	for i := range result {
		results := result[i]
		if results.Eid == "59013858" {
			fmt.Println("results---------->", result[i])
			fmt.Println("results---------->", result[i].Team1)
			fmt.Println("results---------->", result[i].Animation.Path)
			fmt.Println("results---------->", result[i].MatchStatus)
			fmt.Println("results---------->", result[i].MatchStatusCode)
		}
	}
}
