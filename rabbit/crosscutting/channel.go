package crosscutting

import "errors"

func (r *RabbitConnector) initChannel() error {
	channel, err := r.Conn.Channel()
	if err != nil {
		return errors.New("cannot create channel")
	}
	r.Channel = channel
	return nil
}
