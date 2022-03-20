
Simple [DMI API](https://confluence.govcloud.dk/display/FDAPI/Danish+Meteorological+Institute+-+Open+Data) wrapper. Heavily inspired by [this](https://github.com/LasseRegin/dmi-open-data) Python library made by [Lasse Regin](https://github.com/LasseRegin) of [Woodsense](https://en.woodsense.dk/).

Get API keys [here](https://confluence.govcloud.dk/pages/viewpage.action?pageId=26476690).

The two API's used are **Meteorological Observation** (*metObs*) and **Climate Data** (*climateData*).

It is recommended to set the API keys as environmental variables:
`$DMI_CLIMATE_API_KEY` and `$DMI_METOBS_API_KEY`


##### Usage

`go install -i github.com/lindeneg/dmi-open-data-go/v2/cmd/dmi-open-data@latest`

or

```
git clone https://github.com/Lindeneg/dmi-open-data-go.git

cd dmi-open-data-go

go build cmd/dmi-open-data/dmi-open-data.go
```

```
$ dmi-open-data [...ARGS]

ARGS:

-c  --climate-key  [STR]   DMI climateData API key
-m  --metobs-key   [STR]   DMI metObs API key
-p  --file-path    [STR]   path to write json outputs to
-t  --type         [STR]   set parameter 'type'
-st --status       [STR]   set parameter 'status'
-l  --limit        [INT]   set parameter 'limit'
-f  --offset       [INT]   set parameter 'offset'
-pi --param-id     [INT]   set parameter 'parameterId'
-si --station-id   [INT]   set parameter 'stationId'
-a  --lat          [FLOAT] set latitude
-n  --lon          [FLOAT] set longitude
-cd --climate-data 	   run climateData method 'GetClimateData'
-s  --stations 	           run metObs method 'GetStations'
-o  --observations 	   run metObs method 'GetObservations'
-cs --closet-station 	   run metObs method 'GetClosetStation'
-d  --dry-run 	           do not write anything to disk
-sn --silent 	           do not print logs to stdout
-h  --help                 outputs the usage
```