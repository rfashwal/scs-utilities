package crosscutting

import (
	"github.com/streadway/amqp"
)

func (r *RabbitConnector) DeclareQueue() (amqp.Queue, error) {
	queue, err := r.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return amqp.Queue{}, err
	}

	return queue, nil
}
