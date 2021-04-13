package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	NewFile()
}

var fd string = "test.xlsx"

func ReadFile() {
	file, err := xlsx.OpenFile(fd) // 打开一个文件
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
	c := sheet.Cell(0, 0)
	c.Value = "实例"
	err = file.Save(fd)
	fmt.Println(err)
}
