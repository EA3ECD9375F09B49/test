package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"path"
)

func ReadCsvOrXlsxFileData(dst string, check bool) ([][]string, error) {
	var (
		rows [][]string
	)
	fs, err := os.Open(dst)
	if err != nil {
		return rows, err
	}
	defer func(fs *os.File) {
		err = fs.Close()
		if err != nil {
			_ = fmt.Sprintf("readCsvOrXlsxFileData %v", err)
		}
	}(fs)
	fileSuffix := path.Ext(fs.Name())
	switch fileSuffix {
	case ".csv":
		reader := csv.NewReader(fs)
		reader.FieldsPerRecord = -1
		if !check { // 结算 时候 老逻辑使用  rows[][]string,string 模式
			reader.Comma = ','
			reader.LazyQuotes = true
		}
		rows, err = reader.ReadAll()
		if err != nil {
			return rows, err
		}
		return rows, nil
	case ".xlsx":
		xlFile, err := xlsx.OpenFile(dst)
		if err != nil {
			return rows, err
		}
		if len(xlFile.Sheets) == 0 {
			return rows, err
		}
		sheet := xlFile.Sheets[0]
		for _, sRow := range sheet.Rows {
			var (
				innerRowCheck           []string
				innerRowCalculate       []string
				innerRowCalculateString string
			)
			if len(sRow.Cells) == 0 {
				continue
			}
			for i, cell := range sRow.Cells {
				if i < 1 && len(innerRowCalculateString) == 0 { // 检查的 时候 老逻辑使用  rows[][]string 模式
					innerRowCalculateString = cell.Value
				} else { // 结算 时候 老逻辑使用  rows[][]string,string 模式
					innerRowCalculateString += "," + cell.Value
				}
				innerRowCheck = append(innerRowCheck, cell.Value)
			}
			if check { // 检查的 时候 老逻辑使用  rows[][]string 模式
				rows = append(rows, innerRowCheck)
			} else { // 结算 时候 老逻辑使用  rows[][]string,string 模式
				innerRowCalculate = append(innerRowCalculate, innerRowCalculateString)
				rows = append(rows, innerRowCalculate)
			}
		}
		return rows, nil
	default:
		return rows, errors.New("文件格式不支持")
	}
}
