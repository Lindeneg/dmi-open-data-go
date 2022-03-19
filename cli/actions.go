package cli

import (
	"fmt"
	"sync"

	"github.com/lindeneg/dmi-open-data-go/v2/client"
	"github.com/lindeneg/dmi-open-data-go/v2/file"
)

var wg sync.WaitGroup

func runActions(cf *config) {
	didSomething := false
	c := client.NewClimateDataClient(cf.climateKey)
	m := client.NewMetObsClient(cf.metobsKey)
	if cf.getClimateData {
		didSomething = true
		wg.Add(1)
		go func() {
			runAction(cf, "getClimateData", "climate_data", func() (interface{}, error) {
				return c.GetClimateData(getClimateDataConfig(cf))
			})
			wg.Done()
		}()
	}
	if cf.getStations {
		didSomething = true
		wg.Add(1)
		go func() {
			runAction(cf, "getStations", "stations", func() (interface{}, error) {
				return m.GetStations(getStationsConfig(cf))
			})
			wg.Done()
		}()
	}
	if cf.getObservations {
		didSomething = true
		wg.Add(1)
		go func() {
			runAction(cf, "getObservations", "observations", func() (interface{}, error) {
				return m.GetObservations(getObservationsConfig(cf))
			})
			wg.Done()
		}()
	}
	if cf.getClosetStation {
		didSomething = true
		wg.Add(1)
		go func() {
			mc := newConfig()
			if mc.lat == cf.lat || mc.lon == cf.lon {
				colorize(ColorRed, "getClosetStation error: please specify '--lat' and '--lon'")
			} else {
				runAction(cf, "getClosetStation", "closet_station", func() (interface{}, error) {
					return m.GetClosetStation(client.GetClosetStationConfig{Lat: cf.lat, Lon: cf.lon})
				})
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if !didSomething {
		errExit("please invoke at least one method")
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
		colorize(ColorBlue, fmt.Sprintf("%s: request successful", scope))
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
