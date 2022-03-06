package client

import (
	"fmt"
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
