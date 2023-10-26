package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GamegamesImg struct {
	Url   string `db:"url"`
	H5Url string `db:"h5_url"`
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3307)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	var gamesImg GamegamesImg
	db.Raw("select url,h5_url from game_games limit ?", 1).Scan(&gamesImg)
	if db.Error != nil {
		fmt.Printf("get failed err:%v", err)
		return
	}
	fmt.Println(gamesImg)
}
