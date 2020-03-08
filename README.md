# COVID-19 Json API

There are a lots of data about the Coronavirus, but they are usually CSV a bit contrived to parse and read.  
This project was made to provide you some easy JSON endpoints, with the data already merged and parsed!

For example: https://enrichman.github.io/covid19/world/mainland-china/data.json

### Note

Please be aware that the project is in an early stage of development and the data/endpoints are subject to changes!

## Object

The object returned is a **Place**. A Place contains some metadata, an array of children (countries or provinces) and a timeseries field containing the deaths, confirmed and recovered cases.

```json
{
  "id": "world",
  "name": "World",
  "deaths": 3348,
  "confirmed": 97886,
  "recovered": 53797,
  "last_update": "2020-03-05T00:00:00Z",
  "ts": [],
  "countries": []
}
```

The **Timeseries** (**ts**) array contains the _timestamp_ (**t**) in Epoch with the values of the deaths (**d**), confirmed (**c**) and recovered (**r**) cases of the day.
```json
"ts": [{
  "t": 1579651200,
  "d": 17,
  "c": 555,
  "r": 28
}]
```
The **countries** array contains all the countries, that are Places as well.
```json
"countries": [{
  "id": "us",
  "name": "US",
  "deaths": 12,
  "confirmed": 221,
  "recovered": 8,
  "last_update": "2020-03-05T00:00:00Z"
}
```

## Endpoints

The endpoints available are divided by **world**, **country** and **province**.  
For each of it a **full** and a **data** endpoint is available.

The `data.json` endpoint contains the timeseries only of the specified location, while the `full.json` contains the timeseries of the children as well.

```
https://enrichman.github.io/covid19/world/<full/data>.json
https://enrichman.github.io/covid19/world/<country>/<full/data>.json
https://enrichman.github.io/covid19/world/<country>/<province>/full.json
```

For example:

https://enrichman.github.io/covid19/world/full.json  
https://enrichman.github.io/covid19/world/data.json

For a country:  
https://enrichman.github.io/covid19/world/mainland-china/full.json  
https://enrichman.github.io/covid19/world/mainland-china/data.json

For a province:  
https://enrichman.github.io/covid19/world/mainland-china/guangdong/data.json


## Projects

This section will contain a list of the projects using this APIs! :D

## Contributing

Contributions are welcome. Please open an issue or a PR explaining your needs or problems.

### Contributors

- [enrichman](https://github.com/enrichman)
- [esenac](https://github.com/esenac)

## Data

The data are extracted daily from the JHU CSSE:

 - https://github.com/CSSEGISandData/COVID-19  
 - https://systems.jhu.edu/research/public-health/ncov/


