package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "12345678", // no password set
		DB:       0,          // use default DB
	})

	// Test the connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis")

	// Add scores to the sorted set
	err = rdb.ZAdd(context.Background(), "match_scores", &redis.Z{
		Score:  0,
		Member: "红队 1 - 0 蓝队",
	}, &redis.Z{
		Score:  1,
		Member: "红队 2 - 1 蓝队",
	}, &redis.Z{
		Score:  2,
		Member: "红队 3 - 2 蓝队",
	}).Err()
	if err != nil {
		log.Fatalf("Failed to add scores to Redis: %v", err)
	}
	fmt.Println("Scores added to Redis")

	// Update a score in the sorted set
	err = rdb.ZAdd(context.Background(), "match_scores", &redis.Z{
		Score:  2,
		Member: "红队 3 - 3 蓝队",
	}).Err()
	if err != nil {
		log.Fatalf("Failed to update score in Redis: %v", err)
	}
	fmt.Println("Score updated in Redis")

	// Get the latest score
	latestScore, err := rdb.ZRevRange(context.Background(), "match_scores", 0, 0).Result()
	if err != nil {
		log.Fatalf("Failed to get latest score from Redis: %v", err)
	}
	fmt.Printf("最新比分：%s\n", latestScore[0])

	// Get the previous score
	previousScore, err := rdb.ZRevRange(context.Background(), "match_scores", 1, 1).Result()
	if err != nil {
		log.Fatalf("Failed to get previous score from Redis: %v", err)
	}
	fmt.Printf("上一次比分：%s\n", previousScore[0])

	// Compare the current and previous scores
	if latestScore[0] == previousScore[0] {
		fmt.Println("比分未变化")
	} else {
		fmt.Println("比分发生变化")
	}
}
