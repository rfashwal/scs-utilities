package crosscutting

import "github.com/streadway/amqp"

type RabbitConnector struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func InitConnector(connection *amqp.Connection) (*RabbitConnector, error) {
	rabbitConnector := &RabbitConnector{
		Conn: connection,
	}
	err := rabbitConnector.initChannel()
	return rabbitConnector, err
}

func (r *RabbitConnector) Reset() error {
	_ = r.Channel.Close()
	err := r.initChannel()
	if err != nil {
		return err
	}
	if r.Queue == nil {
		queue, err := r.DeclareQueue()
		if err != nil {
			return err
		}
		r.Queue = &queue
	}
	return nil
}
