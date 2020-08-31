package publishing

import (
	"github.com/rfashwal/scs-utilities/rabbit/crosscutting"
	"github.com/streadway/amqp"
)

type Publisher struct {
	*crosscutting.RabbitConnector
}

func InitPublisher(connection *amqp.Connection) (*Publisher, error) {
	connector, err := crosscutting.InitConnector(connection)
	if err != nil {
		return nil, err
	}
	dialer := &Publisher{
		RabbitConnector: connector,
	}
	return dialer, nil
}
