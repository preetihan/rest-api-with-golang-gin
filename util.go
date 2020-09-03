package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func convertIntToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func convertFloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 0, 64)
}

func convertStringToFloat(num string) (result float64, err error) {
	result, err = strconv.ParseFloat(num, 8)
	return
}

func convertStringToDate(d string) time.Time {
	parsedDate, err := time.Parse("02-01-2006", d)
	if err != nil {
		log.Println("convertStringToDate error: ", err)
	}
	return parsedDate
}

func isDateInRange(startDate, endDate string) bool {
	parsedStartDate := convertStringToDate(startDate)
	parsedEndDate := convertStringToDate(endDate)
	dateRange := parsedEndDate.Sub(parsedStartDate).Hours() / 24
	log.Println("Date Range: " + convertFloatToString(dateRange) + " day(s)")
	return dateRange > -1 && dateRange <= 90
}

func retrieveDateRange(startDate, endDate string) func() time.Time {
	parsedStartDate := convertStringToDate(startDate)
	parsedEndDate := convertStringToDate(endDate)

	return func() time.Time {
		if parsedStartDate.After(parsedEndDate) {
			return time.Time{}
		}
		date := parsedStartDate
		parsedStartDate = parsedStartDate.AddDate(0, 1, 0)
		return date
	}
}

func getCurrentDate() string {
	location, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(location).Format("02012006")
}

func getCurrentDateTime() string {
	location, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(location).Format("02-01-2006 15:04:05")
}
