package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Date struct {
	Subject   string
	StartDate string
	EndDate   string
}

type Dates struct {
	Data []Date
}

var layout = "2006-01-02"

func main() {

	jsonFile, err := os.Open("dates.json")
	if err != nil {
		fmt.Println(err)
	}

	var data Dates
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &data.Data)

	for _, date := range data.Data {
		if date.StartDate == "" {
			continue;
		}

		startDate, startDateErr := time.Parse(layout, date.StartDate)
		if startDateErr != nil {
			fmt.Println(startDateErr)
		}

		endDate := time.Now()
		var endDateErr error
		if date.EndDate != "" {
			endDate, endDateErr = time.Parse(layout, date.EndDate)
			if endDateErr != nil {
				fmt.Println(endDateErr)
			}
		}

		year, month, day, hour, min, sec := diff(startDate, endDate)
		duration := endDate.Sub(startDate)

		fmt.Printf("%s: %d days; %d years, %d months, %d days, %d hours, %d mins and %d seconds.",
			date.Subject, int(duration.Hours()/24), year, month, day, hour, min, sec)
		fmt.Println()
	}
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	startYear, startMonth, startDay := a.Date()
	endYear, endMonth, endDay := b.Date()

	startHour, startMin, startSec := a.Clock()
	endHour, endMin, endSec := b.Clock()

	year = int(endYear - startYear)
	month = int(endMonth - startMonth)
	day = int(endDay - startDay)
	hour = int(endHour - startHour)
	min = int(endMin - startMin)
	sec = int(endSec - startSec)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(startYear, startMonth, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
