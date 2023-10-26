package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"

	"time"

	"golang.org/x/net/context"
)

var ctx = context.Background()

func getTimeStamp(timeString string) int64 {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("解析日期时间字符串失败:", err)
		return 0
	}
	return parsedTime.Unix()
}

func main() {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 你的Redis服务器地址
		DB:       1,
		Password: "12345678",
	})

	inputTimeString1 := "2023-10-25 16:40:00"
	inputTimeString2 := "2023-10-25 16:40:34"

	// 转换为时间戳（Unix时间戳）
	timestamp1 := getTimeStamp(inputTimeString1)
	timestamp2 := getTimeStamp(inputTimeString2)

	client.ZAdd(ctx, "myzset", redis.Z{Score: float64(timestamp1), Member: 1})
	client.ZAdd(ctx, "myzset", redis.Z{Score: float64(timestamp2), Member: 2})

	// 设置时间范围
	eg1 := "2023-10-25 16:00:00"
	eg2 := "2023-10-25 17:00:00"
	startTime := getTimeStamp(eg1)
	endTime := getTimeStamp(eg2)

	//startTime := time.Date(2023, 10, 25, 16, 0, 0, 0, time.UTC).Unix()
	//endTime := time.Date(2023, 10, 25, 17, 0, 0, 0, time.UTC).Unix()

	// 获取时间范围内的数据
	result, err := client.ZRangeByScore(ctx, "myzset", &redis.ZRangeBy{
		Min: fmt.Sprintf("%f", float64(startTime)),
		Max: fmt.Sprintf(
			"%f", float64(endTime)),
	}).Result()
	if err != nil {
		panic(err)
	}

	// 打印结果
	for _, z := range result {
		aa := fmt.Sprintf("%v", z)
		fmt.Println(aa)
		//fmt.Printf("Member: %s, Score: %f\n", z.Member.(string), z.Score)
	}
}
