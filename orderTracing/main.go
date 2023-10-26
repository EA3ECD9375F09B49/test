package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"orderTracing/response/IMTY"
	"orderTracing/utils"
	"os"
	"strings"
	"sync"
)

type AdjustCalculationObj struct {
	Error    error
	Wg       *sync.WaitGroup
	faiCount int
	failList []string
	Mutex    sync.Mutex //并发锁
}

func main() {

	var (
		resp            *IMTY.ImSporRecordResp
		filler          IMTY.ToFiller
		successCount    int
		failCount       int
		ourSuccessCount int
		ourFailCount    int
	)
	//
	ThirdData := make(map[string]interface{})
	OurData := make(map[string]interface{})

	jsonFiles := "files/IMTY/response23-12-13.json"
	//upload/批量账户模板20221106 copy.csv
	dst := "files/IMTY/投注记录20221203033838421-1.csv"
	logFileName := "logs/我方缺少数据-23-12-13.txt"

	logFileName2 := "logs/我方多余数据-23-12-13.txt"

	logFileName3 := "logs/需要补单的订单-23-12-13.txt"

	thirdPartyJsonString, _ := ioutil.ReadFile(jsonFiles)
	_ = json.Unmarshal(thirdPartyJsonString, &resp)

	for _, thirdData := range resp.WagerArr {
		ThirdData[thirdData.WagerID] = thirdData
	}

	rows, err := utils.ReadCsvOrXlsxFileData(dst, false)
	if err != nil {
		//c.Errorf("BatchAdjustCreate 批量用户调整-检查  os.Open err:%s", err.Error())
		//base.WebResp(c, http.StatusOK, utils.ErrInternal, nil, "服务器异常")
		return
	}

	f, err := os.Create(logFileName)
	f, err = os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f3, err := os.Create(logFileName3)
	f3, err = os.OpenFile(logFileName3, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer f3.Close()

	for i, item := range rows {
		if i == 0 || i == len(rows)-1 {
			continue
		}
		item[16] = strings.Trim(item[16], "\t")
		OurData[item[16]] = item
	}

	calObj := new(AdjustCalculationObj)
	calObj.Wg = new(sync.WaitGroup)
	for j, item := range ThirdData {

		if _, ok := OurData[j]; ok {
			successCount++
		} else {
			failCount++
			calObj.Wg.Add(2)
			wager := item.(*IMTY.Wager)
			filler.EventName = wager.WagerItemArr[0].EventName
			filler.WagerID = wager.WagerID
			filler.WagerCreationDateTime = wager.WagerCreationDateTime
			filler.LastUpdatedDate = wager.LastUpdatedDate
			filler.IsSettled = wager.IsSettled
			filler.IsSettled = wager.IsSettled
			filler.IsConfirmed = wager.IsConfirmed
			filler.IsConfirmed = wager.IsConfirmed
			go calObj.AdjustParallelCreator(item, f)
			go calObj.AdjustParallelCreator(filler, f3)
		}

	}
	if failCount > 0 {
		calObj.Wg.Wait()
	}

	f2, err := os.Create(logFileName2)
	f2, err = os.OpenFile(logFileName2, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer f2.Close()

	for k, item := range OurData {

		if _, ok := ThirdData[k]; ok {
			ourSuccessCount++
		} else {
			ourFailCount++
			calObj.Wg.Add(1)
			go calObj.AdjustParallelCreator(item, f2)
		}
	}
	if failCount > 0 {
		calObj.Wg.Wait()
	}

	fmt.Printf("done  successCount %d , failCount %d,totalthird %d  totalOurData %d", successCount, failCount, len(ThirdData), len(OurData))
}

func (calObj *AdjustCalculationObj) AdjustParallelCreator(rows interface{}, f *os.File) {
	defer calObj.Wg.Done()
	calObj.Mutex.Lock()

	data, _ := json.Marshal(rows)
	_, _ = f.WriteString(string(data) + "\n")
	calObj.Mutex.Unlock()
}
