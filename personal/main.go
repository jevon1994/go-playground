package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	xlFile, err := xlsx.OpenFile(dir + "/personal/location.xlsx")
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}
	land := make([]string, 1)
	location := make([]string, 1)
	// 遍历所有工作表
	for _, sheet := range xlFile.Sheets {
		fmt.Println("工作表:", sheet.Name)

		// 遍历所有行
		for _, row := range sheet.Rows {
			// 遍历所有单元格
			for i, cell := range row.Cells {
				if i == 3 {
					if strings.HasPrefix(cell.Value, "余政储出") {
						land = append(land, cell.Value)
					} else {
						location = append(location, cell.Value)
						street := getStreetByLocation(cell.Value)
						addCell := row.AddCell()
						if street == "" {
							street = "未查询到"
						}
						addCell.SetValue(street)
					}
				}
			}
			fmt.Println("")
		}
	}
	fmt.Printf("========= land: %s\n ", land)

	fmt.Printf("========= location: %s\n ", location)
}

type GeoResult struct {
	Geocodes []struct {
		FormattedAddress string `json:"formatted_address"`
		AddressComponent struct {
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			Street   string `json:"street"`
			Number   string `json:"number"`
		} `json:"address_component"`
	} `json:"geocodes"`
}

func getStreetByLocation(location string) string {
	// 高德地图 API 的请求 URL，根据实际情况填写 key 和 address 参数
	apiUrl := "https://restapi.amap.com/v3/geocode/geo?key=37c0cb59380dd78e5491d5add98f647f&address=" + location

	// 发送 HTTP GET 请求获取响应
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}

	// 解析响应内容为结构体
	var result GeoResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}

	// 从结构体中提取出街道信息
	if len(result.Geocodes) > 0 {
		street := result.Geocodes[0].AddressComponent.Street
		fmt.Printf("Street: %s\n", street)
	} else {
		fmt.Println("No result found")
	}
	return ""
}
