package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"time"
)

type result struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index   string   `json:"_index"`
			Id      string   `json:"_id"`
			Score   float64  `json:"_score"`
			Ignored []string `json:"_ignored"`
			Source  struct {
				SiteId              int         `json:"site_id"`
				ResettledSum        int         `json:"resettled_sum"`
				OddsType            string      `json:"odds_type"`
				ValidBetAmount      float64     `json:"valid_bet_amount"`
				StartTime           time.Time   `json:"start_time"`
				CashoutReturnAmount int         `json:"cashout_return_amount"`
				GameType            int         `json:"game_type"`
				BetOpportunity      int         `json:"bet_opportunity"`
				MemberId            int64       `json:"member_id"`
				GamePlayInfo        string      `json:"game_play_info"`
				Version             string      `json:"@version"`
				BuybackNetAmount    int         `json:"buyback_netAmount"`
				GameUserIp          string      `json:"game_user_ip"`
				VenueName           string      `json:"venue_name"`
				NetAmount           float64     `json:"net_amount"`
				RemainingBetAmount  int         `json:"remaining_bet_amount"`
				TerminalType        string      `json:"terminal_type"`
				UserRegisterTime    time.Time   `json:"user_register_time"`
				BetTime             time.Time   `json:"bet_time"`
				BuybackAmount       int         `json:"buyback_amount"`
				AgentType           interface{} `json:"agent_type"`
				VenueBillNo         string      `json:"venue_bill_no"`
				MatchId             string      `json:"match_id"`
				Timestamp           time.Time   `json:"@timestamp"`
				BillNo              string      `json:"bill_no"`
				BetAmount           int         `json:"bet_amount"`
				VenueUsername       string      `json:"venue_username"`
				GameCode            string      `json:"game_code"`
				Event               struct {
					Original string `json:"original"`
				} `json:"event"`
				MemberName      string    `json:"member_name"`
				TopId           int64     `json:"top_id"`
				Id              string    `json:"id"`
				PlayInfoJson    string    `json:"play_info_json"`
				GameName        string    `json:"game_name"`
				UpdatedAt       time.Time `json:"updated_at"`
				Flag            int       `json:"flag"`
				SettleTime      time.Time `json:"settle_time"`
				SourceUrl       string    `json:"source_url"`
				IsManySettled   int       `json:"is_many_settled"`
				EsIndexName2    string    `json:"es_index_name2"`
				EarlySettleFlag int       `json:"early_settle_flag"`
				Odds            float64   `json:"odds"`
				Handicap        string    `json:"handicap"`
				EsIndexName     string    `json:"es_index_name"`
				PlayInfo        string    `json:"play_info"`
				CreatedAt       time.Time `json:"created_at"`
				VenueId         int       `json:"venue_id"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	Aggregations struct {
		TypesCount struct {
			Value int `json:"value"`
		} `json:"types_count"`
		ValidBetAmount struct {
			Value float64 `json:"value"`
		} `json:"valid_bet_amount"`
	} `json:"aggregations"`
}

func main() {

	// 读取 Excel 文件
	xlsx, err := excelize.OpenFile("./file/toHandleId1.xlsx")
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return
	}

	// 读取 constant 对应的值
	constantMap := make(map[string]string)
	var idsToCheck []string
	rows := xlsx.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error reading Excel rows:", err)
		return
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		constantMap[row[0]] = row[1]
		idsToCheck = append(idsToCheck, row[0])
	}
	// Read the JSON file
	jsonFile, err := os.Open("./file/y6_2024-05-19.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	// Read the file contents into a byte slice
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal the byte slice into the result struct
	var results result
	json.Unmarshal(byteValue, &results)

	// Check if the IDs exist in the JSON data and write to a file if they do
	for _, idToCheck := range idsToCheck {
		for _, hit := range results.Hits.Hits {
			if hit.Id == idToCheck {
				query := ""
				// If the ID exists, write the Elasticsearch PUT query to a file
				// Construct the PUT query
				//putQuery := fmt.Sprintf(`{"index": {"_id": "%s"} }`, idToCheck)
				putQuery := fmt.Sprintf(`PUT %s/_doc/%s`, hit.Index, idToCheck)

				query = putQuery + "\n" + writePrettyJSON(hit.Source) + "\n"
				// Define the file name
				fileName := "./file/es_put_query.json"

				file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					break
				}
				defer file.Close()

				_, err = file.WriteString(query + "\n")
				if err != nil {
					break
				}
				fmt.Println("Elasticsearch PUT query written to file:", fileName)
				// Here you would send the PUT query to Elasticsearch
				// For example, using an HTTP client to perform the PUT request
			}
		}
	}
}

func writePrettyJSON(data interface{}) string {
	// Convert data to JSON with indentation
	prettyJSON, _ := json.MarshalIndent(data, "", "    ")
	return string(prettyJSON)
}
