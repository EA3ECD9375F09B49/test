package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func convertToPinyin(name string) string {
	var result strings.Builder
	for _, char := range name {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			result.WriteRune(unicode.ToLower(char))
		} else if char > 127 {
			// Convert Chinese characters to Pinyin
			pinyinArr := pinyin.Pinyin(string(char), pinyin.NewArgs())
			if len(pinyinArr) > 0 {
				result.WriteString(pinyinArr[0][0]) // Use the first pronunciation
			}
		} else { // It's a special character
			result.WriteRune('_') // Replace with underscore
		}
	}

	// Replace spaces and special characters with underscores
	re := regexp.MustCompile(`['_\s-]+`)
	return re.ReplaceAllString(result.String(), "_")
}

func renameImageFiles(sourceDir, targetDir string) ([][]string, error) {
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}

	var logData [][]string
	logData = append(logData, []string{"目录名", "原文件名", "最终文件名"}) // Add header row

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has an image extension
		for _, ext := range imageExtensions {
			if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ext) {
				// Separate the file name and extension
				nameWithoutExt := strings.TrimSpace(strings.TrimSuffix(info.Name(), ext))

				// Convert name and construct new name
				newName := convertToPinyin(nameWithoutExt) + ext

				// Create the target path maintaining directory structure
				relativePath, err := filepath.Rel(sourceDir, path)
				if err != nil {
					return err
				}
				newPath := filepath.Join(targetDir, relativePath)

				// Create the directory if it doesn't exist
				if err := os.MkdirAll(filepath.Dir(newPath), os.ModePerm); err != nil {
					return err
				}

				// Log the file names
				logData = append(logData, []string{filepath.Dir(path), info.Name(), newName})
				// Copy the file to the new path with the new name
				newFilePath := filepath.Join(filepath.Dir(newPath), newName)
				fmt.Printf("Copying %s to %s\n", path, newFilePath)
				input, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				return ioutil.WriteFile(newFilePath, input, info.Mode())
			}
		}
		return nil
	})
	return logData, err
}

func writeLogToExcel(logData [][]string, filename string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet := "Sheet1"
	for rowIdx, rowData := range logData {
		for colIdx, cellValue := range rowData {
			cellName, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)
			err := f.SetCellValue(sheet, cellName, cellValue)
			if err != nil {
				return err
			}
		}
	}

	return f.SaveAs(filename)
}

func main() {
	// Specify the source directory and target directory
	sourceDirectory := "./inDir"  // Replace with your source directory
	targetDirectory := "./outDir" // Replace with your target directory

	logData, err := renameImageFiles(sourceDirectory, targetDirectory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Get the current date in the desired format
	currentDate := time.Now().Format("2006-01-02") // "YYYY-MM-DD" format
	// Construct the filename with the date
	fileName := fmt.Sprintf("./outDir/renameInfo_%s.xlsx", currentDate)

	if err := writeLogToExcel(logData, fileName); err != nil {
		fmt.Println("Error writing log to Excel:", err)
	}
}
