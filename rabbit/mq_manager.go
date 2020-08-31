package rabbit

import (
	"github.com/rfashwal/scs-utilities/rabbit/observer"
	"github.com/rfashwal/scs-utilities/rabbit/publishing"
	"github.com/streadway/amqp"
)

type MQManager interface {
	CloseConnection() error
	InitPublisher() (*publishing.Publisher, error)
	InitObserver() (*observer.Observer, error)
}

type rabbitMQManager struct {
	conn *amqp.Connection
}

func NewRabbitMQManager(url string) (MQManager, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return &rabbitMQManager{conn: connection}, nil
}

func (r *rabbitMQManager) CloseConnection() error {
	return r.conn.Close()
}

func (r *rabbitMQManager) InitPublisher() (*publishing.Publisher, error) {
	return publishing.InitPublisher(r.conn)
}

func (r *rabbitMQManager) InitObserver() (*observer.Observer, error) {
	return observer.InitObserver(r.conn)
}
