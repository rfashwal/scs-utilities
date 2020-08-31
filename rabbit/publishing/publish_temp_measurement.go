package publishing

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/rfashwal/scs-utilities/rabbit/domain"
)

func (d *Publisher) PublishTemperatureMeasurement(
	routingKey,
	topic,
	service,
	sensorId string,
	value float64) {

	now := time.Now()

	msg := domain.TemperatureMeasurement{
		ProcessId:   uuid.New().String(),
		Value:       value,
		Service:     service,
		PublishedOn: now,
		SensorId:    sensorId,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Warn("could not marshal msg")
		return
	}
	d.Reset()
	d.Publish(
		topic,
		routingKey,
		string(bytes))

	d.Reset()
}
