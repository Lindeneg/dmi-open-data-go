package cli

import (
	"fmt"
	"sync"

	"github.com/lindeneg/dmi-open-data-go/v2/client"
	"github.com/lindeneg/dmi-open-data-go/v2/file"
)

type tch chan func()

type action struct {
	fn func()
}

type actions struct {
	cf *config
	wg *sync.WaitGroup
	ch tch
	as []action
	c  client.ClimateDataClient
	m  client.MetObsClient
}

func newActions(cf *config) *actions {
	var wg sync.WaitGroup
	ch := make(tch)
	c := client.NewClimateDataClient(cf.climateKey)
	m := client.NewMetObsClient(cf.metobsKey)
	ac := actions{ch: ch, wg: &wg, cf: cf, c: c, m: m}
	addActions(&ac)
	return &ac
}

func (ac *actions) runAll() {
	for _, a := range ac.as {
		go a.fn()
	}
	go func() {
		ac.wg.Wait()
		close(ac.ch)
	}()

	for fn := range ac.ch {
		fn()
	}

	if len(ac.as) < 1 {
		errExit("please invoke at least one method")
	}
}

func runAction(
	cf *config,
	ch tch,
	wg *sync.WaitGroup,
	scope string,
	filename string,
	fn func() (interface{}, error),
) {
	var (
		res  interface{}
		err  error
		name string
	)
	defer (*wg).Done()
	colorize(ColorBlue, fmt.Sprintf("%s: initialize request", scope))
	if res, err = fn(); err != nil {
		colorize(ColorRed, fmt.Sprintf("%s error: %v", scope, err))
	} else {
		colorize(ColorBlue, fmt.Sprintf("%s: request successful", scope))
		if !cf.dryRun {
			if name, err = write(cf.filePath, filename, res); err != nil {
				ch <- func() { colorize(ColorRed, fmt.Sprintf("%s error: %v", scope, err)) }
			} else {
				ch <- func() { colorize(ColorBlue, fmt.Sprintf("%s: wrote file '%s'", scope, name)) }
			}
		}
	}
}

func addActions(ac *actions) {
	cf, ch, wg := ac.cf, ac.ch, ac.wg
	if ac.cf.getClimateData {
		wg.Add(1)
		ac.as = append(ac.as, action{fn: func() {
			runAction(cf, ch, wg, "getClimateData", "climate_data", func() (interface{}, error) {
				return ac.c.GetClimateData(getClimateDataConfig(cf))
			})
		}})
	}
	if ac.cf.getStations {
		wg.Add(1)
		ac.as = append(ac.as, action{fn: func() {
			runAction(cf, ch, wg, "getStations", "stations", func() (interface{}, error) {
				return ac.m.GetStations(getStationsConfig(cf))
			})
		}})
	}
	if ac.cf.getObservations {
		wg.Add(1)
		ac.as = append(ac.as, action{fn: func() {
			runAction(cf, ch, wg, "getObservations", "observations", func() (interface{}, error) {
				return ac.m.GetObservations(getObservationsConfig(cf))
			})
		}})
	}
	if ac.cf.getClosetStation {
		mc := newConfig()
		if mc.lat == cf.lat || mc.lon == cf.lon {
			colorize(ColorRed, "getClosetStation error: please specify '--lat' and '--lon'")
		} else {
			wg.Add(1)
			ac.as = append(ac.as, action{fn: func() {
				if mc.lat == cf.lat || mc.lon == cf.lon {
					ch <- func() { colorize(ColorRed, "getClosetStation error: please specify '--lat' and '--lon'") }
					wg.Done()
				} else {
					runAction(cf, ch, wg, "getClosetStation", "closet_station", func() (interface{}, error) {
						res, dist, err := ac.m.GetClosetStation(client.GetClosetStationConfig{Lat: cf.lat, Lon: cf.lon})
						if err == nil {
							colorize(ColorBlue, fmt.Sprintf("getClosetStation distance: found station %dkm away", int(dist)))
						}
						return res, err
					})
				}
			}})
		}
	}
}

func write(path string, filename string, data interface{}) (string, error) {
	name := fmt.Sprintf("%s/%s.json", path, filename)
	err := file.Write(name, data)
	return name, err
}
