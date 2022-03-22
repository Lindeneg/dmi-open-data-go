package cli

import (
	"fmt"
	"os"

	"github.com/lindeneg/dmi-open-data-go/pkg/client"
)

type Color string

const (
	ColorRed   Color = "\u001b[31m"
	ColorGreen Color = "\u001b[32m"
	ColorBlue  Color = "\u001b[36m"
	ColorReset Color = "\u001b[0m"
)

var (
	silent = false
)

func colorize(color Color, message string) {
	if !silent {
		fmt.Println(fmt.Sprintf("%s%s%s", string(color), message, string(ColorReset)))
	}
}

func errExit(m string) {
	colorize(ColorRed, "error: "+m+", run with '-h' or '--help' for usage\n")
	os.Exit(1)
}

func getClimateDataConfig(cf *config) client.GetClimateDataConfig {
	param, err := client.GetClimateDataParameter(cf.parameterId)
	if cf.parameterId != "" && err != nil {
		colorize(ColorRed, fmt.Sprintf("error: %v", err))
	}
	return client.GetClimateDataConfig{
		Parameter: param,
		StationId: int(cf.stationId),
		Limit:     int(cf.limit),
		Offset:    int(cf.offset),
	}
}

func getStationsConfig(cf *config) client.GetStationsConfig {
	return client.GetStationsConfig{
		Status:    cf.status,
		Type:      cf.ptype,
		StationId: int(cf.stationId),
		Limit:     int(cf.limit),
		Offset:    int(cf.offset),
	}
}

func getObservationsConfig(cf *config) client.GetObservationsConfig {
	param, err := client.GetMetObsParameter(cf.parameterId)
	if cf.parameterId != "" && err != nil {
		colorize(ColorRed, fmt.Sprintf("error: %v", err))
	}
	return client.GetObservationsConfig{
		Parameter: param,
		StationId: int(cf.stationId),
		Limit:     int(cf.limit),
		Offset:    int(cf.offset),
	}
}

func usage() {
	fmt.Println(`usage: dmi-open-data
src:   https://github.com/lindeneg/dmi-open-data-go

go run dmi-open-data-go [...ARGS]

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
-h  --help                 outputs the usage`)
}
