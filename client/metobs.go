package client

import (
	"encoding/json"
	"errors"

	"github.com/lindeneg/dmi-open-data-go/v2/constants"
	"github.com/lindeneg/dmi-open-data-go/v2/request"
)

type metObsClient struct {
	r *request.RequestConfig
}

func NewMetObsClient(key string) metObsClient {
	return metObsClient{request.New(key, constants.APIMetObs)}
}

// Get collection of stations in unmarshaled geo-json format
// Accepts a range of parameters as specified in GetStationsConfig
func (c metObsClient) GetStations(cf GetStationsConfig) (getStationsResponse, error) {
	var (
		err  error
		res  getStationsResponse
		body []byte
	)
	query := request.Query{
		"stationId": maybeParam(cf.StationId, nil, 0),
		"status":    maybeParam(cf.Status, nil, ""),
		"type":      maybeParam(cf.Status, nil, ""),
		"limit":     maybeParam(cf.Limit, 10000, 0),
		"offset":    cf.Offset,
	}
	if body, err = c.r.Fetch("collections/station/items", query); err != nil {
		return res, err
	}
	if err = json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

// Get raw observations in unmarshaled geo-json format
// Accepts a range of parameters as specified in GetObservationsConfig
func (c metObsClient) GetObservations(cf GetObservationsConfig) (getObservationsResponse, error) {
	var (
		err  error
		res  getObservationsResponse
		body []byte
	)
	query := request.Query{
		"parameterId": maybeParam(string(cf.Parameter), nil, ""),
		"stationId":   maybeParam(cf.StationId, nil, 0),
		"datetime":    constructTimeParam(cf.FromTime, cf.ToTime),
		"limit":       maybeParam(cf.Limit, 10000, 0),
		"offset":      cf.Offset,
	}
	if body, err = c.r.Fetch("collections/observation/items", query); err != nil {
		return res, err
	}
	if err = json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

// Get closet station given a set of coordinates
// Accepts a range of parameters as specified in GetClosetStationConfig
func (c metObsClient) GetClosetStation(cf GetClosetStationConfig) (metObsStationFeature, error) {
	var (
		err  error
		res  getStationsResponse
		cs   metObsStationFeature
		dist float64 = 0xffffff
	)
	if res, err = c.GetStations(cf.Filter); err != nil {
		return cs, err
	}
	st := res.Features
	for _, s := range st {
		lat := s.Geometry.Coordinates[1]
		lon := s.Geometry.Coordinates[0]
		if curdist := distance(cf.Lat, cf.Lon, lat, lon); curdist < dist {
			dist = curdist
			cs = s
		}
	}
	return cs, nil
}

// Lists supported MetObsParameters
// MetObsParameters also accessible on 'constants'
// All MetObsParameters in 'constants' are prefixed 'M'
func (metObsClient) GetParameters() []constants.MetObsParameter {
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
