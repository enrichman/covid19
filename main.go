package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

}

func printWorldFull(place World, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printWorldData(place World, out string) {
	place.Countries = nil
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printCountryFull(place Country, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}
func printCountryData(place Country, out string) {
	place.Provinces = nil
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printProvince(place Province, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}
