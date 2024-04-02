package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"time"
)

var Cjson = jsoniter.ConfigCompatibleWithStandardLibrary

type TCGData1 struct {
	BetAmount       int     `json:"betAmount"`
	GameCode        string  `json:"gameCode"`
	BetOrderNo      string  `json:"betOrderNo"`
	BetTime         string  `json:"betTime"`
	TransTime       string  `json:"transTime"`
	BetContentId    string  `json:"betContentId"`
	PlayCode        string  `json:"playCode"`
	OrderNum        string  `json:"orderNum"`
	Chase           string  `json:"chase"`
	Numero          string  `json:"numero"`
	MerchantCode    string  `json:"merchantCode"`
	BettingContent  string  `json:"bettingContent"`
	PlayId          int     `json:"playId"`
	FreezeTime      string  `json:"freezeTime"`
	Multiple        int     `json:"multiple"`
	Remark          string  `json:"remark"`
	PlayName        string  `json:"playName"`
	GameGroupName   string  `json:"gameGroupName"`
	Device          string  `json:"device"`
	PlanBetAmount   int     `json:"planBetAmount"`
	BetNum          int     `json:"betNum"`
	WinAmount       float64 `json:"winAmount"`
	NetPNL          float64 `json:"netPNL"`
	BetStatus       string  `json:"betStatus"`
	SettlementTime  string  `json:"settlementTime"`
	ActualBetAmount string  `json:"actualBetAmount"`
	ExceedWinAmount float64 `json:"exceedWinAmount"`
	Details         string  `json:"details"`
	WinningNumber   string  `json:"winningNumber"`
	ClientIp        string  `json:"clientIp"`
	PlayBonus       float64 `json:"playBonus"`
	Username        string  `json:"username"`
	ProductType     string  `json:"productType"`
	Single          bool    `json:"single"`
	DetailStatusId  int     `json:"detailStatusId"`
}

type TCGData struct {
	BetAmount       string `json:"betAmount"`
	GameCode        string `json:"gameCode"`
	BetOrderNo      string `json:"betOrderNo"`
	BetTime         string `json:"betTime"`
	TransTime       string `json:"transTime"`
	BetContentId    string `json:"betContentId"`
	PlayCode        string `json:"playCode"`
	OrderNum        string `json:"orderNum"`
	Chase           string `json:"chase"`
	Numero          string `json:"numero"`
	MerchantCode    string `json:"merchantCode"`
	BettingContent  string `json:"bettingContent"`
	PlayId          string `json:"playId"`
	FreezeTime      string `json:"freezeTime"`
	Multiple        string `json:"multiple"`
	Remark          string `json:"remark"`
	PlayName        string `json:"playName"`
	GameGroupName   string `json:"gameGroupName"`
	Device          string `json:"device"`
	WinAmount       string `json:"winAmount"`
	NetPNL          string `json:"netPNL"`
	BetStatus       string `json:"betStatus"`
	SettlementTime  string `json:"settlementTime"`
	ActualBetAmount string `json:"actualBetAmount"`
	ExceedWinAmount string `json:"exceedWinAmount"`
	Details         string `json:"details"`
	WinningNumber   string `json:"winningNumber"`
	ClientIp        string `json:"clientIp"`
	PlayBonus       string `json:"playBonus"`
	UserName        string `json:"username"`
	ProductType     string `json:"productType"`
	Single          bool   `json:"single"`
	DetailStatusId  string `json:"detailStatusId"`
}

type response struct {
	List []*TCGData `json:"list"`
	Page struct {
		CurrentPage int `json:"currentPage"`
		PageSize    int `json:"pageSize"`
		Total       int `json:"total"`
	} `json:"page"`
}

type response1 struct {
	List []*TCGData1 `json:"list"`
	Page struct {
		CurrentPage int `json:"currentPage"`
		PageSize    int `json:"pageSize"`
		Total       int `json:"total"`
	} `json:"page"`
}

type Data struct {
	aa string
}

type User struct {
	Aaa  string `json:"aaa"`
	Data []*TCGData1
}

func (p *Data) NewPartnerErr(format string, sleep time.Duration, a ...interface{}) *Data {
	time.Sleep(sleep * time.Second)
	p.aa = fmt.Sprintf(format, a...)
	return &Data{fmt.Sprintf(format, a...)}
}

type MyStruct struct {
	Field1 string
	Field2 int
	Field3 time.Time
}

func myMethod(str string, slice interface{}, fieldName string) {
	value := reflect.ValueOf(slice)
	for i := 0; i < value.Len(); i++ {
		//reflect.ValueOf(value.Index(i).Interface()).Elem().FieldByName(fieldName).SetString(str)
		fmt.Println(reflect.ValueOf(value.Index(i).Interface()).Elem().FieldByName(fieldName).Interface().(time.Time).Unix())
	}
}

var (
	platformCount = 2
	indexFormula  = "p%s"
	dataformula   = "p%s_id"
	innerloop     = []string{"CPU", "ntp", "傻瓜rw", "叼毛"}
	finalstring   string
	outer1        = `{
%s
}`
	outer2     = `"%s":{"alert_id":{*}}`
	outer3     = `"%s":"%s"`
	outer2Text string
)

type NameGroup struct {
	PlayerName          string
	VenueUserName       string //去掉三方场馆前缀的用户名
	PossibleMemberName1 string //可能的用户名1
	PossibleMemberName2 string //可能的用户名2
	PossibleMemberName3 string //可能的用户名3
	PossibleSiteId      int
}

func main() {
	loginLog := []string{
		"7dd57043cb405d768a1dc532f8d823b91170c349a2085b155b739573eec12d50c48c6564893b26e795847a17089eae73",
		"7dd57043cb405d768a1dc532f8d823b94c9b42075522c5feff381c38f114d340a8a010d5cfe26b6f4da95e814738c190",
		"7dd57043cb405d768a1dc532f8d823b94c9b42075522c5feff381c38f114d3400cc9e1e36b526afd0b3293b840a76f3b",
		"7dd57043cb405d768a1dc532f8d823b91170c349a2085b155b739573eec12d5064b2028976587769ff6ff24f97156925",
		"7dd57043cb405d768a1dc532f8d823b94c9b42075522c5feff381c38f114d340279755ed1113da9bffea95a731f24e9d",
		"7dd57043cb405d768a1dc532f8d823b9bee0cca3d025caa662597f5ec83499fccdd10f8acad07df3ede53c5632d78c52",
		"7dd57043cb405d768a1dc532f8d823b9bee0cca3d025caa662597f5ec83499fc86da13c6248d7eb26f2c36d9e4eae8e6",
		"7dd57043cb405d768a1dc532f8d823b9bee0cca3d025caa662597f5ec83499fc158ad2a731f7961ac025f8b2cabcce49",
		"7dd57043cb405d768a1dc532f8d823b9bee0cca3d025caa662597f5ec83499fc176d0a5213aa7f976444c626683795ff",
		"7dd57043cb405d768a1dc532f8d823b9eb682504dc65fd7af3f1cd080f2ff252d12ed6fd25aacc264395211d63abc849",
	}

	for k := range loginLog {
		token := loginLog[k]
		fmt.Println(token)
	}

}
