package constants

type ApiType string

const (
	APIClimateData ApiType = "climateData"
	APIMetObs      ApiType = "metObs"
	BaseUrl        string  = "https://dmigw.govcloud.dk"
)

type GGeometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type GLinks struct {
	Href  string `json:"href"`
	Rel   string `json:"rel"`
	Type  string `json:"type"`
	Title string `json:"title"`
}
