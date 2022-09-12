package support

import (
	"log"
	"math"
	"time"
)

const (
	layoutDate = time.RFC3339
)

func TimeDiffDatesInMinute(data1, data2 time.Time) float64 {
	diff := data1.Sub(data2)
	return math.Round(diff.Minutes())
}

func TimeDiffDatesInSecond(data1, data2 time.Time) float64 {
	diff := data2.Sub(data1)
	return math.Round(diff.Seconds())
}

// Parse time to string with layout "2006-01-02T15:04:05Z07:00"
func ParseTimeFormatToSrt(t time.Time) string {
	return t.Format(time.RFC3339)
}

// Parse string to time with layout "2006-01-02 15:04:05 -0700 -07"
func ParseTimeFormatToTime(t string) (date time.Time, err error) {
	date, err = time.Parse(layoutDate, t)
	return date, err
}

// Now time string with layout "2006-01-02 15:04:05 -0700"
func NewTimeStr() string {
	return time.Now().Format(layoutDate)
}

func TimeNowFormat() string {
	layout := "2006-01-02 15:04:05 -07:00"
	return time.Now().Format(layout)
}

func ParseStringToTime(date string) time.Time {
	layout := "2006-01-02 15:04:05 -07:00"
	parsed, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}

func TimeFormatTime(date string) string {
	layout := "2006-01-02 15:04:05"

	parsed, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return parsed.String()
}

func DateToTimeZone(data string) string {

	dataPosicaoTime, _ := time.Parse("2006-01-02 15:04:05", data)
	location, _ := time.LoadLocation("America/Sao_Paulo")
	data = dataPosicaoTime.In(location).Format("2006-01-02 15:04:05 -07:00")
	return data
}
