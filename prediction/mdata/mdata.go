package mdata

import (
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	SPORT        = "1"
	IMSPORT      = "2"
	SB           = "3"
	MysqlPrefix  = "video_"
	Cjson        = jsoniter.ConfigCompatibleWithStandardLibrary
	ValidSportId = []string{"1", "2"}
	RwLock       sync.RWMutex

	TypeMap = map[string]int{
		"188":  5,
		"im":   11,
		"ob":   19,
		"saba": 12,
		"obdj": 50,
		"imdj": 51,
	}

	PositionMap = map[string]string{
		"0":  "-",
		"1":  "门卫",
		"2":  "后卫",
		"3":  "后卫",
		"4":  "后卫",
		"5":  "中场",
		"6":  "中场",
		"7":  "中场",
		"8":  "边卫",
		"9":  "前锋",
		"10": "前锋",
		"11": "边卫",
	}

	ApiSbMap = map[string]string{
		"api":      "http://oddsfeed.jiouoiufs.com/api/GetEventsAndMarkets",
		"api_sp":   "http://oddsfeed.jiouoiufs.com/api/GetEventInfo",
		"apiOld":   "http://oddsfeed.jiouoiufs.com/api/GetEventsAndMarkets",
		"vendorId": "UUrcQFKOOh0",
	}

	SportTypeMap = map[uint]string{
		1:  "美式足球",
		2:  "篮球",
		3:  "足球",
		4:  "冰上曲棍球",
		5:  "网球",
		6:  "排球",
		7:  "台球",
		8:  "棒球",
		9:  "羽毛球",
		10: "高尔夫",
		11: "赛车",
		16: "拳击",
		24: "手球",
		26: "橄榄球",
		43: "电子竞技",
		50: "板球",
	}

	CateData = map[string]string{
		"Football":   "足球",
		"Basketball": "篮球",
		"Tennis":     "网球",
		"Baseball":   "棒球",
		"Volleyball": "排球",
		"Snooker":    "斯诺克",
		"Esport":     "电子竞技",
	}

	CateDataList = map[string]int{
		"足球":  1,
		"篮球":  1,
		"网球":  1,
		"棒球":  1,
		"排球":  1,
		"斯诺克": 1,
		"台球":  1,
		"其他":  1,
	}
	EsportsMap = []string{
		"Counter strike: the global offensive",
		"League of Legends",
		"Honor of Kings",
		"DOTA2",
		"Hearthstone legend: heroes of warcraft",
		"Heroes of the Storm",
		"Overwatch",
		"StarCraft",
		"StarCraft 2",
		"Rocket League",
		"World of Warcraft 3",
	}
	ImCateMap = map[string]string{
		"1":  "足球",
		"2":  "篮球",
		"3":  "网球",
		"6":  "田径",
		"7":  "羽毛球",
		"8":  "棒球",
		"11": "拳击",
		"15": "飞镖",
		"18": "草地曲棍球",
		"19": "美式足球",
		"21": "高尔夫球",
		"23": "手球",
		"25": "冰上曲棍球",
		"29": "赛车运动",
		"31": "橄榄球",
		"32": "帆船",
		"34": "斯诺克/英式台球",
		"36": "乒乓球",
		"39": "虚拟足球",
		"40": "排球",
		"41": "水球",
		"43": "虚拟篮球",
		"44": "虚拟世界杯",
		"45": "娱乐投注 ",
		"46": "虚拟国家杯",
		"47": "虚拟足球英国联赛",
		"49": "虚拟足球西班牙友谊赛",
		"52": "虚拟足球西班牙联赛",
		"53": "虚拟足球意大利联赛",
	}
)

type GameConfig struct {
}

type Config struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
}

type Data100InfoOrig struct {
	EventIdInt       uint64 `json:"event_id"`
	EventId          string `json:"even_id_str"`
	CompetitionName  string `json:"competitionName"`
	HomeTeam         string `json:"team1"`
	AwayTeam         string `json:"team2"`
	Flv              string `json:"flv"`
	M3u8             string `json:"m3u8"`
	Ani              string `json:"ani"`
	EventDate        string `json:"eventDate"`
	HomeTeamLogoUrl  string `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl string `json:"guestTeamLogoUrl"`
	Category         string `json:"category"`
	Status           int    `json:"status"`
}

//==========================以下是188最终放进redis的结构
type T188Data struct {
	Video            interface{} `json:"video"`
	Ani              *T188Ani    `json:"ani"`
	EventId          interface{} `json:"EventId"`
	CompetitionName  string      `json:"CompetitionName"`
	HomeTeam         string      `json:"HomeTeam"`
	AwayTeam         string      `json:"AwayTeam"`
	EventDate        string      `json:"EventDate"`
	Category         string      `json:"category"`
	HomeTeamLogoUrl  string      `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl string      `json:"guestTeamLogoUrl"`
}

type Map188Video map[string]*T188Video

type T188Video struct {
	Flv  string `json:"flv"`
	M3u8 string `json:"m3u8"`
}

type T188Ani struct {
	One   interface{} `json:"one"`
	Two   interface{} `json:"two"`
	Three interface{} `json:"three"`
}

//公用视频源数据结构(2,3)
type CommonVideoInfo struct {
	EventId          uint64 `json:"eventid"`
	EventIdStr       string `json:"eventidStr"`
	M3u8             string `json:"m3u8"`
	Flv              string `json:"flv"`
	AnimationUrl     string `json:"animationUrl"`
	Line             string `json:"line"`
	TypeName         string `json:"typeName"`
	IsBob            bool   `json:"isBob"`
	GuestTeamLogoURL string `json:"guestTeamLogoUrl"`
	HomeTeamLogoURL  string `json:"homeTeamLogoUrl"`
	EventDate        string `json:"eventDate"`
	Status           int    `json:"status"`
	Team1            string `json:"team1"`
	Team2            string `json:"team2"`
	BrEventId        uint64 `json:"rowEventId"`
}

type Data100HttpData struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	FlvUrl  string                 `json:"flv_url"`
	M3u8Url string                 `json:"m3u8_url"`
	Data    []*Data100HttpDataOrig `json:"data"`
}

type Data100HttpDataOrig struct {
	Category string `json:"category"`
	Bid      string `json:"bid"`
	Name     string `json:"name"`
	Team1    string `json:"team1"`
	Team2    string `json:"team2"`
	Start    string `json:"start"`
	StreamId string `json:"stream_id"`
	Hd       int64  `json:"hd"`
	HlsId    string `json:"hls_id"`
	KeepFlv  string
	KeepM3u8 string
}

//188second
type Data188ThirdHttpData struct {
	Data  []*Data188ThirdHttpDataOrig `json:"data"`
	Ttype bool
}

type Data188ThirdHttpDataOrig struct {
	T188Id             string      `json:"188_id"`
	StreamId           interface{} `json:"stream_id"`
	AniId              interface{} `json:"ani_id"`
	StartDate          string      `json:"start_date"`
	Cate               string      `json:"cate"`
	League             string      `json:"league"`
	Team1              string      `json:"team1"`
	Team2              string      `json:"team2"`
	LeagueEn           string      `json:"leagueen"`
	Team1En            string      `json:"team1en"`
	Team2En            string      `json:"team2en"`
	HomeTeamLogoUrl    string      `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl   string      `json:"guestTeamLogoUrl"`
	HomeTeamYellowCard string      `json:"homeTeamYellowCard"`
	AwayTeamYellowCard string      `json:"awayTeamYellowCard"`
	HomeTeamRedCard    string      `json:"homeTeamRedCard"`
	AwayTeamRedCard    string      `json:"awayTeamRedCard"`
	Status             interface{} `json:"status"`
	AnimationUrl       string      `json:"animationUrl"`
}

//188second
type Data188SecondHttpData struct {
	Code    int                          `json:"code"`
	Message string                       `json:"message"`
	Url     string                       `json:"url"`
	Data    []*Data188SecondHttpDataOrig `json:"data"`
}

type Data188SecondHttpDataOrig struct {
	Vid       string `json:"vid"`
	Type      string `json:"type"`
	TypeEn    string `json:"typeEn"`
	MatchNo   uint64 `json:"match_no"`
	Start     uint64 `json:"start"`
	League    string `json:"league"`
	LeagueEn  string `json:"leagueEn"`
	Home      string `json:"home"`
	HomeEn    string `json:"homeEn"`
	Away      string `json:"away"`
	AwayEn    string `json:"awayEn"`
	FlvUrl    string `json:"flv_url"`
	M3u8Url   string `json:"m3u8_url"`
	StartDate string `json:"start_date"`
}

type VideoLineData struct {
	PathFlv   string `json:"path_flv"`
	PathM3U8  string `json:"path_m3u8"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Num       int    `json:"num"`
	StartDate string `json:"startDate"`
}

type DataSbEventsHttpData struct {
	StatusCode int                                    `json:"status_code"`
	Message    string                                 `json:"message"`
	Data       map[string][]*DataSbEventsHttpDataOrig `json:"data"`
}

type DataSbEventsHttpDataOrig struct {
	EventId           uint64
	EventDate         string
	ZhCompetitionName string
	CompetitionName   string
	ZhHomeTeam        string
	ZhAwayTeam        string
	HomeTeam          string
	AwayTeam          string
	Md5Str            string
	SportName         string
	SportNameLocal    string
}

//播控返的数据结构
type BKvideosData struct {
	Cate             string `json:"cate"`
	Eid              string `json:"eid"`
	GuestTeamLogoURL string `json:"guestTeamLogoUrl"`
	HomeTeamLogoURL  string `json:"homeTeamLogoUrl"`
	League           string `json:"league"`
	MatchID          string `json:"match_id"`
	StartDate        string `json:"start_date"`
	Team1            string `json:"team1"`
	Team1En          string `json:"team1en"`
	Team2            string `json:"team2"`
	Team2En          string `json:"team2en"`
	AniID            string `json:"ani_id,omitempty"`
	Animation        struct {
		Path string `json:"path"`
	} `json:"animation,omitempty"`
	Animation3 []struct {
		Path      string `json:"path"`
		StyleName string `json:"style_name"`
	} `json:"animation3,omitempty"`
	LiveVideo struct { //直播视频信息
		HD       string `json:"HD"`
		Online   int    `json:"online"`
		PathFlv  string `json:"path_flv"`
		PathM3U8 string `json:"path_m3u8"`
		Status   int    `json:"status"`
	} `json:"liveVideo"`
	AnchorVideo     []*AnchorVideoData `json:"anchorVideo"` //主播视频列表(数组)
	VideoLines      []*VideoLineData   `json:"video_lines"`
	TournamentLogo  string             `json:"tournament_logo"`
	LeagueLogoUrl   string             `json:"league_logo_url"`
	LiveStatus      int                `json:"liveStatus"`
	RowEventId      int                `json:"rowEventId"`               //im专用可视化投注预测赛事ID
	MatchStatus     int                `json:"match_status"`             //比赛状态：1未开始 2进行中 3已结束4已取消
	MatchStatusCode int                `json:"match_status_code,string"` // 0 未开始，1 上半场，2 下半场。11 第一节，，12， 第二节，13，第三节，14 第四节。
}

type AnchorVideoData struct {
	HD         string      `json:"HD"`
	Anchor     *AnchorInfo `json:"anchor"`
	Online     string      `json:"online"`
	PathFlv    string      `json:"path_flv"`
	PathM3U8   string      `json:"path_m3u8"`
	Screenshot string      `json:"screenshot"`
	Status     string      `json:"status"`
	StartDate  string      `json:"start_date"`
}

type AnchorInfo struct {
	Grade           string `json:"grade"`
	ID              string `json:"id"`
	LogoRectangle   string `json:"logo_rectangle"`
	LogoSquare      string `json:"logo_square"`
	Nickname        string `json:"nickname"`
	PersonalProfile string `json:"personal_profile"`
	RecomTimes      int    `json:"recomTimes"`
	RecomTime       int64  `json:"recomTime"`
	Sex             string `json:"sex"`
	Labels          string `json:"labels"`
}

//==========================188End
type Category struct {
	SportId      uint32
	SportName    string
	EarlyFECount uint32
	TodayFECount uint32
	RBFECount    uint32
}

type ImSportListV2En struct {
	EventId           uint64
	EventGroupId      int
	EnCompetitionName string
	EnHomeTeam        string
	EnAwayTeam        string
	Market            int
	SportId           uint32
}

type ImSportListV2 struct {
	EventId          uint64
	EventGroupId     int
	CompetitionName  string
	EventStatusId    int
	EventDate        string
	HomeTeam         string
	AwayTeam         string
	Market           int
	HasVisualization bool
	BREventId        int
	LiveStreamingUrl interface{}
	SportId          uint32
	RowEventId       int `json:"rowEventId"`
}

//下面是获取赛事
type ImSportV2EVENTINFO struct {
	PageSize   int
	StatusCode int64
	StatusDesc string
	Sports     []*ImSportV2Sport
}

type ImSportV2Sport struct {
	SportId     int64
	SportName   string
	OrderNumber int
	Events      []*ImSportV2Event
}

type ImSportV2Event struct {
	EventGroupId     int
	HasVisualization bool
	BREventId        int
	EventStatusId    int
	OrderNumber      int
	EventId          uint64
	HomeTeam         string
	AwayTeam         string
	EventDate        string
	LiveStreamingUrl interface{}
	Competition      *ImSportV2Competition
}

type ImSportV2Competition struct {
	CompetitionId   uint64
	CompetitionName string
	PMOrderNumber   uint64
	RBOrderNumber   uint64
}

type SabaEnData struct {
	CompetitionName string
	HomeTeam        string
	AwayTeam        string
}

//下面是抓取sb视频
type DataSbOrig struct {
	ErrorCode int `json:"error_code"`
	Data      map[string][]*SbMatchOrig
}

type SbMatchOrig struct {
	SportType  uint
	SportName  string
	LeagueId   uint64
	LeagueName string               `json:"LeagueName"`
	Matches    []*SbMatchDetailOrig `json:"matches"`
}

type SbMatchDetailOrig struct {
	HomeName string `json:"HomeName"`
	AwayTeam string `json:"AwayName"`
	ShowTime string `json:"ShowTime"`
	MatchId  uint64 `json:"MatchId"`
}

type DataSbEventsHttpDataOrigAlone struct {
	EventId           uint64
	EventDate         string
	ZhCompetitionName string
	CompetitionName   string
	ZhHomeTeam        string
	ZhAwayTeam        string
	HomeTeam          string
	AwayTeam          string
	StreamingUrlCn    string
	StreamingUrlNoCn  string
	SportName         string
	Md5Str            string
	Eventstatus       string
}

//下面是抓取sb视频
type DataSbOrigAlone struct {
	ErrorCode int `json:"error_code"`
	Data      map[string][]*SbMatchOrigAlone
}

type SbMatchOrigAlone struct {
	SportType  uint                      `json:"sport_type"`
	LeagueId   uint64                    `json:"league_id"`
	LeagueName string                    `json:"league_name"`
	Matches    []*SbMatchDetailOrigAlone `json:"matches"`
}

type SbMatchDetailOrigAlone struct {
	MatchId          uint64 `json:"match_id"`
	HomeName         string `json:"home_team_name"`
	AwayTeam         string `json:"away_team_name"`
	ShowTime         string `json:"show_time"`
	KickoffTime      string `json:"kickoff_time"`
	EventStatus      string `json:"event_status"`
	StreamingUrlCn   string `json:"streaming_url_cn"`
	StreamingUrlNoCn string `json:"streaming_url_non_cn"`
}

type SportCate struct {
	SportCount []*CateSportCount
	ServerTime string
	StatusCode int
	StatusDesc string
}

type CateSportCount struct {
	SportId         uint32
	SportName       string
	OrderNumber     int
	IsCombo         bool
	IsHasLive       bool
	EarlyFECount    uint32
	TodayFECount    uint32
	ORCount         uint32
	RBFECount       uint32
	Count           uint32
	EventGroupTypes []*CateEventGroupType
}

type CateEventGroupType struct {
	EventGroupTypeId uint32
	Count            uint32
	IsHasLive        bool
	EarlyFECount     uint32
	TodayFECount     uint32
	ORCount          uint32
	RBFECount        uint32
}

type DataFirstVideoHttpData struct {
	Code    int                           `json:"code"`
	Message string                        `json:"message"`
	FlvUrl  string                        `json:"flv_url"`
	M3u8Url string                        `json:"m3u8_url"`
	Data    []*DataFirstVideoHttpDataOrig `json:"data"`
}

type DataFirstVideoHttpDataOrig struct {
	Category string `json:"category"`
	Bid      string `json:"bid"`
	Name     string `json:"name"`
	Team1    string `json:"team1"`
	Team2    string `json:"team2"`
	Start    string `json:"start"`
	StreamId string `json:"stream_id"`
	HlsId    string `json:"hls_id"`
	KeepFlv  string `json:"flv"` //后期组装的数据
	KeepM3u8 string `json:"m3u8"`
	FlvHls   string `json:"flv_hls"`
	M3u8Hls  string `json:"m3u8_hls"`
}

//im second
type DataImSecondHttpDataNew struct {
	Code    int                            `json:"code"`
	Message string                         `json:"message"`
	Data    []*DataImSecondHttpDataOrigNew `json:"data"`
	Url     string                         `json:"url"`
}

type DataImSecondHttpDataOrigNew struct {
	Vid               string `json:"vid"`
	Type              string `json:"type"`
	TypeEn            string `json:"typeEn"`
	EventId           uint64 `json:"EventId"`
	CompetitionName   string `json:"CompetitionName"`
	EnCompetitionName string `json:"EnCompetitionName"`
	EventDate         string `json:"EventDate"`
	HomeTeam          string `json:"HomeTeam"`
	EnHomeTeam        string `json:"EnHomeTeam"`
	AwayTeam          string `json:"AwayTeam"`
	EnAwayTeam        string `json:"EnAwayTeam"`
	StartDate         string `json:"start_date"`
	FlvUrl            string `json:"flv_url"`
	M3u8Url           string `json:"m3u8_url"`
	Ani               string `json:"ani"`
	RowEventId        uint64 `json:"RowEventId"`
	//下面是和三方商议的新字段，不用再循环去拿了
	Flv    string `json:"flv_im"`     //标清
	FlvLd  string `json:"flv_im_ld"`  //高清
	M3u8   string `json:"m3u8_im"`    //标清
	M3u8Ld string `json:"m3u8_im_ld"` //高清
}

//im second
type DataImSecondAniOrigHttp struct {
	Code    int                          `json:"code"`
	Message string                       `json:"message"`
	Data    []*DataImSecondAniOrigDetail `json:"data"`
	Url     string                       `json:"url"`
}

type DataImSecondAniOrigDetail struct {
	Id           int                            `json:"id"`
	SportId      int                            `json:"sports_id"`
	SportName    string                         `json:"sports_name"`
	SportsNameZh string                         `json:"sports_name_zh"`
	SportsNameEn string                         `json:"sports_name_en"`
	StartDate    string                         `json:"start_date"`
	AniDetail    *DataImSecondAniOrigDetailMore `json:"im"`
	AniUrl       string                         `json:"ani_url"`
}

type DataImSecondAniOrigDetailMore struct {
	EventId           uint64 `json:"EventId"`
	CompetitionName   string `json:"CompetitionName"`
	EnCompetitionName string `json:"EnCompetitionName"`
	EventDate         string `json:"EventDate"`
	HomeTeam          string `json:"HomeTeam"`
	EnHomeTeam        string `json:"EnHomeTeam"`
	AwayTeam          string `json:"AwayTeam"`
	EnAwayTeam        string `json:"EnAwayTeam"`
	RowEventId        uint64 `json:"RowEventId"`
}

//sb second
type DataSbSecondAniOrigHttp struct {
	Code    int                          `json:"code"`
	Message string                       `json:"message"`
	Data    []*DataSbSecondAniOrigDetail `json:"data"`
	Url     string                       `json:"url"`
}

type DataSbSecondAniOrigDetail struct {
	Id           int                            `json:"id"`
	SportId      int                            `json:"sports_id"`
	SportName    string                         `json:"sports_name"`
	SportsNameZh string                         `json:"sports_name_zh"`
	SportsNameEn string                         `json:"sports_name_en"`
	StartDate    string                         `json:"start_date"`
	AniDetail    *DataImSecondAniOrigDetailMore `json:"sba"`
	AniUrl       string                         `json:"ani_url"`
}

type DataSbSecondAniOrigDetailMore struct {
	EventId           uint64 `json:"EventId"`
	CompetitionName   string `json:"CompetitionName"`
	EnCompetitionName string `json:"EnCompetitionName"`
	EventDate         string `json:"EventDate"`
	HomeTeam          string `json:"HomeTeam"`
	EnHomeTeam        string `json:"EnHomeTeam"`
	AwayTeam          string `json:"AwayTeam"`
	EnAwayTeam        string `json:"EnAwayTeam"`
}

type DataThirdImHttpData struct {
	Data []*DataThirdImHttpDataOrig `json:"data"`
}

type DataThirdImHttpDataOrig struct {
	ImId               string      `json:"im_id"`
	StreamId           string      `json:"stream_id"`
	StartDate          string      `json:"start_date"`
	AniId              interface{} `json:"ani_id"`
	Cate               string      `json:"cate"`
	League             string      `json:"league"`
	Team1              string      `json:"team1"`
	Team2              string      `json:"team2"`
	Leagueen           string      `json:"leagueen"`
	Team1en            string      `json:"team1en"`
	Team2en            string      `json:"team2en"`
	HomeTeamLogoUrl    string      `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl   string      `json:"guestTeamLogoUrl"`
	HomeTeamYellowCard string      `json:"homeTeamYellowCard"`
	AwayTeamYellowCard string      `json:"awayTeamYellowCard"`
	HomeTeamRedCard    string      `json:"homeTeamRedCard"`
	AwayTeamRedCard    string      `json:"awayTeamRedCard"`
	Status             interface{} `json:"status"`
	AnimationUrl       string      `json:"animationUrl"`
}

type DataThirdSbHttpData struct {
	Data []*DataThirdSbHttpDataOrig `json:"data"`
}

type DataThirdSbHttpDataOrig struct {
	SbId               string      `json:"sb_id"`
	StreamId           string      `json:"stream_id"`
	AniId              interface{} `json:"ani_id"`
	StartDate          string      `json:"start_date"`
	Cate               string      `json:"cate"`
	League             string      `json:"league"`
	Team1              string      `json:"team1"`
	Team2              string      `json:"team2"`
	Leagueen           string      `json:"leagueen"`
	Team1en            string      `json:"team1en"`
	Team2en            string      `json:"team2en"`
	HomeTeamLogoUrl    string      `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl   string      `json:"guestTeamLogoUrl"`
	HomeTeamYellowCard string      `json:"homeTeamYellowCard"`
	AwayTeamYellowCard string      `json:"awayTeamYellowCard"`
	HomeTeamRedCard    string      `json:"homeTeamRedCard"`
	AwayTeamRedCard    string      `json:"awayTeamRedCard"`
	Status             interface{} `json:"status"`
	AnimationUrl       string      `json:"animationUrl"`
}

type AutoGenerated struct {
	Cate             string `json:"cate"`
	Eid              string `json:"eid"`
	GuestTeamLogoURL string `json:"guestTeamLogoUrl"`
	HomeTeamLogoURL  string `json:"homeTeamLogoUrl"`
	LeagueLogoUrl    string `json:"league_logo_url"`
	League           string `json:"league"`
	MatchID          string `json:"match_id"`
	StartDate        string `json:"start_date"`
	Team1            string `json:"team1"`
	Team1En          string `json:"team1en"`
	Team2            string `json:"team2"`
	Team2En          string `json:"team2en"`
	AniID            string `json:"ani_id,omitempty"`
	Animation        struct {
		Path string `json:"path"`
	} `json:"animation,omitempty"`
	Animation3 []struct {
		Path      string `json:"path"`
		StyleName string `json:"style_name"`
	} `json:"animation3,omitempty"`
	LiveVideo struct {
		HD       string `json:"HD"`
		Online   int    `json:"online"`
		PathFlv  string `json:"path_flv"`
		PathM3U8 string `json:"path_m3u8"`
		Status   int    `json:"status"`
	} `json:"liveVideo"`
	AnchorVideo    []*AnchorVideoData `json:"anchorVideo"`
	TournamentLogo string             `json:"tournament_logo"`
}

type VideoLiveAdminData struct {
	Eid     string `json:"eid"`
	MatchID string `json:"match_id"`
	TypeTxt string `json:"type"`
}

type TeamsNamesZh struct {
	HomeTeamEn string `json:"homeTeamEn"`
	AwayTeamEn string `json:"awayTeamEn"`
	HomeTeam   string `json:"homeTeam"`
	AwayTeam   string `json:"awayTeam"`
	HomeTeamId string `json:"homeId"`
	AwayTeamId string `json:"awayId"`
}

type MatchLogo struct {
	Cate             string `json:"cate"`
	Eid              string `json:"eid"`
	GuestTeamLogoURL string `json:"guestTeamLogoUrl"`
	HomeTeamLogoURL  string `json:"homeTeamLogoUrl"`
	League           string `json:"league"`
	MatchID          string `json:"match_id"`
	Team1            string `json:"team1"`
	Team1En          string `json:"team1en"`
	Team2            string `json:"team2"`
	Team2En          string `json:"team2en"`
}

type AnchorVideoDataAdmin struct {
	Anchor           *AnchorInfo `json:"anchor"`
	Online           string      `json:"online"`
	Status           string      `json:"status"`
	Cate             string      `json:"cate"`
	Eid              string      `json:"eid"`
	League           string      `json:"league"`
	MatchID          string      `json:"match_id"`
	StartDate        string      `json:"start_date"`
	Team1            string      `json:"team1"`
	Team1En          string      `json:"team1en"`
	Team2            string      `json:"team2"`
	Team2En          string      `json:"team2en"`
	Type             string      `json:"type"`
	Recommand        string      `json:"recomStatus"`
	LiveStatus       int         `json:"liveStatus"`
	LiveOnline       int         `json:"liveOnline"`
	Flv              string      `json:"flv"`
	M3u8             string      `json:"m3u8"`
	HomeTeamLogoUrl  string      `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl string      `json:"guestTeamLogoUrl"`
	MatchStatus      int         `json:"matchStatus"`     //比赛状态：1未开始 2进行中 3已结束4已取消
	MatchStatusCode  int         `json:"matchStatusCode"` // 0 未开始，1 上半场，2 下半场。11 第一节，，12， 第二节，13，第三节，14 第四节。
}

type ThirdMatch struct {
	AwayTeamID   string `json:"away_team_id"`
	AwayTeamLogo string `json:"away_team_logo"`
	AwayTeamName string `json:"away_team_name"`
	HomeTeamID   string `json:"home_team_id"`
	HomeTeamLogo string `json:"home_team_logo,omitempty"`
	HomeTeamName string `json:"home_team_name"`
	LmtMode      string `json:"lmt_mode"`
	MatchID      string `json:"match_id"`
	Reverse      string `json:"reverse"`
	SportID      string `json:"sport_id"`
	//Stream       []interface{} `json:"stream"`
	ThirdMatchID string `json:"third_match_id"`
	IsHan        bool   `json:"is_han"`
}

type LatestMatchJason struct {
	Createtime string `json:"createtime"`
	ID         int    `json:"id"`
	Path       string `json:"path"`
}

type ThirdMatchLiveScoreData struct {
	AggregateWinner         string          `json:"aggregate_winner,omitempty"`
	AwayFormation           string          `json:"away_formation"`
	AwayTeamHalfTimeScore   string          `json:"away_team_half_time_score"`
	AwayTeamID              string          `json:"away_team_id"`
	AwayTeamName            string          `json:"away_team_name"`
	AwayTeamNormalTimeScore string          `json:"away_team_normal_time_score"`
	AwayTeamScore           string          `json:"away_team_score"`
	Distance                string          `json:"distance"`
	EndTime                 string          `json:"end_time"`
	Events                  []*StaticEvents `json:"events"`
	Phrase                  []*StaticPhrase `json:"phrase"`
	GroupID                 string          `json:"group_id"`
	HasInjury               string          `json:"has_injury"`
	HasNews                 string          `json:"has_news"`
	HomeFormation           string          `json:"home_formation"`
	HomeTeamHalfTimeScore   string          `json:"home_team_half_time_score"`
	HomeTeamID              string          `json:"home_team_id"`
	HomeTeamName            string          `json:"home_team_name"`
	HomeTeamNormalTimeScore string          `json:"home_team_normal_time_score"`
	HomeTeamScore           string          `json:"home_team_score"`
	ID                      string          `json:"id"`
	IsVisible               string          `json:"is_visible"`
	Lineup                  []struct {
		PlayerID     string `json:"player_id"`
		PlayerName   string `json:"player_name"`
		PlayerPicURL string `json:"player_pic_url,omitempty"`
		Position     string `json:"position"`
		ShirtNumber  string `json:"shirt_number"`
		Substitute   string `json:"substitute"`
		TeamID       string `json:"team_id"`
	} `json:"lineup"`
	LmtMode         string           `json:"lmt_mode"`
	MatchTime       string           `json:"match_time"`
	Neutral         string           `json:"neutral"`
	PlayerTechnic   []interface{}    `json:"player_technic"`
	PreviousMatchID string           `json:"previous_match_id,omitempty"`
	Round           string           `json:"round"`
	RoundType       string           `json:"round_type"`
	Scores          []*HistoryScores `json:"scores"`
	Season          string           `json:"season"`
	SeasonID        string           `json:"season_id"`
	SportID         string           `json:"sport_id"`
	StadiumCnName   string           `json:"stadium_cn_name,omitempty"`
	StadiumEnName   string           `json:"stadium_en_name,omitempty"`
	StadiumID       string           `json:"stadium_id,omitempty"`
	Statics         []*Mstatic       `json:"statics"`
	StaticsEx       []*Mstatic       `json:"statics_ex"`
	Status          string           `json:"status"`
	StatusCode      string           `json:"status_code"`
	StatusName      string           `json:"status_name"`
	TimePlayed      string           `json:"time_played"`
	TimeRemaining   string           `json:"time_remaining"`
	TimeRunning     string           `json:"time_running"`
	TimeUpdate      string           `json:"time_update"`
	TournamentID    string           `json:"tournament_id"`
	UpdateTimestamp string           `json:"update_timestamp"`
	Weather         string           `json:"weather"`
	WeatherDesc     string           `json:"weather_desc"`
	Winner          string           `json:"winner"`
	RefereeCnName   string           `json:"referee_cn_name,omitempty"`
	RefereeEnName   string           `json:"referee_en_name,omitempty"`
	RefereeID       string           `json:"referee_id,omitempty"`
}

type HistoryScores struct {
	MatchID    string `json:"match_id"`
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	Type       string `json:"type"`
	UpdateTime string `json:"update_time"`
}

type Mstatic struct {
	Period     string `json:"period"`
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	TypeCnName string `json:"type_cn_name"`
	TypeEnName string `json:"type_en_name"`
	TypeID     string `json:"type_id"`
}

type StaticPhrase struct {
	ID      string `json:"id"`
	Team    string `json:"team"`
	Time    string `json:"time"`
	CnText  string `json:"cn_text"`
	EventId string `json:"event_id"`
	Period  string `json:"period"`
	score   string `json:"score"`
}

type StaticEvents struct {
	GoalType     string `json:"goal_type"`
	ID           string `json:"id"`
	InjuryTime   string `json:"injury_time"`
	Sort         string `json:"sort"`
	Team         string `json:"team"`
	Time         string `json:"time"`
	TypeID       string `json:"type_id"`
	TypeName     string `json:"type_name"`
	PlayerID     string `json:"player_id,omitempty"`
	PlayerName   string `json:"player_name,omitempty"`
	PlayerPicURL string `json:"player_pic_url,omitempty"`
	Scores       string `json:"scores,omitempty"`
	X            string `json:"x,omitempty"`
	Y            string `json:"y,omitempty"`
	PlayerName2  string `json:"player_name2,omitempty"`
}

//文字直播
type LiveScore struct {
	CnText  string      `json:"cn_text"`
	ID      interface{} `json:"id"`
	MatchID string      `json:"match_id"`
	Period  string      `json:"period"`
	Score   string      `json:"score"`
	SportID string      `json:"sport_id"`
	Team    string      `json:"team"`
	Time    string      `json:"time"`
}

type PhrasePathData struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
}
type HistoryMatchs struct {
	AwayFormation           string                  `json:"away_formation"`
	HomeFormation           string                  `json:"home_formation"`
	AwayTeamHalfTimeScore   string                  `json:"away_team_half_time_score,omitempty"`
	AwayTeamID              string                  `json:"away_team_id"`
	AwayTeamName            string                  `json:"away_team_name"`
	AwayTeamNormalTimeScore string                  `json:"away_team_normal_time_score,omitempty"`
	AwayTeamScore           string                  `json:"away_team_score,omitempty"`
	EndTime                 string                  `json:"end_time,omitempty"`
	Events                  []*HistoryEvents        `json:"events"`
	GroupID                 string                  `json:"group_id"`
	HasEvent                string                  `json:"has_event"`
	HasEventphase           string                  `json:"has_eventphase"`
	HasInjury               string                  `json:"has_injury"`
	HasLineup               string                  `json:"has_lineup"`
	HasLiveodds             string                  `json:"has_liveodds"`
	HasNews                 string                  `json:"has_news"`
	HasPhase                string                  `json:"has_phase"`
	HasPlayerStatistics     string                  `json:"has_player_statistics"`
	HasPreodds              string                  `json:"has_preodds"`
	HasStatistics           string                  `json:"has_statistics"`
	HomeTeamHalfTimeScore   string                  `json:"home_team_half_time_score,omitempty"`
	HomeTeamID              string                  `json:"home_team_id"`
	HomeTeamName            string                  `json:"home_team_name"`
	HomeTeamNormalTimeScore string                  `json:"home_team_normal_time_score,omitempty"`
	HomeTeamScore           string                  `json:"home_team_score,omitempty"`
	ID                      string                  `json:"id"`
	IsVisible               string                  `json:"is_visible"`
	Lineup                  []*HistoryMatchesLineUp `json:"lineup"`
	LmtMode                 string                  `json:"lmt_mode"`
	MatchTime               string                  `json:"match_time"`
	Neutral                 string                  `json:"neutral"`
	//Phrases                 []interface{}     `json:"phrases"`
	PlayerStatics []*HistoryPlayerSeasonStatics `json:"player_statics"`
	Round         string                        `json:"round"`
	RoundType     string                        `json:"round_type"`
	Scores        []*HistoryScores              `json:"scores"`
	Season        string                        `json:"season"`
	SeasonID      string                        `json:"season_id"`
	SportID       string                        `json:"sport_id"`
	StadiumCnName string                        `json:"stadium_cn_name,omitempty"`
	StadiumEnName string                        `json:"stadium_en_name,omitempty"`
	StadiumID     string                        `json:"stadium_id,omitempty"`
	Statics       []*HistoryStatics             `json:"statics"`
	//StaticsEx       []interface{}     `json:"statics_ex"`
	Status          string `json:"status"`
	StatusCode      string `json:"status_code"`
	StatusName      string `json:"status_name"`
	TimePlayed      string `json:"time_played,omitempty"`
	TimeRemaining   string `json:"time_remaining,omitempty"`
	TimeRunning     string `json:"time_running,omitempty"`
	TimeUpdate      string `json:"time_update,omitempty"`
	TournamentID    string `json:"tournament_id"`
	UpdateTimestamp string `json:"update_timestamp"`
	Weather         string `json:"weather,omitempty"`
	WeatherDesc     string `json:"weather_desc,omitempty"`
	Winner          string `json:"winner"`
	Distance        string `json:"distance,omitempty"`
	RefereeCnName   string `json:"referee_cn_name,omitempty"`
	RefereeEnName   string `json:"referee_en_name,omitempty"`
	RefereeID       string `json:"referee_id,omitempty"`
}

type HistoryEvents struct {
	ID         string `json:"id"`
	PlayerID   string `json:"player_id"`
	PlayerID2  string `json:"player_id2"`
	Sort       string `json:"sort"`
	Team       string `json:"team"`
	Time       string `json:"time"`
	TypeID     string `json:"type_id"`
	TypeName   string `json:"type_name"`
	GoalType   string `json:"goal_type,omitempty"`
	Scores     string `json:"scores,omitempty"`
	InjuryTime string `json:"injury_time,omitempty"`
}

type HistoryMatchesLineUp struct {
	PlayerID     string `json:"player_id"`
	PlayerName   string `json:"player_name"`
	PlayerPicURL string `json:"player_pic_url"`
	Position     string `json:"position"`
	ShirtNumber  string `json:"shirt_number"`
	Substitute   string `json:"substitute"`
	TeamID       string `json:"team_id"`
}

type HistoryPlayerSeasonStatics struct {
	PlayerId string `json:"player_id,omitempty"`
	HistorySeasonStatics
}

type HistorySeasonStatics struct {
	SeasonId     string `json:"season_id,omitempty"`
	TourNamentId string `json:"tournament_id,omitempty"`
	TeamId       string `json:"team_id,omitempty"`
	TypeId       string `json:"type_id,omitempty"`
	TypeValue    string `json:"type_value,omitempty"`
	Range        string `json:"range,omitempty"`
}

type HistoryStatics struct {
	Period     string `json:"period"`
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	TypeCnName string `json:"type_cn_name"`
	TypeEnName string `json:"type_en_name"`
	TypeID     string `json:"type_id"`
}

//赛事新闻
type MatchNews struct {
	/*	AwayCoach struct {
			Association     string `json:"association"`
			AssociationLogo string `json:"association_logo"`
			Birthdate       string `json:"birthdate"`
			EnName          string `json:"en_name"`
			GameCount       string `json:"game_count"`
			ID              string `json:"id"`
			Logo            string `json:"logo"`
			Score           string `json:"score"`
		} `json:"away_coach"`
		AwayFormation           string `json:"away_formation"`
		AwayTeamHalfTimeScore   string `json:"away_team_half_time_score"`
		AwayTeamID              string `json:"away_team_id"`
		AwayTeamNormalTimeScore string `json:"away_team_normal_time_score"`
		AwayTeamScore           string `json:"away_team_score"`
		Distance                string `json:"distance"`
		GroupID                 string `json:"group_id"`
		HasInjury               string `json:"has_injury"`
		HasNews                 string `json:"has_news"`
	*/
	AwayTeamID   string `json:"away_team_id"`
	AwayTeamName string `json:"away_team_name"`

	HistoryH2HMatchs []*MatchHistory `json:"historyH2HMatchs"`
	/*	HomeCoach struct {
		Association     string `json:"association"`
		AssociationLogo string `json:"association_logo"`
		Birthdate       string `json:"birthdate"`
		EnName          string `json:"en_name"`
		GameCount       string `json:"game_count"`
		ID              string `json:"id"`
		Logo            string `json:"logo"`
		Score           string `json:"score"`
		Style           string `json:"style"`
	} `json:"home_coach"`*/
	HomeFormation           string               `json:"home_formation"`
	HomeTeamHalfTimeScore   string               `json:"home_team_half_time_score"`
	HomeTeamID              string               `json:"home_team_id"`
	HomeTeamName            string               `json:"home_team_name"`
	HomeTeamNormalTimeScore string               `json:"home_team_normal_time_score"`
	HomeTeamScore           string               `json:"home_team_score"`
	ID                      string               `json:"id"`
	Informatins             []*MatchInformations `json:"informatins"`
	IsVisible               string               `json:"is_visible"`
	LmtMode                 string               `json:"lmt_mode"`
	MatchTime               string               `json:"match_time"`
	Neutral                 string               `json:"neutral"`
	Referee                 struct {
		CnName             string `json:"cn_name"`
		CountryCnName      string `json:"country_cn_name"`
		CountryEnName      string `json:"country_en_name"`
		CountryID          string `json:"country_id"`
		EnName             string `json:"en_name"`
		ID                 string `json:"id"`
		RedCardsPerGame    string `json:"redCardsPerGame"`
		YellowCardsPerGame string `json:"yellowCardsPerGame"`
	} `json:"referee"`
	RefereeID string `json:"referee_id"`
	Round     string `json:"round"`
	RoundType string `json:"round_type"`
	Scores    []struct {
		MatchID    string `json:"match_id"`
		Team1      string `json:"team1"`
		Team2      string `json:"team2"`
		Type       string `json:"type"`
		UpdateTime string `json:"update_time"`
	} `json:"scores"`
	Season   string `json:"season"`
	SeasonID string `json:"season_id"`
	SportID  string `json:"sport_id"`
	Stadium  struct {
		CnName        string `json:"cn_name"`
		CountryCnName string `json:"country_cn_name"`
		CountryEnName string `json:"country_en_name"`
		EnName        string `json:"en_name"`
		ID            string `json:"id"`
	} `json:"stadium"`

	StadiumID       string            `json:"stadium_id"`
	Status          string            `json:"status"`
	StatusCode      string            `json:"status_code"`
	StatusName      string            `json:"status_name"`
	TimePlayed      string            `json:"time_played"`
	TimeRemaining   string            `json:"time_remaining"`
	TimeRunning     string            `json:"time_running"`
	TimeUpdate      string            `json:"time_update"`
	TournamentID    string            `json:"tournament_id"`
	UpdateTimestamp string            `json:"update_timestamp"`
	Weather         string            `json:"weather"`
	WeatherDesc     string            `json:"weather_desc"`
	Winner          string            `json:"winner"`
	WinningOdds     *MatchWinningOdds `json:"winningOdds"`
}

type MatchHistory struct {
	AwayTeamHalfTimeScore   string `json:"away_team_half_time_score"`
	AwayTeamID              string `json:"away_team_id"`
	AwayTeamName            string `json:"away_team_name"`
	AwayTeamNormalTimeScore string `json:"away_team_normal_time_score"`
	AwayTeamScore           string `json:"away_team_score"`
	HomeTeamHalfTimeScore   string `json:"home_team_half_time_score"`
	HomeTeamID              string `json:"home_team_id"`
	HomeTeamName            string `json:"home_team_name"`
	HomeTeamNormalTimeScore string `json:"home_team_normal_time_score"`
	HomeTeamScore           string `json:"home_team_score"`
	ID                      string `json:"id"`
	MatchTime               string `json:"match_time"`
	Status                  string `json:"status"`
	StatusCode              string `json:"status_code"`
	StatusName              string `json:"status_name"`
	TournamentCnName        string `json:"tournament_cn_name"`
	TournamentID            string `json:"tournament_id"`
	Winner                  string `json:"winner"`
	AggregateWinner         string `json:"aggregate_winner,omitempty"`
}

type MatchInformations struct {
	Benefit  string `json:"benefit"`
	Content  string `json:"content"`
	TypeID   string `json:"type_id"`
	TypeName string `json:"type_name"`
}

type MatchWinningOdds struct {
	AwayActual       string `json:"away_actual"`
	AwayDecimalValue string `json:"away_decimal_value"`
	AwayDesc         string `json:"away_desc"`
	AwayExpected     string `json:"away_expected"`
	HomeActual       string `json:"home_actual"`
	HomeDecimalValue string `json:"home_decimal_value"`
	HomeDesc         string `json:"home_desc"`
	HomeExpected     string `json:"home_expected"`
}

type HistoryZip struct {
	ID         int    `json:"id"`
	Path       string `json:"path"`
	Updatetime string `json:"updatetime"`
}

type AllHistoryData struct {
	Groups              []*HistoryGroups              `json:"groups"`
	Matchs              []*HistoryMatchs              `json:"matchs"`
	PlayerSeasonStatics []*HistoryPlayerSeasonStatics `json:"player_season_statics"`
	Players             []*HistoryPlayers             `json:"players"`
	Season              *HistorySeason                `json:"season"`
	Tables              []*HistoryTables              `json:"tables"`
	TeamSeasonStatics   []*HistorySeasonStatics       `json:"team_season_statics"`
	Teams               []*HistoryTeam                `json:"teams"`
	Tournament          *HistoryTournament            `json:"tournament"`
}

type HistoryGroups struct {
	CnAlias string `json:"cn_alias"`
	CnName  string `json:"cn_name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	EnAlias string `json:"en_alias,omitempty"`
	EnName  string `json:"en_name,omitempty"`
}

type HistoryPlayers struct {
	Birthdate     string `json:"birthdate,omitempty"`
	CnAlias       string `json:"cn_alias"`
	CnName        string `json:"cn_name"`
	EnAlias       string `json:"en_alias"`
	EnName        string `json:"en_name"`
	Height        string `json:"height,omitempty"`
	ID            string `json:"id"`
	Nationality   string `json:"nationality,omitempty"`
	Position      string `json:"position,omitempty"`
	PreferredFoot string `json:"preferred_foot,omitempty"`
	ShirtNumber   string `json:"shirt_number,omitempty"`
	SportID       string `json:"sport_id"`
	//Statics       []interface{} `json:"statics"`
	TeamID string `json:"team_id"`
	Weight string `json:"weight,omitempty"`
	PicURL string `json:"pic_url,omitempty"`
}

type HistorySeason struct {
	HasPlayerStatistics string `json:"has_player_statistics"`
	HasRanking          string `json:"has_ranking"`
	ID                  string `json:"id"`
	Year                string `json:"year"`
}

type HistoryTables struct {
	Id      string                  `json:"id"`
	CnName  string                  `json:"cn_name"`
	EnName  string                  `json:"en_name"`
	GroupId string                  `json:"group_id"`
	Statics []*HistoryTablesStatics `json:"statics"`
}

type HistoryTablesStatics struct {
	TeamId          string `json:"team_id"`
	TypeId          string `json:"type_id"`
	TypeName        string `json:"type_name"`
	TypeEnName      string `json:"type_en_name"`
	TypeCnName      string `json:"type_cn_name"`
	Value           string `json:"value"`
	PromotionId     string `json:"promotion_id"`
	PromotionCnName string `json:"promotion_cn_name"`
	PromotionEnName string `json:"promotion_en_name"`
}

type HistoryTeam struct {
	CnAlias       string `json:"cn_alias"`
	CnName        string `json:"cn_name"`
	CoachCnName   string `json:"coach_cn_name,omitempty"`
	CoachEnName   string `json:"coach_en_name,omitempty"`
	CoachID       string `json:"coach_id,omitempty"`
	EnAlias       string `json:"en_alias"`
	EnName        string `json:"en_name"`
	ID            string `json:"id"`
	IsCountry     string `json:"is_country,omitempty"`
	LogoURL       string `json:"logo_url"`
	StadiumCnName string `json:"stadium_cn_name,omitempty"`
	StadiumEnName string `json:"stadium_en_name,omitempty"`
	StadiumID     string `json:"stadium_id,omitempty"`
	EstablishDate string `json:"establish_date,omitempty"`
}

type HistoryTournament struct {
	CategoryCnName    string `json:"category_cn_name"`
	CategoryContinent string `json:"category_continent"`
	CategoryID        string `json:"category_id"`
	CnAlias           string `json:"cn_alias"`
	CnName            string `json:"cn_name"`
	CurrentSeasonID   string `json:"current_season_id"`
	EnAlias           string `json:"en_alias"`
	EnName            string `json:"en_name"`
	ID                string `json:"id"`
	Level             string `json:"level"`
	LogoURL           string `json:"logo_url"`
	PeriodType        string `json:"period_type"`
	Reverse           string `json:"reverse"`
	SportID           string `json:"sport_id"`
	Type              string `json:"type"`
}

type TeamData struct {
	CnAlias           string               `json:"cn_alias"`
	CnName            string               `json:"cn_name"`
	CoachCnName       string               `json:"coach_cn_name"`
	CoachEnName       string               `json:"coach_en_name"`
	CoachID           string               `json:"coach_id"`
	EnAlias           string               `json:"en_alias"`
	EnName            string               `json:"en_name"`
	EstablishDate     string               `json:"establish_date"`
	ID                string               `json:"id"`
	IsCountry         string               `json:"is_country"`
	LogoURL           string               `json:"logo_url"`
	Players           []*TeamPlayer        `json:"players"`
	PlayersMissing    []*PlayersMissing    `json:"players_missing"`
	SportID           string               `json:"sport_id"`
	StadiumCnName     string               `json:"stadium_cn_name"`
	StadiumEnName     string               `json:"stadium_en_name"`
	StadiumID         string               `json:"stadium_id"`
	Tables            []*TeamTables        `json:"tables"`
	UpdateTime        string               `json:"update_time"`
	TeamSeasonStatics []*TeamSeasonStatics `json:"team_season_statics"`
}

type TeamSeasonStatics struct {
	SeasonId     string `json:"season_id"`
	TournamentId string `json:"tournament_id"`
	TeamId       string `json:"team_id"`
	TypeId       string `json:"type_id"`
	TypeValue    string `json:"type_value"`
	Range        string `json:"range"`
}

type TeamPlayer struct {
	Birthdate     string `json:"birthdate"`
	CnAlias       string `json:"cn_alias"`
	CnName        string `json:"cn_name"`
	EnAlias       string `json:"en_alias"`
	EnName        string `json:"en_name"`
	Height        string `json:"height,omitempty"`
	ID            string `json:"id"`
	Nationality   string `json:"nationality"`
	PicURL        string `json:"pic_url,omitempty"`
	Position      string `json:"position"`
	PreferredFoot string `json:"preferred_foot,omitempty"`
	ShirtNumber   string `json:"shirt_number"`
	UpdateTime    string `json:"update_time"`
	Weight        string `json:"weight,omitempty"`
}

type PlayersMissing struct {
	ID             string `json:"id"`
	PlayerID       string `json:"player_id"`
	PlayerName     string `json:"player_name"`
	Reason         string `json:"reason"`
	RegisteredTime string `json:"registeredTime"`
	TournamentID   string `json:"tournament_id"`
}

type TeamTables struct {
	CnName         string         `json:"cn_name"`
	EnName         string         `json:"en_name"`
	GroupCnName    string         `json:"group_cn_name"`
	GroupID        string         `json:"group_id"`
	ID             string         `json:"id"`
	SeasonID       string         `json:"season_id"`
	SeasonYear     string         `json:"season_year"`
	Statics        []*TeamStatics `json:"statics"`
	TournamentID   string         `json:"tournament_id"`
	TournamentName string         `json:"tournament_name"`
}

type TeamStatics struct {
	TypeCnName string `json:"type_cn_name"`
	TypeEnName string `json:"type_en_name"`
	TypeID     string `json:"type_id"`
	Value      string `json:"value"`
}

type CenterBaseInfoResp struct {
	AwayFormation string `json:"away_formation"`
	HomeFormation string `json:"home_formation"`
	CenterBaseInfo
	StaticMatchesInfoStoreDataRespPlayers
}

//公用元素
type CenterBaseInfo struct {
	AwayLogoBase string `json:"awayLogo"`
	HomeLogoBase string `json:"homeLogo"`
	AwayNameBase string `json:"awayName"`
	HomeNameBase string `json:"homeName"`
	AwayIdBase   string `json:"awayId"`
	HomeIdBase   string `json:"homeId"`
}

type StaticMatchesInfoStoreDataRespPlayers struct {
	StartingLineUp   map[string][]*HistoryMatchesLineUp `json:"starting_lineUp"`
	InjureLineUp     map[string][]*HistoryMatchesLineUp `json:"injure_lineUp"`
	SubstituteLineUp map[string][]*HistoryMatchesLineUp `json:"substitute_lineUp"`
}

type XjMatchesData struct {
	EventId          string `json:"EventId"`
	CompetitionName  string `json:"CompetitionName"`
	HomeTeam         string `json:"HomeTeam"`
	AwayTeam         string `json:"AwayTeam"`
	EventDate        string `json:"EventDate"`
	Category         string `json:"category"`
	HomeTeamLogoUrl  string `json:"homeTeamLogoUrl"`
	GuestTeamLogoUrl string `json:"guestTeamLogoUrl"`
}

type Anchor struct {
	Grade           string `json:"grade"`
	ID              string `json:"id"`
	LogoRectangle   string `json:"logo_rectangle"`
	LogoSquare      string `json:"logo_square"`
	Nickname        string `json:"nickname"`
	PersonalProfile string `json:"personal_profile"`
	RecomTimes      int    `json:"recomTimes"`
	RecomTime       int64  `json:"recomTime"`
	Sex             string `json:"sex"`
	Labels          string `json:"labels"`
	IsDelete        string `json:"is_delete"`
	Active          string `json:"active"`
}
