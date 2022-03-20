package client

import (
	"time"

	"github.com/lindeneg/dmi-open-data-go/v2/internal/constants"
)

type GetStationsConfig struct {
	StationId int
	Status    string
	Type      string
	Limit     int
	Offset    int
}

type metObsStation struct {
	Country           string   `json:"country"`
	Name              string   `json:"name"`
	Owner             string   `json:"owner"`
	ParameterID       []string `json:"parameterId"`
	StationID         string   `json:"stationId"`
	Status            string   `json:"status"`
	TimeCreated       string   `json:"timeCreated"`
	TimeOperationFrom string   `json:"timeOperationFrom"`
	TimeOperationTo   string   `json:"timeOperationTo"`
	TimeUpdated       string   `json:"timeUpdated"`
	TimeValidFrom     string   `json:"timeValidFrom"`
	TimeValidTo       string   `json:"timeValidTo"`
	Type              string   `json:"type"`
}

type metObsStationFeature struct {
	Geometry   constants.GGeometry `json:"geometry"`
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Properties metObsStation       `json:"properties"`
}

type getStationsResponse struct {
	Type           string                 `json:"type"`
	Features       []metObsStationFeature `json:"features"`
	TimeStamp      time.Time              `json:"timeStamp"`
	NumberReturned int                    `json:"numberReturned"`
	Links          []constants.GLinks     `json:"links"`
}

type GetObservationsConfig struct {
	Parameter constants.MetObsParameter
	StationId int
	FromTime  time.Time
	ToTime    time.Time
	Limit     int
	Offset    int
}

type getObservationsResponse struct {
	Type     string `json:"type"`
	Features []struct {
		Geometry   constants.GGeometry `json:"geometry"`
		ID         string              `json:"id"`
		Type       string              `json:"type"`
		Properties struct {
			Created     time.Time `json:"created"`
			Observed    time.Time `json:"observed"`
			ParameterID string    `json:"parameterId"`
			StationID   string    `json:"stationId"`
			Value       float64   `json:"value"`
		} `json:"properties"`
	} `json:"features"`
	TimeStamp      time.Time          `json:"timeStamp"`
	NumberReturned int                `json:"numberReturned"`
	Links          []constants.GLinks `json:"links"`
}

type GetClosetStationConfig struct {
	Lat    float64
	Lon    float64
	Filter GetStationsConfig
}
