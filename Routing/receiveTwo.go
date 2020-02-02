package main

import "../RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("logs", "two")
	rabbitmq.ReceiveRouting()
}