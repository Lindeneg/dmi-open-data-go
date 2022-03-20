package cli

import (
	"fmt"
	"os"
	"strconv"
)

type config struct {
	climateKey       string
	metobsKey        string
	filePath         string
	ptype            string
	status           string
	parameterId      string
	dryRun           bool
	getClimateData   bool
	getStations      bool
	getObservations  bool
	getClosetStation bool
	stationId        int64
	limit            int64
	offset           int64
	lat              float64
	lon              float64
}

func newConfig() *config {
	return &config{}
}

var args []string = os.Args[1:]

func (c *config) Parse() {
	checkEnv(&c.climateKey, &c.metobsKey)
	valid := 0
	if len(args) > 0 {
		var ignore interface{} = nil
		for i, arg := range args {
			if i == ignore {
				continue
			}
			e, v := parseArg(c, arg, i)
			valid += e
			ignore = v
		}
		getPath(&c.filePath)
		if valid > 0 {
			colorize(ColorBlue, "successfully parsed config")
			if c.dryRun {
				colorize(ColorBlue, "dry-run detected, will not write to disk")
			}
			return
		}
	}
	errExit("no valid arguments specified")
}

func getPath(filePath *string) {
	r, _ := os.Getwd()
	s := *filePath
	l := len(s) - 1
	if s != "" {
		if l > -1 && string(s[l]) == "/" {
			r += "/" + s[:l]
		} else {
			r += "/" + s
		}
	}
	*filePath = r
}

func checkEnv(c *string, m *string) {
	if *c == "" {
		*c = os.Getenv("DMI_CLIMATE_API_KEY")
	}
	if *m == "" {
		*m = os.Getenv("DMI_METOBS_API_KEY")
	}
}

func getNext(n int) string {
	if len(args) > n {
		return args[n]
	}
	return ""
}

func getInt(str string) int64 {
	r, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return r
}

func getFloat(str string) float64 {
	r, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return r
}

func parseArg(c *config, arg string, i int) (int, interface{}) {
	next := i + 1
	switch arg {

	case "-c", "--climate-key":
		c.climateKey = getNext(next)
		return 1, next

	case "-m", "--metobs-key":
		c.metobsKey = getNext(next)
		return 1, next

	case "-p", "--file-path":
		c.filePath = getNext(next)
		return 1, next

	case "-d", "--dry-run":
		c.dryRun = true
		return 1, nil

	case "-sn", "--silent":
		silent = true
		return 1, nil

	case "-cd", "--climate-data":
		c.getClimateData = true
		return 1, nil

	case "-s", "--stations":
		c.getStations = true
		return 1, nil

	case "-o", "--observations":
		c.getObservations = true
		return 1, nil

	case "-cs", "--closet-station":
		c.getClosetStation = true
		return 1, nil

	case "-l", "--limit":
		c.limit = getInt(getNext(next))
		return 1, next

	case "-f", "--offset":
		c.offset = getInt(getNext(next))
		return 1, next

	case "-a", "--lat":
		c.lat = getFloat(getNext(next))
		return 1, next

	case "-n", "--lon":
		c.lon = getFloat(getNext(next))
		return 1, next

	case "-t", "--type":
		c.ptype = getNext(next)
		return 1, next

	case "-st", "--status":
		c.status = getNext(next)
		return 1, next

	case "-pi", "--param-id":
		c.parameterId = getNext(next)
		return 1, next

	case "-si", "--station-id":
		c.stationId = getInt(getNext(next))
		return 1, next

	case "-h", "--help":
		usage()
		os.Exit(0)

	default:
		colorize(ColorRed, fmt.Sprintf("warning: argument '%s' is not recognized", arg))
	}
	return 0, nil
}
