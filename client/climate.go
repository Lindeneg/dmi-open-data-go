package client

import (
	"encoding/json"
	"errors"

	"github.com/lindeneg/dmi-open-data-go/v2/constants"
	"github.com/lindeneg/dmi-open-data-go/v2/request"
)

type climateDataClient struct {
	r *request.RequestConfig
}

func NewClimateDataClient(k string) climateDataClient {
	return climateDataClient{request.New(k, constants.APIClimateData)}
}

// Get raw climate data in unmarshaled geo-json format
// Accepts a range of parameters as specified in GetClimateDataConfig.
func (c climateDataClient) GetClimateData(cf GetClimateDataConfig) (getClimateDataResponse, error) {
	var (
		err  error
		res  getClimateDataResponse
		body []byte
	)
	query := request.Query{
		"parameterId":    maybeParam(string(cf.Parameter), nil, ""),
		"stationId":      maybeParam(cf.StationId, nil, 0),
		"datetime":       constructTimeParam(cf.FromTime, cf.ToTime),
		"timeResolution": maybeParam(cf.TimeResolution, nil, ""),
		"limit":          maybeParam(cf.Limit, 1000, 0),
		"offset":         cf.Offset,
	}
	if body, err = c.r.Fetch("collections/stationValue/items", query); err != nil {
		return res, err
	}
	if err = json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

// Lists supported ClimateDataParameters
// ClimateDataParameters are also accessible on 'constants'
// All ClimateDataParameters in 'constants' are prefixed 'C'
func (climateDataClient) GetParameters() []constants.ClimateDataParameter {
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
