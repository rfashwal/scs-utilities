package observer

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func (o *Observer) Observe() <-chan amqp.Delivery {
	deliveries, err := o.Channel.Consume(
		o.Queue.Name,
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)

	if err != nil {
		log.Panic(fmt.Sprintf("could not start observing with queue: [%s] ", o.Queue.Name))
	}
	return deliveries
}
