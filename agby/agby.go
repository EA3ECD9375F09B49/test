package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type Result struct {
	XMLName  xml.Name `xml:"result"`
	Info     int      `xml:"info"`
	Rows     []Row    `xml:"row"`
	Addition Addition `xml:"addition"`
}

type Row struct {
	BillNo            string  `xml:"billno,attr"`            // 游戏订单号
	ProductID         string  `xml:"productid,attr"`         // 订单所属产品ID，即cagent
	Username          string  `xml:"username,attr"`          // 用户名
	RoomID            int     `xml:"roomid,attr"`            // 房号
	BetX              int     `xml:"betx,attr"`              // 倍率
	SceneID           string  `xml:"sceneid,attr"`           // 场景号
	FishID            int     `xml:"fishid,attr"`            // 鱼号
	FishCost          int     `xml:"fishcost,attr"`          // 鱼价值
	Hunted            int     `xml:"hunted,attr"`            // 捕获结果 (0：捕获失败，1：捕获成功)
	BillTime          int     `xml:"billtime,attr"`          // 下注时间，格式为Unix TimeStamp(整数)
	ReckonTime        int     `xml:"reckontime,attr"`        // 派彩时间，格式为Unix TimeStamp(整数)
	Currency          string  `xml:"currency,attr"`          // 币种
	Account           int     `xml:"account,attr"`           // 下注金额
	CusAccount        int     `xml:"cus_account,attr"`       // 赢输
	ValidAccount      int     `xml:"valid_account,attr"`     // 有效下注金额
	SrcAmount         float64 `xml:"src_amount,attr"`        // 下注前额度
	DstAmount         float64 `xml:"dst_amount,attr"`        // 下注后额度
	BetIP             string  `xml:"betIp,attr"`             // IP地址
	GameType          string  `xml:"gametype,attr"`          // 游戏类型
	DeviceType        int     `xml:"devicetype,attr"`        // 投注设备
	JackpotContribute int     `xml:"jackpotcontribute,attr"` // jackpot抽水
	WeaponID          string  `xml:"weaponid,attr"`          // 这条鱼被那一个武器所杀
	Flag              int     `xml:"flag,attr"`              // 订单状态 (0：异常，请联系客服；1：已派彩；-8：取消指定局注单；-9：取消指定注单)
}

type Addition struct {
	Total       int `xml:"total"`        // 总计数据，记录条数
	NumPerPage  int `xml:"num_per_page"` // 每页记录条数
	CurrentPage int `xml:"currentpage"`  // 当前页码
	TotalPage   int `xml:"totalpage"`    // 总页数
	PerPage     int `xml:"perpage"`      // 当前页记录数
}

func main() {
	xmlFile := "./xmlfile/agby.xml"
	xmlData, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		log.Fatal("Error reading XML file:", err)
	}

	var result Result
	err = xml.Unmarshal(xmlData, &result)
	if err != nil {
		log.Fatal("Error unmarshaling XML data:", err)
	}

	fmt.Printf("%+v\n", result)
}
