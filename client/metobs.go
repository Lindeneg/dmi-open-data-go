package client

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lindeneg/dmi-open-data-go/v2/constants"
	"github.com/lindeneg/dmi-open-data-go/v2/request"
)

type metObsClient struct {
	request  *request.Request
	response *request.Response
}

func NewMetObsClient(key string) metObsClient {
	req, res := request.NewRequest(key, constants.APIMetObs)
	return metObsClient{request: req, response: res}
}

type GetStationsConfig struct {
	Limit  int
	Offset int
}

/* I have a hard time extending structs in Go.
To my knowledge you can't destructure properties
from struct A into struct B, you'll have to compose
it but that means adding on an extra property

So I have to repeat some type definitions
until I understand whats up.. */

type getStationsResponse struct {
	Type     string `json:"type"`
	Features []struct {
		Geometry   constants.GGeometry `json:"geometry"`
		ID         string              `json:"id"`
		Type       string              `json:"type"`
		Properties struct {
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
		} `json:"properties"`
	} `json:"features"`
	TimeStamp      time.Time          `json:"timeStamp"`
	NumberReturned int                `json:"numberReturned"`
	Links          []constants.GLinks `json:"links"`
}

// Get collection of stations in unmarshaled geo-json format
// Accepts a range of parameters as specified in GetStationsConfig
func (client metObsClient) GetStations(config GetStationsConfig) (getStationsResponse, error) {
	result := getStationsResponse{}
	query := request.Query{
		"limit":  maybeParam(config.Limit, 10000, 0),
		"offset": config.Offset,
	}
	client.request.Get("collections/station/items", query)
	if client.response.Err != nil {
		return result, client.response.Err
	}
	err := json.Unmarshal(client.response.Body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
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

// Get raw observations in unmarshaled geo-json format
// Accepts a range of parameters as specified in GetObservationsConfig.
func (client metObsClient) GetObservations(config GetObservationsConfig) (getObservationsResponse, error) {
	result := getObservationsResponse{}
	query := request.Query{
		"parameterId": maybeParam(string(config.Parameter), nil, ""),
		"stationId":   maybeParam(config.StationId, nil, 0),
		"datetime":    constructTimeParam(config.FromTime, config.ToTime),
		"limit":       maybeParam(config.Limit, 10000, 0),
		"offset":      config.Offset,
	}
	client.request.Get("collections/observation/items", query)
	if client.response.Err != nil {
		return result, client.response.Err
	}
	err := json.Unmarshal(client.response.Body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// TODO GetClosetStation
// take tuple of coords
// call client.GetStations()
// use haversine to cmp dist
// return closet station

// Lists supported MetObsParameters
// MetObsParameters also accessible on 'constants'
// All MetObsParameters in 'constants' are prefixed 'M'
func (client metObsClient) GetParameters() []constants.MetObsParameter {
	return constants.MetObsParameters
}

// Get ClimateDataParameter from string
func GetMetObsParameter(s string) (constants.MetObsParameter, error) {
	for _, v := range constants.MetObsParameters {
		if string(v) == s {
			return v, nil
		}
	}
	return "", errors.New(s + " could not be matched with a MetObsParameter")
}
