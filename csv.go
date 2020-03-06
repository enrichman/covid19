package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Record struct {
	Province   string
	ProvinceID string
	Country    string
	CountryID  string
	Lat        float64
	Lng        float64
	Type       string
	Status     string
	LastUpdate time.Time
	Timeseries []TimeseriesData
}

type TimeseriesData struct {
	UnixTime       int64 `json:"t"`
	time           time.Time
	ValueDeath     int `json:"d"`
	ValueConfirmed int `json:"c"`
	ValueRecovered int `json:"r"`
}

func GetRecords() []Record {
	records := getDeaths()
	records = append(records, getConfirmed()...)
	records = append(records, getRecovered()...)
	return records
}

func getDeaths() []Record {
	f, err := os.Open("csv/deaths.csv")
	if err != nil {
		panic(err)
	}
	return parseRecords(f, "deaths")
}

func getConfirmed() []Record {
	f, err := os.Open("csv/confirmed.csv")
	if err != nil {
		panic(err)
	}
	return parseRecords(f, "confirmed")
}

func getRecovered() []Record {
	f, err := os.Open("csv/recovered.csv")
	if err != nil {
		panic(err)
	}
	return parseRecords(f, "recovered")
}

func parseRecords(reader io.Reader, recordType string) []Record {
	csvReader := csv.NewReader(reader)
	records, _ := csvReader.ReadAll()

	dates := []time.Time{}
	for i := 4; i < len(records[0]); i++ {
		dateString := records[0][i]
		date, _ := time.Parse("1/_2/06", dateString)
		dates = append(dates, date)
	}

	recordsParsed := []Record{}
	for i := 1; i < len(records); i++ {
		lat, _ := strconv.ParseFloat(records[i][2], 64)
		lng, _ := strconv.ParseFloat(records[i][3], 64)

		rec := Record{
			Province:   strings.TrimSpace(records[i][0]),
			ProvinceID: slug.Make(records[i][0]),
			Country:    strings.TrimSpace(records[i][1]),
			CountryID:  slug.Make(records[i][1]),
			Status:     recordType,
			Timeseries: []TimeseriesData{},
			LastUpdate: dates[len(dates)-1],
			Lat:        lat,
			Lng:        lng,
		}

		for j := 4; j < len(records[i]); j++ {
			value, _ := strconv.Atoi(records[i][j])
			tsData := TimeseriesData{
				UnixTime: dates[j-4].Unix(),
				time:     dates[j-4],
			}

			switch recordType {
			case "deaths":
				tsData.ValueDeath = value
			case "confirmed":
				tsData.ValueConfirmed = value
			case "recovered":
				tsData.ValueRecovered = value
			}

			rec.Timeseries = append(rec.Timeseries, tsData)
		}
		recordsParsed = append(recordsParsed, rec)
	}

	return recordsParsed
}
