package main

//
type TrainData struct {
	TrainDataLink []TrainDataLink `json:"links"`
	Lines         []TrainLine     `json:"lines"`
}

type TrainDataLink struct {
	HttpLink string `json:"href"`
	Type     string `json:"type"`
	Rel      string `json:"rel"`
}

type TrainLine struct {
	Code           string           `json:"code"`
	Network        TrainLineNetwork `json:"network"`
	PhysicalModes  []TrainMode      `json:"physical_modes"`
	CommercialMode TrainMode        `json:"commercial_mode"`
	ClosingTime    string           `json:"closing_time"`
	OpeningTime    string           `json:"opening_time"`
	Id             string           `json:"id"`
	Name           string           `json:"name"`
	Routes         []TrainRoute     `json:"routes"`
}

type TrainLineNetwork struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TrainRoute struct {
	TrainDirection TrainDirection `json:"direction"`
	Name           string         `json:"name"`
	Id             string         `json:"id"`
	IsFrequent     string         `json:"is_frequence"`
	DirectionType  string         `json:"direction_type"`
}

type TrainDirection struct {
	EmbeddedType string        `json:"embedded_type"`
	StopARea     TrainStopArea `json:"stop_area"`
	Name         string        `json:"name"`
	Id           string        `json:"id"`
}

type TrainStopArea struct {
	Label       string             `json:"label"`
	Id          string             `json:"id"`
	Coordinates TrainStopAreaCoord `json:"coord"`
}

type TrainStopAreaCoord struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}

type TrainMode struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
