package client

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lindeneg/dmi-open-data-go/constants"
	"github.com/lindeneg/dmi-open-data-go/request"
)

type climateDataClient struct {
	request  *request.Request
	response *request.Response
}

func NewClimateDataClient(key string) climateDataClient {
	req, res := request.NewRequest(key, constants.APIClimateData)
	return climateDataClient{request: req, response: res}
}

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

// Get raw climate data in unmarshaled geo-json format
// Accepts a range of parameters as specified in GetClimateDataConfig.
func (client climateDataClient) GetClimateData(config GetClimateDataConfig) (getClimateDataResponse, error) {
	result := getClimateDataResponse{}
	query := request.Query{
		"parameterId":    maybeParam(string(config.Parameter), nil, ""),
		"stationId":      maybeParam(config.StationId, nil, 0),
		"datetime":       constructTimeParam(config.FromTime, config.ToTime),
		"timeResolution": maybeParam(config.TimeResolution, nil, ""),
		"limit":          maybeParam(config.Limit, 1000, 0),
		"offset":         config.Offset,
	}
	client.request.Get("collections/stationValue/items", query)
	if client.response.Err != nil {
		return result, client.response.Err
	}
	err := json.Unmarshal(client.response.Body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Lists supported ClimateDataParameters
// ClimateDataParameters are also accessible on 'constants'
// All ClimateDataParameters in 'constants' are prefixed 'C'
func (client climateDataClient) ListParameters() []constants.ClimateDataParameter {
	return constants.ClimateDataParameters
}

// Get ClimateDataParameter from string
func GetClimateDataParameter(s string) (constants.ClimateDataParameter, error) {
	for _, v := range constants.ClimateDataParameters {
		if string(v) == s {
			return v, nil
		}
	}
	return "", errors.New(s + " could not be matched with a ClimateDataParameter")
}
