package main

import (
	"fmt"
	"time"
)

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
	fmt.Println(parseLocalTime("20-May-2022 09:35", "Asia/Kolkata"))
}
