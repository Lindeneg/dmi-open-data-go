package cli

import "flag"

var (
	climateKey       = flag.String("c", "", "DMI climateData API key (defaults to $DMI_CLIMATE_API_KEY")
	metobsKey        = flag.String("m", "", "DMI metObs API key (defaults to $DMI_METOBS_API_KEY")
	filePath         = flag.String("p", "", "path to write json outputs to")
	dryRun           = flag.Bool("d", false, "do not write anything to disk")
	getClimateData   = flag.Bool("cd", false, "Run climateData method 'GetClimateData'")
	getStations      = flag.Bool("gs", false, "Run metObs method 'GetStations'")
	getObservations  = flag.Bool("go", false, "Run metObs method 'GetObservations'")
	getClosetStation = flag.Bool("cs", false, "Run metObs method 'GetClosetStation'")
	limit            = flag.Int("li", 10000, "Parameter: limit")
	offset           = flag.Int("of", 0, "Parameter: offset")
	ptype            = flag.String("ty", "", "Parameter: type")
	status           = flag.String("st", "", "Parameter: status")
	parameterID      = flag.String("pi", "", "Parameter: parameterId")
	stationId        = flag.String("si", "", "Parameter: stationId")
)
