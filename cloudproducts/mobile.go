package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetProductJson() MobileJson {
	url := "https://ecloud.10086.cn/api/query/menu/v4/menu/tree?situationCode=product"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var mobile MobileJson
	json.Unmarshal(body, &mobile)
	return mobile
}

type BindDataItem struct {
	Name          string         `json:"name"`
	Desc          string         `json:"desc"`
	Level         int64          `json:"level"`
	BindMetadatas []BindMetadata `json:"bindMetadatas"`
}

type BindMetadata struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Level int64  `json:"level"`
}

type MobileJson struct {
	RequestId interface{} `json:"requestId"`
	State     string      `json:"state"`
	Body      []struct {
		SituationCode string `json:"situationCode"`
		BindData      []struct {
			Id            string        `json:"id"`
			Name          string        `json:"name"`
			Sort          int           `json:"sort"`
			Level         int           `json:"level"`
			ParentId      string        `json:"parentId"`
			Type          string        `json:"type"`
			BindLabel     string        `json:"bindLabel"`
			BindMetadatas []interface{} `json:"bindMetadatas"`
			Children      []struct {
				Id            string `json:"id"`
				Name          string `json:"name"`
				Sort          int    `json:"sort"`
				Level         int    `json:"level"`
				ParentId      string `json:"parentId"`
				Type          string `json:"type,omitempty"`
				BindLabel     string `json:"bindLabel"`
				BindMetadatas []struct {
					Id                   string `json:"id"`
					Name                 string `json:"name"`
					ProductType          string `json:"productType,omitempty"`
					Type                 string `json:"type,omitempty"`
					Sort                 int    `json:"sort"`
					SDesc                string `json:"sDesc,omitempty"`
					Status               string `json:"status"`
					IsPay                bool   `json:"isPay,omitempty"`
					IsTest               bool   `json:"isTest,omitempty"`
					IsLeoTemplate        bool   `json:"isLeoTemplate"`
					BackgroundUrl        string `json:"backgroundUrl"`
					IntroduceUrl         string `json:"introduceUrl,omitempty"`
					OrderUrl             string `json:"orderUrl,omitempty"`
					ConsoleUrl           string `json:"consoleUrl,omitempty"`
					OrderUrlH5           string `json:"orderUrlH5,omitempty"`
					ConsoleUrlH5         string `json:"consoleUrlH5,omitempty"`
					PriceDetailUrl       string `json:"priceDetailUrl,omitempty"`
					PriceCalculatorUrl   string `json:"priceCalculatorUrl,omitempty"`
					PriceCalculatorUrlH5 string `json:"priceCalculatorUrlH5,omitempty"`
					HelpCenterUrl        string `json:"helpCenterUrl,omitempty"`
					Brief                string `json:"brief,omitempty"`
					SpecMarker           string `json:"specMarker,omitempty"`
					SubName              string `json:"subName,omitempty"`
					ServiceProtocolUrl   string `json:"serviceProtocolUrl,omitempty"`
				} `json:"bindMetadatas"`
				Children []struct {
					Id            string `json:"id"`
					Name          string `json:"name"`
					Sort          int    `json:"sort"`
					Level         int    `json:"level"`
					ParentId      string `json:"parentId"`
					BindLabel     string `json:"bindLabel"`
					BindMetadatas []struct {
						Id                   string `json:"id"`
						Name                 string `json:"name"`
						ProductType          string `json:"productType"`
						Type                 string `json:"type"`
						Sort                 int    `json:"sort"`
						SDesc                string `json:"sDesc"`
						Status               string `json:"status"`
						IsPay                bool   `json:"isPay"`
						IsTest               bool   `json:"isTest"`
						IsLeoTemplate        bool   `json:"isLeoTemplate"`
						BackgroundUrl        string `json:"backgroundUrl"`
						IntroduceUrl         string `json:"introduceUrl"`
						OrderUrl             string `json:"orderUrl"`
						ConsoleUrl           string `json:"consoleUrl"`
						OrderUrlH5           string `json:"orderUrlH5,omitempty"`
						ConsoleUrlH5         string `json:"consoleUrlH5,omitempty"`
						PriceDetailUrl       string `json:"priceDetailUrl"`
						PriceCalculatorUrl   string `json:"priceCalculatorUrl,omitempty"`
						PriceCalculatorUrlH5 string `json:"priceCalculatorUrlH5,omitempty"`
						HelpCenterUrl        string `json:"helpCenterUrl"`
						Brief                string `json:"brief,omitempty"`
					} `json:"bindMetadatas"`
					Children []interface{} `json:"children"`
				} `json:"children"`
			} `json:"children"`
		} `json:"bindData"`
	} `json:"body"`
}
