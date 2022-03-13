package client

import (
	"fmt"
	"math"
	"time"
)

func isDefaultTime(t time.Time) bool {
	return t.Equal(time.Time{})
}

func constructTimeParam(fromTime time.Time, toTime time.Time) interface{} {
	ftd := isDefaultTime(fromTime)
	ttd := isDefaultTime(toTime)
	if ftd && ttd {
		return nil
	}
	if !ftd && ttd {
		return fromTime.Format(time.RFC3339)
	}
	if ftd && !ttd {
		return toTime.Format(time.RFC3339)
	}
	return fmt.Sprintf("%s/%s", fromTime.Format(time.RFC3339), toTime.Format(time.RFC3339))
}

func maybeParam(target interface{}, fallback interface{}, cmp interface{}) interface{} {
	if target == cmp {
		return fallback
	}
	return target
}

func distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	p := math.Pi / 180
	a := 0.5 - math.Cos((lat2-lat1)*p)/2.0 + math.Cos(lat1*p)*math.Cos(lat2*p)*(1.0-math.Cos((lon2-lon1)*p))/2.0
	d := 12742.0 * math.Asin(math.Sqrt(a))
	return d
}
