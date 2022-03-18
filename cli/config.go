package cli

import (
	"flag"
	"os"
)

func rmFilePathTailSlash() {
	s := *filePath
	l := len(s) - 1
	if l > -1 && string(s[l]) == "/" {
		*filePath = s[:l]
	}
}

func parseConfig() {
	flag.Parse()
	if *climateKey == "" {
		*climateKey = os.Getenv("DMI_CLIMATE_API_KEY")
	}
	if *metobsKey == "" {
		*metobsKey = os.Getenv("DMI_METOBS_API_KEY")
	}
	rmFilePathTailSlash()
}
