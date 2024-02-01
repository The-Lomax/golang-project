package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func test() {
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

func callGCPCrawler() error {
	testAwb := "023-88991431"

	fmt.Println("sending request to GCP")
	client := &http.Client{}

	resource := "/track"
	params := url.Values{}
	params.Add("providerType", "crawler")
	params.Add("awb", testAwb)

	u, _ := url.ParseRequestURI("https://api.tnt.cargoai.co")
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return err
	}
	req.Header.Add("X-API-Key", "im0dvz86lb903byaedoe9vc5v9hf2ofm")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return err
	}

	if resp == nil {
		fmt.Println("empty response from GCP")
		return err
	}

	if resp.StatusCode != 200 {
		fmt.Println("GCP response status code:", resp.StatusCode)
		return err
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	fmt.Printf("GCP response: %v\n", string(resBody))

	return nil
}

func main() {
	// test()
	fmt.Println(callGCPCrawler())
}
