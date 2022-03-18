package cli

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/lindeneg/dmi-open-data-go/v2/client"
	"github.com/lindeneg/dmi-open-data-go/v2/file"
)

type actions struct {
	data map[string]interface{}
	err  map[string]interface{}
	done chan struct{}
	mux  sync.Mutex
	wg   sync.WaitGroup
}

func runActions() {
	actions := actions{}
	runs := 0
	c := client.NewClimateDataClient(*climateKey)
	m := client.NewMetObsClient(*metobsKey)
	if *getClimateData {
		go func() {
			colorize(ColorBlue, "running getClimateData")
			actions.mux.Lock()
			actions.data["getClimateData"], actions.err["getClimatedata"] = c.GetClimateData(client.GetClimateDataConfig{})
			actions.mux.Unlock()
			write(*filePath, actions.data["getClimateData"])
		}()
		runs++
	}
	if *getStations {
		colorize(ColorBlue, "running getStations")
		runs++
	}
	if *getObservations {
		colorize(ColorBlue, "running getObservations")
		runs++
	}
	if *getClosetStation {
		colorize(ColorBlue, "running getClosetStation")
		runs++
	}
	checkRuns(runs)

}

func checkRuns(runs int) {
	if runs == 0 {
		colorize(ColorRed, "please specify at least one action")
		flag.Usage()
		os.Exit(1)
	}
}

func write(filename string, data interface{}) {
	if *dryRun {
		return
	}
	name := fmt.Sprintf("%s/%s", *filePath, filename)
	colorize(ColorBlue, fmt.Sprintf("writing %s to disk", name))
	file.Write(name, data)
}
