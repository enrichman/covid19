package main

import (
	"encoding/json"
	"io/ioutil"
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

	printFullWorld(world, "world_full.json")
	printFullCountry(world.Countries[0], "country_full.json")
}

func printFullWorld(place World, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}

func printFullCountry(place Country, out string) {
	b, _ := json.Marshal(place)
	_ = ioutil.WriteFile(out, b, 0644)
}
