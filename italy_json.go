package main

import (
	"encoding/json"
	"os"
	"sort"
	"time"

	"github.com/gosimple/slug"
)

type RecordIT struct {
	CodiceRegione        int    `json:"codice_regione,omitempty"`
	DenominazioneRegione string `json:"denominazione_regione,omitempty"`

	CodiceProvincia        int    `json:"codice_provincia,omitempty"`
	DenominazioneProvincia string `json:"denominazione_provincia,omitempty"`
	SiglaProvincia         string `json:"sigla_provincia,omitempty"`

	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"long,omitempty"`

	DataStr string `json:"data,omitempty"`

	TimeSeriesDataIT
}

type TimeSeriesDataIT struct {
	Time int64 `json:"t,omitempty"`
	data time.Time
	Valori
}

type Valori struct {
	TotaleCasi                int `json:"totale_casi"`
	Deceduti                  int `json:"deceduti"`
	DimessiGuariti            int `json:"dimessi_guariti"`
	TotaleAttualmentePositivi int `json:"totale_attualmente_positivi"`
	NuoviAttualmentePositivi  int `json:"nuovi_attualmente_positivi"`
	TotaleOspedalizzati       int `json:"totale_ospedalizzati"`
	RicoveratiConSintomi      int `json:"ricoverati_con_sintomi"`
	TerapiaIntensiva          int `json:"terapia_intensiva"`
	IsolamentoDomiciliare     int `json:"isolamento_domiciliare"`
	Tamponi                   int `json:"tamponi"`
}

type Italy struct {
	Stato string `json:"stato,omitempty"`
	Valori
	regionMap  map[string]Regione
	Regioni    []Regione          `json:"regioni,omitempty"`
	Timeseries []TimeSeriesDataIT `json:"ts,omitempty"`
}

type Regione struct {
	CodiceRegione        int     `json:"codice_regione,omitempty"`
	DenominazioneRegione string  `json:"denominazione_regione,omitempty"`
	Lat                  float64 `json:"lat,omitempty"`
	Lng                  float64 `json:"long,omitempty"`
	Valori
	provinceMap map[int]Provincia
	Province    []Provincia        `json:"province,omitempty"`
	Timeseries  []TimeSeriesDataIT `json:"ts,omitempty"`
}

type Provincia struct {
	CodiceProvincia        int                `json:"codice_provincia,omitempty"`
	DenominazioneProvincia string             `json:"denominazione_provincia,omitempty"`
	SiglaProvincia         string             `json:"sigla_provincia,omitempty"`
	Lat                    float64            `json:"lat,omitempty"`
	Lng                    float64            `json:"long,omitempty"`
	TotaleCasi             int                `json:"totale_casi"`
	Timeseries             []TimeSeriesDataIT `json:"ts,omitempty"`
}

func parseItaly() Italy {
	f, _ := os.Open("json/italy.json")
	italyRecords := []RecordIT{}
	json.NewDecoder(f).Decode(&italyRecords)

	f, _ = os.Open("json/regioni.json")
	regionsRecords := []RecordIT{}
	json.NewDecoder(f).Decode(&regionsRecords)

	f, _ = os.Open("json/province.json")
	provinceRecords := []RecordIT{}
	json.NewDecoder(f).Decode(&provinceRecords)

	italy := Italy{
		Stato:      "Italia",
		Timeseries: []TimeSeriesDataIT{},
		regionMap:  make(map[string]Regione),
	}
	italy.Valori = italyRecords[len(italyRecords)-1].Valori

	// italia
	for _, record := range italyRecords {
		ts := TimeSeriesDataIT{}
		ts.Valori = record.Valori
		ts.data, _ = time.Parse("2006-01-02 15:04:05", record.DataStr)
		ts.Time = ts.data.Unix()

		italy.Timeseries = append(italy.Timeseries, ts)
	}

	// regioni
	for _, record := range regionsRecords {
		ts := TimeSeriesDataIT{}
		ts.Valori = record.Valori
		ts.data, _ = time.Parse("2006-01-02 15:04:05", record.DataStr)
		ts.Time = ts.data.Unix()

		regione, found := italy.regionMap[slug.Make(record.DenominazioneRegione)]
		if !found {
			regione = Regione{
				CodiceRegione:        record.CodiceRegione,
				DenominazioneRegione: record.DenominazioneRegione,
				Lat:                  record.Lat,
				Lng:                  record.Lng,
				provinceMap:          make(map[int]Provincia),
				Timeseries:           []TimeSeriesDataIT{},
			}
		}
		regione.Timeseries = append(regione.Timeseries, ts)

		italy.regionMap[slug.Make(record.DenominazioneRegione)] = regione
	}

	// province
	for _, record := range provinceRecords {
		ts := TimeSeriesDataIT{}
		ts.Valori = record.Valori
		ts.data, _ = time.Parse("2006-01-02 15:04:05", record.DataStr)
		ts.Time = ts.data.Unix()

		regione := italy.regionMap[slug.Make(record.DenominazioneRegione)]
		provincia, found := regione.provinceMap[record.CodiceProvincia]
		if !found {
			provincia = Provincia{
				CodiceProvincia:        record.CodiceProvincia,
				DenominazioneProvincia: record.DenominazioneProvincia,
				SiglaProvincia:         record.SiglaProvincia,
				Lat:                    record.Lat,
				Lng:                    record.Lng,
				TotaleCasi:             record.TotaleCasi,
				Timeseries:             []TimeSeriesDataIT{},
			}
		}

		provincia.Timeseries = append(provincia.Timeseries, ts)

		regione.provinceMap[record.CodiceProvincia] = provincia
		italy.regionMap[slug.Make(record.DenominazioneRegione)] = regione
	}

	for _, regione := range italy.regionMap {
		for _, provincia := range regione.provinceMap {
			provincia.TotaleCasi = provincia.Timeseries[len(provincia.Timeseries)-1].TotaleCasi
			regione.Province = append(regione.Province, provincia)
		}

		regione.Valori = regione.Timeseries[len(regione.Timeseries)-1].Valori
		italy.Regioni = append(italy.Regioni, regione)
	}

	sort.Slice(italy.Regioni, func(i, j int) bool {
		return italy.Regioni[i].TotaleCasi > italy.Regioni[j].TotaleCasi
	})
	for _, r := range italy.Regioni {
		sort.Slice(r.Province, func(i, j int) bool {
			return r.Province[i].TotaleCasi > r.Province[j].TotaleCasi
		})
	}

	return italy
}
