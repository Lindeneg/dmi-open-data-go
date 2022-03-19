package cli

import (
	"fmt"

	"github.com/lindeneg/dmi-open-data-go/v2/client"
	"github.com/lindeneg/dmi-open-data-go/v2/file"
)

func runActions(cf *config) {
	c := client.NewClimateDataClient(cf.climateKey)
	//m := client.NewMetObsClient(*metobsKey)
	if cf.getClimateData {
		func() {
			runAction(cf, "getClimateData", "climate_data", func() (interface{}, error) {
				return c.GetClimateData(client.GetClimateDataConfig{})
			})
		}()
	}
	if cf.getStations {
		colorize(ColorBlue, "getStations: initialize")
	}
	if cf.getObservations {
		colorize(ColorBlue, "getObservations: initialize")
	}
	if cf.getClosetStation {
		colorize(ColorBlue, "getClosetStation: initialize")
	}

}

func runAction(cf *config, scope string, filename string, fn func() (interface{}, error)) {
	var (
		res  interface{}
		err  error
		name string
	)
	colorize(ColorBlue, fmt.Sprintf("%s: initialize", scope))
	if res, err = fn(); err != nil {
		colorize(ColorRed, fmt.Sprintf("%s error: %v", scope, err))
	} else {
		colorize(ColorGreen, fmt.Sprintf("%s: request successful", scope))
		if !cf.dryRun {
			if name, err = write(cf.filePath, filename, res); err != nil {
				colorize(ColorRed, fmt.Sprintf("%s error: %v", scope, err))
			} else {
				colorize(ColorGreen, fmt.Sprintf("%s: wrote file '%s'", scope, name))
			}
		}
	}
}

func write(path string, filename string, data interface{}) (string, error) {
	name := fmt.Sprintf("%s/%s.json", path, filename)
	err := file.Write(name, data)
	return name, err
}
