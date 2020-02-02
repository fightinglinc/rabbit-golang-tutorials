package main

import "../RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("logs", "one")
	rabbitmq.ReceiveRouting()
}