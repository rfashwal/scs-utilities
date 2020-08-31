package crosscutting

import (
	"github.com/streadway/amqp"
)

func (r *RabbitConnector) BindQueue(queue amqp.Queue, routingKey, exchangeName string) error {
	return r.Channel.QueueBind(queue.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)
}
