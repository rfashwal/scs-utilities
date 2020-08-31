package crosscutting

func (r *RabbitConnector) DeclareTopicExchange(exchangeName string) error {
	return r.Channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
