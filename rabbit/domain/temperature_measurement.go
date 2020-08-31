package domain

import "time"

type TemperatureMeasurement struct {
	ProcessId   string    `json:"processId"`
	SensorId    string    `json:"sensorId"`
	Service     string    `json:"service"`
	Value       float64   `json:"value"`
	PublishedOn time.Time `json:"publishedOn"`
}
