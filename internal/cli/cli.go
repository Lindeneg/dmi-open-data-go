package cli

import (
	"fmt"
	"time"
)

func Init() {
	start := time.Now()
	cf := newConfig()
	cf.Parse()
	newActions(cf).runAll()
	colorize(ColorGreen, fmt.Sprintf("dmi-open-data took %s", time.Since(start)))
}
