# COVID-19 Json API

## Endpoints

The endpoints available are divided by `world`, `country` and `province`. For each of it a `full` and a `data` endpoint is availble.
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



## Data

https://github.com/CSSEGISandData/COVID-19
