package main

import "../RabbitMQ"


func main()  {
	// Should receive only topic.two
	rabbitmq := RabbitMQ.NewRabbitMQTopic("topicExample", "*.two")
	rabbitmq.ReceiveTopic()
}

