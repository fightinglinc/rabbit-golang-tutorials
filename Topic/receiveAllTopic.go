package main

import "../RabbitMQ"

func main() {
	// Use "#", should get all the messages
	rabbitmq := RabbitMQ.NewRabbitMQTopic("topicExample", "#")
	rabbitmq.ReceiveTopic()
}
