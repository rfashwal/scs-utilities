package observer

import (
	"github.com/rfashwal/scs-utilities/rabbit/crosscutting"
	"github.com/streadway/amqp"
)

type Observer struct {
	*crosscutting.RabbitConnector
	Queue amqp.Queue
}

func InitObserver(connection *amqp.Connection) (*Observer, error) {
	connector, err := crosscutting.InitConnector(connection)
	if err != nil {
		return nil, err
	}
	observer := &Observer{
		RabbitConnector: connector,
	}
	queue, err := observer.DeclareQueue()
	if err != nil {
		return nil, err
	}
	observer.Queue = queue
	return observer, nil
}

func (o *Observer) ShutdownObserver() {
	_ = o.Channel.Close()
}
