```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	. "gocmd/colorPrint"
	"io/ioutil"
	"os"
	"path/filepath"
)
var FontColor Color = Color{Blue: 1, Green: 2, Cyan: 3, Red: 4, Purple: 5, Yellow: 6, LightGray: 7, Gray: 8, LightBlue: 9, LightGreen: 10, LightCyan: 11, LightRed: 12, LightPurple: 13, LightYellow: 14, White: 15}
func main() {
	source := "i18n.xlsx"
	// fmt.Println(os.Args)
	if len(os.Args) > 1 {
		source = os.Args[1]
	}
	currentDir,_ := os.Getwd()
	docPath := filepath.Join(currentDir, source)
	zhPath := filepath.Join(currentDir, "zh_CN.json")
	enPath := filepath.Join(currentDir, "en_US.json")
	ColorPrint("", FontColor.Blue)
	if fileExist(zhPath) {
		err := os.Remove(zhPath)
		msg := "清除成功：" + zhPath
		if err != nil {
			msg = "清除失败：" + zhPath
		}
		fmt.Println(msg)
	}
	if fileExist(enPath) {
		err := os.Remove(enPath)
		if err != nil {
			fmt.Println("清除失败：", enPath)
		} else {
			fmt.Println("清除成功：", enPath)	
		}
	}
	f, err := excelize.OpenFile(docPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	zhMap := make(map[string] string)
	enMap := make(map[string] string)
	rows, err := f.GetRows("Sheet1")
	rowLoop:
	for rowIndex, row := range rows {
		for colIndex, cell := range row {
			// fmt.Print(cell, "\t\t")
			if cell == "" {
				if colIndex == 0 {
					break rowLoop
				}
				break 
			} 
		}
		if rowIndex > 0 {
			zhMap[row[0]] = row[1]	
			enMap[row[0]] = row[2]	
		}
		// fmt.Println()
	}
	zhData, err := json.MarshalIndent(zhMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(zhData))
	err = ioutil.WriteFile(zhPath, zhData, 0644)
	if err == nil {
		fmt.Println("ZH-Json输出成功：", zhPath)
	}
	enData, err := json.MarshalIndent(enMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(enPath, enData, os.ModeAppend)
	if err == nil {
		fmt.Println("EN-Json输出成功：", enPath)
	}
	ColorPrint("", FontColor.White)
}

func fileExist(path string) bool {
  _, err := os.Lstat(path)
  return !os.IsNotExist(err)
}
```
