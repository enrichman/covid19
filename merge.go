package main

import (
	"time"
)

func Merge(records []Record) World {
	world := NewWorld(records[0].Timeseries)

	for _, rec := range records {

		country, found := world.countryMap[rec.CountryID]
		if !found {
			country = NewCountry(rec.CountryID, rec.Country, records[0].Timeseries)
		}

		province, found := country.provinceMap[rec.ProvinceID]
		if !found {
			province = NewProvince(rec.ProvinceID, rec.Province, rec.Lat, rec.Lng, country, records[0].Timeseries)
		}

		// update data
		for i, ts := range rec.Timeseries {
			world.Timeseries[i].ValueDeath += ts.ValueDeath
			world.Timeseries[i].ValueConfirmed += ts.ValueConfirmed
			world.Timeseries[i].ValueRecovered += ts.ValueRecovered

			country.Timeseries[i].ValueDeath += ts.ValueDeath
			country.Timeseries[i].ValueConfirmed += ts.ValueConfirmed
			country.Timeseries[i].ValueRecovered += ts.ValueRecovered

			province.Timeseries[i].ValueDeath += ts.ValueDeath
			province.Timeseries[i].ValueConfirmed += ts.ValueConfirmed
			province.Timeseries[i].ValueRecovered += ts.ValueRecovered
		}

		world.Deaths += rec.Timeseries[len(rec.Timeseries)-1].ValueDeath
		world.Confirmed += rec.Timeseries[len(rec.Timeseries)-1].ValueConfirmed
		world.Recovered += rec.Timeseries[len(rec.Timeseries)-1].ValueRecovered

		country.Deaths += rec.Timeseries[len(rec.Timeseries)-1].ValueDeath
		country.Confirmed += rec.Timeseries[len(rec.Timeseries)-1].ValueConfirmed
		country.Recovered += rec.Timeseries[len(rec.Timeseries)-1].ValueRecovered

		province.Deaths += rec.Timeseries[len(rec.Timeseries)-1].ValueDeath
		province.Confirmed += rec.Timeseries[len(rec.Timeseries)-1].ValueConfirmed
		province.Recovered += rec.Timeseries[len(rec.Timeseries)-1].ValueRecovered

		world.countryMap[rec.CountryID] = country
		country.provinceMap[rec.ProvinceID] = province
	}

	for _, country := range world.countryMap {
		for _, province := range country.provinceMap {
			country.Provinces = append(country.Provinces, province)
		}
		world.Countries = append(world.Countries, country)
	}

	return world
}

func NewWorld(timeseries []TimeseriesData) World {
	world := World{}
	world.ID = "world"
	world.Name = "World"
	world.countryMap = make(map[string]Country)
	world.Countries = []Country{}
	world.Timeseries = []TimeseriesData{}
	world.LastUpdate = timeseries[len(timeseries)-1].time.Format(time.RFC3339)

	// initialize ts
	for _, ts := range timeseries {
		world.Timeseries = append(world.Timeseries, TimeseriesData{UnixTime: ts.UnixTime, time: ts.time})
	}

	return world
}

func NewCountry(countryID string, countryName string, timeseries []TimeseriesData) Country {
	country := Country{}
	country.ID = countryID
	country.Name = countryName
	country.provinceMap = make(map[string]Province)
	country.Provinces = []Province{}
	country.Timeseries = []TimeseriesData{}
	country.LastUpdate = timeseries[len(timeseries)-1].time.Format(time.RFC3339)

	// initialize ts
	for _, ts := range timeseries {
		country.Timeseries = append(country.Timeseries, TimeseriesData{UnixTime: ts.UnixTime, time: ts.time})
	}

	return country
}

func NewProvince(provinceID, provinceName string, lat, lng float64, country Country, timeseries []TimeseriesData) Province {
	province := Province{}
	province.LastUpdate = timeseries[len(timeseries)-1].time.Format(time.RFC3339)

	province.ID = provinceID
	if province.ID == "" {
		province.ID = country.ID
	}

	province.Name = provinceName
	if province.Name == "" {
		province.Name = country.Name
	}

	province.Lat = lat
	province.Lng = lng

	// initialize ts
	for _, ts := range timeseries {
		province.Timeseries = append(province.Timeseries, TimeseriesData{UnixTime: ts.UnixTime, time: ts.time})
	}

	return province
}
