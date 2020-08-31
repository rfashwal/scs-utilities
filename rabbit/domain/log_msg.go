package domain

import "time"

type LogMsg struct {
	ProcessId   string    `json:"processId"`
	SensorId    string    `json:"sensorId"`
	Service     string    `json:"service"`
	LogMsg      string    `json:"logMsg"`
	PublishedOn time.Time `json:"publishedOn"`
}
