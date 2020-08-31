package crosscutting

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func (r *RabbitConnector) Publish(topic, routingkey, body string) error {
	err := r.Channel.Publish(
		topic,
		routingkey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	if err != nil {
		log.Errorf(
			fmt.Sprintf("could not publish msg for topic: [%s] and routingkey: [%s]", topic, routingkey),
			err)
		return err
	}

	log.Infof(fmt.Sprintf("msg published to topic: [%s] and routingkey: [%s]", topic, routingkey))
	return nil
}
