#!/bin/bash

#global
curl -s -O https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Confirmed.csv
mv time_series_19-covid-Confirmed.csv csv/confirmed.csv

curl -s -O https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Deaths.csv
mv time_series_19-covid-Deaths.csv csv/deaths.csv

curl -s -O https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Recovered.csv
mv time_series_19-covid-Recovered.csv csv/recovered.csv

#italy
curl -s -O https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-json/dpc-covid19-ita-andamento-nazionale.json
mv dpc-covid19-ita-andamento-nazionale.json json/italy.json

curl -s -O https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-json/dpc-covid19-ita-regioni.json
mv dpc-covid19-ita-regioni.json json/regioni.json

curl -s -O https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-json/dpc-covid19-ita-province.json
mv dpc-covid19-ita-province.json json/province.json
