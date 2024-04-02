package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	// 打开 Excel 文件
	excelFileName := "./checkFile/toCheckUrl.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("打开 Excel 文件时出错:", err)
		return
	}

	// 创建新的 Excel 文件
	newFile := xlsx.NewFile()
	sheetR, err := newFile.AddSheet("结果")
	if err != nil {
		fmt.Println("创建新的 Excel 文件时出错:", err)
		return
	}

	// 创建等待组
	var wg sync.WaitGroup

	// 循环处理每行数据
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			urlCell := row.Cells[1]        // 假设 URL 字段在第一列
			departmentCell := row.Cells[2] // 假设 Department 字段在第二列
			url := urlCell.String()

			// 增加等待组计数
			wg.Add(1)
			// 将处理后的数据写入新的 Excel 文件
			newRow := sheetR.AddRow()
			// 并行处理 URL 的可访问性
			go func(newRow *xlsx.Row) {
				defer wg.Done()
				if url != "" {
					resp, err := http.Get(url)
					if err != nil {
						fmt.Println("无法访问 URL:", url)
						departmentCell.SetString("不可访问")
					} else {
						fmt.Println("可以访问 URL:", url)
						departmentCell.SetString("可以访问")
					}
					fmt.Println("访问 URL 结果:", resp)
					cell1 := newRow.AddCell()
					cell1.Value = url
					cell2 := newRow.AddCell()
					cell2.Value = departmentCell.String()
				}
			}(newRow)

			// 并行处理 IP 地址和端口的可访问性
			//go func(ip, port string, departmentCell *xlsx.Cell) {
			//	defer wg.Done()
			//	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, port), 2*time.Second)
			//	if err != nil {
			//		fmt.Println("无法访问 IP 地址和端口:", ip, port)
			//		departmentCell.SetString("不可访问")
			//	} else {
			//		defer conn.Close()
			//		fmt.Println("可以访问 IP 地址和端口:", ip, port)
			//		departmentCell.SetString("可以访问")
			//	}
			//}(ip, port, departmentCell)

		}
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 保存新的 Excel 文件
	_, fileName := filepath.Split(excelFileName)
	newFileName := "./Result/" + strings.TrimSuffix(fileName, ".xlsx") + "_result.xlsx"
	newFile.Save(newFileName)
	fmt.Println("处理完成，结果已保存到", newFileName)
}
