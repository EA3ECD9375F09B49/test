package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type Result struct {
	Info     int      `xml:"info"`
	Rows     []Row    `xml:"row"`
	Addition Addition `xml:"addition"`
}

type Row struct {
	BillNo         string `xml:"billNo,attr"`         // 游戏订单号
	PlayName       string `xml:"playName,attr"`       // 用户名
	GameCode       string `xml:"gameCode,attr"`       // 局号
	NetAmount      string `xml:"netAmount,attr"`      // 派彩额度
	BetTime        string `xml:"betTime,attr"`        // 下注时间
	BetAmount      string `xml:"betAmount,attr"`      // 投注额度
	ValidBetAmount string `xml:"validBetAmount,attr"` // 有效投注额度
	Flag           int    `xml:"flag,attr"`           // 订单状态
	// 0：异常(请联系客服)
	// 1：已派彩
	// -8：取消指定局注单
	// -9：取消指定注单
	PlayType       string `xml:"playType,attr"`        // 玩法类型，请参考附件：playType玩法类型
	Currency       string `xml:"currency,attr"`        // 币种，请参考附件：currency币种
	TableCode      string `xml:"tableCode,attr"`       // 桌台号 (此处为虚拟桌号，非实际桌号)
	RecalcuTime    string `xml:"recalcuTime,attr"`     // 派彩时间
	BeforeCredit   string `xml:"beforeCredit,attr"`    // 投注前余额
	BetIP          string `xml:"betIP,attr"`           // 投注 IP
	PlatformType   string `xml:"platformType,attr"`    // 平台类型，AGIN为视讯真人
	Remark         string `xml:"remark,attr"`          // 注示，请参考：remark备注说明
	Round          string `xml:"round,attr"`           // 平台内的大厅类型，请参考：round大厅类型
	Result         string `xml:"result,attr"`          // 此处为空
	GameType       string `xml:"gameType,attr"`        // 游戏类型，请参考：gameType游戏列表
	DeviceType     string `xml:"deviceType,attr"`      // 投注设备，请参考：deviceType投注设备
	Seat           string `xml:"Seat,attr"`            // 座位号 (只显示于gameType="BJ")
	CardIndex      string `xml:"cardindex,attr"`       // 可忽略
	Odds           string `xml:"odds,attr"`            // 可忽略
	ShoeCode       string `xml:"shoecode,attr"`        // 靴的编码，在短时间内不会重复
	ShoeRoundIndex string `xml:"shoe_roundindex,attr"` // 在同一靴下的第几局
	JpInvest       string `xml:"jp_invest,attr"`       // Jackpot投资 (只于jackpot=1时返回)
	JpPayout       string `xml:"jp_payout,attr"`       // Jackpot派彩 (只于jackpot=1时返回)
}

type Addition struct {
	Total       int `xml:"total"`
	NumPerPage  int `xml:"num_per_page"`
	CurrentPage int `xml:"currentpage"`
	TotalPage   int `xml:"totalpage"`
	PerPage     int `xml:"perpage"`
}

func main() {
	xmlFile := "./xmlfile/agzr.xml"
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
