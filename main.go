package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gosimple/slug"
)

type Place struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	Deaths     int              `json:"deaths"`
	Confirmed  int              `json:"confirmed"`
	Recovered  int              `json:"recovered"`
	LastUpdate string           `json:"last_update"`
	Timeseries []TimeseriesData `json:"ts,omitempty"`
}

type World struct {
	Place
	countryMap map[string]Country
	Countries  []Country `json:"countries,omitempty"`
}

type Country struct {
	Place
	provinceMap map[string]Province
	Provinces   []Province `json:"provinces,omitempty"`
}

type Province struct {
	countryID string
	Lat       float64
	Lng       float64
	Place
}

func main() {
	recordsParsed := GetRecords()

	world := Merge(recordsParsed)

	os.Mkdir("world", os.ModePerm)

	for _, country := range world.Countries {
		for _, province := range country.Provinces {
			os.MkdirAll(fmt.Sprintf("world/%s/%s", country.ID, province.ID), os.ModePerm)
			printProvince(province, fmt.Sprintf("world/%s/%s/data.json", country.ID, province.ID))
		}
		printCountryFull(country, fmt.Sprintf("world/%s/full.json", country.ID))
		printCountryData(country, fmt.Sprintf("world/%s/data.json", country.ID))
	}
	printWorldFull(world, "world/full.json")
	printWorldData(world, "world/data.json")

	// italy data
	italy := parseItaly()

	for _, regione := range italy.Regioni {
		regID := slug.Make(regione.DenominazioneRegione)

		for _, provincia := range regione.Province {
			provID := slug.Make(provincia.SiglaProvincia)
			if provID == "" {
				provID = "xx"
			}
			os.MkdirAll(fmt.Sprintf("local/italy/%s/%s", regID, provID), os.ModePerm)
			printProvincia(provincia, fmt.Sprintf("local/italy/%s/%s/data.json", regID, provID))
		}

		printRegioneFull(regione, fmt.Sprintf("local/italy/%s/full.json", regID))
		printRegioneData(regione, fmt.Sprintf("local/italy/%s/data.json", regID))
	}

	printItalyFull(italy, "local/italy/full.json")
	printItalyData(italy, "local/italy/data.json")
}

func printWorldFull(place World, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printWorldData(place World, out string) {
	for i := range place.Countries {
		place.Countries[i].Timeseries = nil
		place.Countries[i].Provinces = nil
	}

	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printCountryFull(place Country, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}
func printCountryData(place Country, out string) {
	for i := range place.Provinces {
		place.Provinces[i].Timeseries = nil
	}

	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printProvince(place Province, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

// LOCAL - ITALY

func printItalyFull(italy Italy, out string) {
	b, _ := json.Marshal(italy)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printItalyData(italy Italy, out string) {
	for i := range italy.Regioni {
		italy.Regioni[i].Timeseries = nil
		italy.Regioni[i].Province = nil
	}

	b, _ := json.Marshal(italy)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printRegioneFull(regione Regione, out string) {
	b, _ := json.Marshal(regione)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printRegioneData(regione Regione, out string) {
	for i := range regione.Province {
		regione.Province[i].Timeseries = nil
	}

	b, _ := json.Marshal(regione)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printProvincia(provincia Provincia, out string) {
	b, _ := json.Marshal(provincia)
	_ = ioutil.WriteFile(out, b, 0644)
}
