package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	NewFile()
}

func ReadFile() {
	filePath := "frame/excel/test.xlsx"
	file, err := xlsx.OpenFile(filePath) // 打开一个文件
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sheet := range file.Sheets {
		for _, rows := range sheet.Rows {
			for _, cell := range rows.Cells {
				fmt.Println(cell.Value)
			}
		}
	}
}

func NewFile() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("创建工作表失败", err)
		return
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "我是一个表格"
	err = file.Save("frame/excel/test.xlsx")
	fmt.Println(err)
}
