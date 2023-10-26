package IMTY

type ImSporRecordResp struct {
	StatusCode int      `json:"statusCode"`
	StatusDesc string   `json:"statusDesc"`
	WagerArr   []*Wager `json:"wagers"`
}

type Wager struct {
	WagerID               string       `json:"wagerId"`
	WagerCreationDateTime string       `json:"wagerCreationDateTime"`
	WagerEventDateTime    string       `json:"wagerEventDateTime"`
	MemberCode            string       `json:"memberCode"`
	CurrencyCode          string       `json:"currencyCode"`
	InputtedStakeAmount   float64      `json:"inputtedStakeAmount"`
	MemberExposure        float64      `json:"memberExposure"`
	PayoutAmount          float64      `json:"payoutAmount"`
	MemberWinLossAmount   float64      `json:"memberWinLossAmount"`
	OddsType              string       `json:"oddsType"`
	WagerType             string       `json:"wagerType"`
	BettingPlatform       string       `json:"bettingPlatform"`
	IsSettled             string       `json:"isSettled"`
	IsConfirmed           string       `json:"isConfirmed"`
	IsCancelled           string       `json:"isCancelled"`
	BetTradeStatus        string       `json:"betTradeStatus"`
	BetTradeCommission    float64      `json:"betTradeCommission"`
	BetTradeBuybackAmount float64      `json:"betTradeBuybackAmount"`
	LastUpdatedDate       string       `json:"lastUpdatedDate"`
	SettlementDate        string       `json:"settlementDate"`
	ComboType             string       `json:"comboType"`
	WagerItemArr          []*WagerItem `json:"wagerItems"`
	IsResettled           string       `json:"IsResettled"` //注意json字段是大驼峰命名
}

type WagerItem struct {
	Market               string  `json:"market"`
	EventName            string  `json:"eventName"`
	EventDateTime        string  `json:"eventDateTime"`
	CompetitionName      string  `json:"competitionName"`
	CompetitionID        string  `json:"competitionId"`
	HomeTeamName         string  `json:"homeTeamName"`
	AwayTeamName         string  `json:"awayTeamName"`
	FavTeam              string  `json:"favTeam"`
	BetType              string  `json:"betType"`
	BetTypeDesc          string  `json:"betTypeDesc"`
	Period               string  `json:"period"`
	Selection            string  `json:"Selection"`
	Odds                 float64 `json:"Odds"`
	HomeTeamHTScore      string  `json:"homeTeamHTScore"`
	AwayTeamHTScore      string  `json:"awayTeamHTScore"`
	HomeTeamFTScore      string  `json:"homeTeamFTScore"`
	AwayTeamFTScore      string  `json:"awayTeamFTScore"`
	WagerHomeTeamScore   string  `json:"wagerHomeTeamScore"`
	WagerAwayTeamScore   string  `json:"wagerAwayTeamScore"`
	Handicap             string  `json:"handicap"`
	IswagerItemCancelled string  `json:"IswagerItemCancelled"`
	SportsName           string  `json:"sportsName"`
	EventID              string  `json:"eventID"`
	Specifier            string  `json:"specifier"`
}

type ToFiller struct {
	WagerID               string `json:"wagerId"`
	WagerCreationDateTime string `json:"wagerCreationDateTime"`
	EventName             string `json:"eventName"`
	LastUpdatedDate       string `json:"lastUpdatedDate"`
	IsSettled             string `json:"isSettled"`
	IsConfirmed           string `json:"isConfirmed"`
	IsCancelled           string `json:"isCancelled"`
}
