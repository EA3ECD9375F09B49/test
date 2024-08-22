package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 读取 Excel 文件
	xlsx, err := excelize.OpenFile("./y_gamesite2.xlsx")
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return
	}

	// 读取 constant 对应的值
	constantMap := make(map[string]string)
	rows := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		constantMap[row[0]] = row[1]
	}

	// 遍历指定文件夹中的 Go 文件
	//root := "./../../../coderep.mdhue.com/site/gameSite/controller"
	//root := "./../../../y8/uploadservice/controller"
	//root := "./../../../y8/gameSite/controller"
	//root := "./../../../y8/apiSite/controller"
	root := "./../../../y8/gameSite/service/gameSer"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// 读取文件内容
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return nil
			}

			//// 检查是否已导入 util 包，如果没有则导入
			//if !strings.Contains(string(fileData), "util") {
			//	fileData = append([]byte("import \"your_package_path/utils\"\n\n"), fileData...)
			//}

			// 替换文件内容
			for key, value := range constantMap {
				fileData = []byte(strings.ReplaceAll(string(fileData), `"`+key+`"`, `base.I18nSprintf(utils.`+value+`)`))
			}

			// 写入替换后的内容
			err = ioutil.WriteFile(path, fileData, 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return nil
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}
}
