package main

import (
	"bufio"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"

	"time"

	"golang.org/x/net/context"
)

var ctx = context.Background()

const layout = "2006-01-02 15:04:05"

func GetBjTimeLoc() *time.Location {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return bjLoc
}

func ToBjTime(timeStr string) (time.Time, error) {
	beiJinLocation := GetBjTimeLoc()
	bgTime, err := time.ParseInLocation(layout, timeStr, beiJinLocation)
	if err != nil {
		return time.Time{}, err
	}
	return bgTime, nil
}

func getTime(timeString string) time.Time {
	parsedTime, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("解析日期时间字符串失败:", err)
		return time.Now()
	}
	return parsedTime
}

func initRedis() *redis.Client {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 你的Redis服务器地址
		DB:       1,
		Password: "12345678",
	})
	return client
}

type TimePair struct {
	Start, End time.Time
}

func SplitTime(start, end time.Time, interval time.Duration) []TimePair {
	if interval == 0 {
		return []TimePair{{start, end}}
	}

	split := []TimePair{}
	begin := start
	for end.Sub(begin) > 0 {
		next := begin.Add(interval)
		if next.Sub(end) > 0 {
			next = end
		}
		split = append(split, TimePair{begin, next})
		begin = next
	}

	return split
}

func main() {
	interval := 1 * time.Hour * 24 * 30
	client := initRedis()
	reader := bufio.NewReader(os.Stdin)
	var (
		eg1 string
		eg2 string
		err error
	)
	for {
		fmt.Printf("please enter range 1 2023-01-01 00:00:00 format \n")
		eg1, err = reader.ReadString('\n')
		eg1 = strings.Trim(eg1, "\f\t\r\n ")
		// convert CRLF to LF
		eg1 = strings.Replace(eg1, "\n", "", -1)
		if err != nil {
			return
		}
		fmt.Printf("The range 1 is %s \n",
			eg1)
		fmt.Printf("please enter range 2 2023-01-01 00:00:00 format \n")
		eg2, err = reader.ReadString('\n')
		eg2 = strings.Trim(eg2, "\f\t\r\n ")
		// convert CRLF to LF
		eg2 = strings.Replace(eg2, "\n", "", -1)
		if err != nil {
			return
		}
		fmt.Printf("The range 2 is %s \n",
			eg2)

		zsetName := "mysetmonth"
		zcardResult, err := client.ZCard(ctx, zsetName).Result()
		if err != nil {
			fmt.Printf("获取有序集合成员数量失败: %v", err)
		}
		fmt.Printf("有序集合 %s 中的成员数量: %d\n", zsetName, zcardResult)
		sTime, err := ToBjTime(eg1)
		eTime, err := ToBjTime(eg2)
		splitTime := SplitTime(sTime, eTime, interval)
		for _, timepair := range splitTime {
			startTime := timepair.Start.Unix()
			// 获取当前时间戳
			currentTime := time.Now()
			// 计算半年前的时间戳
			sixMonthsAgo := currentTime.AddDate(0, -6, 0)
			sixMonthsAgoTimestamp := sixMonthsAgo.Unix() //时间戳在半年内，允许通过
			if startTime >= sixMonthsAgoTimestamp {
				client.ZAdd(ctx, "mysetmonth", redis.Z{Score: float64(startTime), Member: timepair.Start.String()})
				fmt.Printf("时间戳在半年内，允许通过  %v \n\n", timepair.Start.String())
			} else {
				fmt.Printf("时间戳在半年前，不允许通过  %v \n\n", timepair.Start.String())
			}
		}

	}

}
