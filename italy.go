package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gosimple/slug"
)

var italianRegions map[string]Record

func getRegions() []Record {
	f, err := os.Open("csv/italy/dpc-covid19-ita-regioni.csv")
	if err != nil {
		panic(err)
	}
	return parseItalianRegions(f)
}

func parseItalianRegions(reader io.Reader) []Record {
	italianRegions = make(map[string]Record)
	csvReader := csv.NewReader(reader)
	records, _ := csvReader.ReadAll()

	for i := 1; i < len(records); i++ {
		provinceID := slug.Make(records[i][3]) // denominazione_regione

		confirmed, _ := strconv.Atoi(records[i][10])                  // totale_attualmente_positivi
		death, _ := strconv.Atoi(records[i][13])                      // deceduti
		recovered, _ := strconv.Atoi(records[i][12])                  // dimessi_guariti
		tsTime, _ := time.Parse("2006-01-02 15:04:05", records[i][0]) // data
		lat, _ := strconv.ParseFloat(records[i][4], 64)               // lat
		lng, _ := strconv.ParseFloat(records[i][5], 64)               // long
		tsData := TimeseriesData{
			ValueConfirmed: confirmed,
			ValueDeath:     death,
			ValueRecovered: recovered,
			time:           tsTime,
			UnixTime:       tsTime.Unix(),
		}

		var rec Record
		existing, ok := italianRegions[provinceID]
		if ok {
			rec = existing
			rec.Timeseries = append(rec.Timeseries, tsData)
			rec.LastUpdate = tsTime
		} else {
			rec = Record{
				LastUpdate: tsTime,
				Province:   records[i][3], // denominazione_regione
				ProvinceID: provinceID,
				Country:    "Italy",
				CountryID:  "italy",
				Lat:        lat,
				Lng:        lng,
				Timeseries: []TimeseriesData{
					tsData,
				},
			}
		}
		italianRegions[provinceID] = rec
	}
	recordsParsed := []Record{}
	for _, r := range italianRegions {
		recordsParsed = append(recordsParsed, r)
	}
	return recordsParsed
}
