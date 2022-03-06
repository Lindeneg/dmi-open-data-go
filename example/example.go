package main

import (
	"fmt"
	"os"

	"github.com/lindeneg/dmi-open-data-go/client"
	"github.com/lindeneg/dmi-open-data-go/constants"
)

func main() {
	if len(os.Args) >= 2 {
		arg := os.Args[1]
		if arg == "metObs" {
			fmt.Println("METOBS API")
			metObsExample()
			return
		}
		if arg == "climateData" {
			fmt.Println("CLIMATE API")
			climateExample()
			return
		}
		fmt.Println("unknown argument:", arg)
		return
	}
	fmt.Println("no args specified.. please run with 'metObs' or 'climateData'")
}

func metObsExample() {
	c := client.NewMetObsClient(os.Getenv("DMI_METOBS_API_KEY"))

	obs, err := c.GetStations(client.GetStationsConfig{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Returned %d entries\n", obs.NumberReturned)
	if obs.NumberReturned > 0 {
		fmt.Printf("First entry: %v\n", obs.Features[0])
	}
}

func climateExample() {
	c := client.NewClimateDataClient(os.Getenv("DMI_CLIMATE_API_KEY"))

	obs, err := c.GetClimateData(client.GetClimateDataConfig{Parameter: constants.CMaxPressure})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Returned %d entries\n", obs.NumberReturned)
	if obs.NumberReturned > 0 {
		fmt.Printf("First entry: %v\n", obs.Features[0])
	}
}
