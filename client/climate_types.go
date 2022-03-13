package client

import (
	"time"

	"github.com/lindeneg/dmi-open-data-go/v2/constants"
)

type GetClimateDataConfig struct {
	Parameter      constants.ClimateDataParameter
	StationId      int
	FromTime       time.Time
	ToTime         time.Time
	TimeResolution string
	Limit          int
	Offset         int
}

type getClimateDataResponse struct {
	Type     string `json:"type"`
	Features []struct {
		Geometry   constants.GGeometry `json:"geometry"`
		ID         string              `json:"id"`
		Type       string              `json:"type"`
		Properties struct {
			CalculatedAt   time.Time `json:"calculatedAt"`
			Created        time.Time `json:"created"`
			From           time.Time `json:"from"`
			ParameterID    string    `json:"parameterId"`
			QcStatus       string    `json:"qcStatus"`
			StationID      string    `json:"stationId"`
			TimeResolution string    `json:"timeResolution"`
			To             time.Time `json:"to"`
			Validity       bool      `json:"validity"`
			Value          float64   `json:"value"`
		} `json:"properties"`
	} `json:"features"`
	TimeStamp      time.Time          `json:"timeStamp"`
	NumberReturned int                `json:"numberReturned"`
	Links          []constants.GLinks `json:"links"`
}
