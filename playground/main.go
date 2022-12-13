package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/FerdinaKusumah/excel2json"
)

type Entry struct {
	ECode        string `json:"eCode"`
	EDate        string `json:"eDate"`
	ELocation    string `json:"eLocation"`
	EPieces      string `json:"ePieces"`
	EVolume      string `json:"eVolume"`
	EWeight      string `json:"eWeight"`
	FDestination string `json:"fDestination"`
	FNumber      string `json:"fNumber"`
	FOrigin      string `json:"fOrigin"`
	Prefix       string `json:"prefix"`
	SDestination string `json:"sDestination"`
	SOrigin      string `json:"sOrigin"`
	SPieces      string `json:"sPieces"`
	SVolume      string `json:"sVolume"`
	SWeight      string `json:"sWeight"`
}

func parseLocalTime(eDateTime, eLocation string) string {
	loc, _ := time.LoadLocation(eLocation)
	eDate := eDateTime
	// orgForm needs to match format of eDate
	orgForm := "02-Jan-2006 15:04"
	isoForm := "2006-01-02T15:04:05-07:00"

	locTime, _ := time.ParseInLocation(orgForm, eDate, loc)
	return locTime.Format(isoForm)
}

func main() {
	// fmt.Println(parseLocalTime("20-May-2022 09:35", "Asia/Kolkata"))
	var (
		result []*map[string]interface{}
		err    error
		url    = "./airlines.xlsx"
		// select sheet name
		sheetName = "Coverage"
		// select only selected field
		// if you want to show all headers just passing nil or empty list
		// headers = []string{}
	)
	if result, err = excel2json.GetExcelFilePath(url, sheetName, nil); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	resArr := []Entry{}
	for _, val := range result {
		res, _ := json.Marshal(val)
		entry := Entry{}
		if err := json.Unmarshal(res, &entry); err == nil {
			resArr = append(resArr, entry)
		} else {
			fmt.Println("unable to parse entry")
		}
	}

	if res, err := json.Marshal(resArr); err == nil {
		_ = ioutil.WriteFile("test.json", res, 0o644)
	} else {
		fmt.Println("unable to write to file")
	}
}
