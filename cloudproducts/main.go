package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

//
func main() {
	mobileJson := GetProductJson()
	//items := make([]BindDataItem, 0)
	data := mobileJson.Body[0].BindData
	dir, _ := os.Getwd()
	file := excelize.NewFile()
	id := 1
	for index, value := range data {
		if index == 0 || index == 13 || index == 14 {
			continue
		}
		cell := fmt.Sprintf("A%d", id)
		sheet := "Sheet1"
		file.SetCellValue(sheet, cell, value.Name)
		for _, child := range value.Children {
			childCell := fmt.Sprintf("B%d", id)
			file.SetCellValue(sheet, childCell, child.Name)
			for _, metadata := range child.BindMetadatas {
				bindCell := fmt.Sprintf("C%d", id)
				file.SetCellValue(sheet, bindCell, metadata.Name)
				Desc := fmt.Sprintf("D%d", id)
				file.SetCellValue(sheet, Desc, metadata.SDesc)
				id++
			}
		}
	}

	//xlFile, err := xlsx.OpenFile(dir + "/cloudproducts/mobile.xlsx")
	//if err != nil {
	//	fmt.Println("打开文件错误:", err)
	//	return
	//}
	// 遍历所有工作表
	file.SaveAs(dir + "/cloudproducts/mobile.xlsx")
}
