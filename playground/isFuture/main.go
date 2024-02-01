package main

import (
	"fmt"
	"time"
)

func isFuture(date string) {
	eventDateInUTC, err := time.Parse(time.RFC3339, date)
	isFuture := time.Now().Before(eventDateInUTC)
	fmt.Printf("time in local: %+v\ntime in UTC: %+v\ntime now in UTC: %v\nisFuture: %v\nerr: %v", eventDateInUTC, eventDateInUTC.UTC(), time.Now().UTC(), isFuture, err)
}

func isNFDOnTime(plannedNFD, actualNFD string) {
	planned, err1 := time.Parse(time.RFC3339, plannedNFD)
	actual, err2 := time.Parse(time.RFC3339, actualNFD)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	isNFDOnTime := planned.Before(actual)
	fmt.Printf("date1 in local: %+v,\tdate1 in UTC: %+v,\terr: %v\n", planned, planned.UTC(), err1)
	fmt.Printf("date2 in local: %+v,\tdate2 in UTC: %+v,\terr: %v\n", actual, actual.UTC(), err2)
	fmt.Printf("isBefore: %v\n", isNFDOnTime)
	fmt.Printf("NFDdelay: %v\n", actual.Sub(planned).Minutes())
	fmt.Printf("%v\n", actual.In(loc).Format(time.RFC3339))
}

func main() {
	// eventDate := "2023-06-01T01:00:00Z"
	// isFuture(eventDate)

	//date1 := "2023-05-14T10:00:01-07:00"
	//date2 := "2023-05-14T11:30:01Z"
	//isNFDOnTime(date1, date2)

	tz := "America/New_York"
	loc, _ := time.LoadLocation(tz)
	eventTime := "2023-11-13T21:29:00Z"
	layout := "2006-01-02T15:04:05Z"
	t2, _ := time.Parse(time.RFC3339, eventTime)
	t1, _ := time.ParseInLocation(layout, t2.Format(time.RFC3339), loc)
	fmt.Println(t1.Format(time.RFC3339), t2.Format(time.RFC3339))
}
