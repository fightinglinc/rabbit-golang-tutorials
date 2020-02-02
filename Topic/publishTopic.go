package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	first := RabbitMQ.NewRabbitMQTopic("topicExample", "topic.one")
	second := RabbitMQ.NewRabbitMQTopic("topicExample", "topic.two")

	for i := 0; i <= 10; i++ {
		first.PublishTopic("This is topic one: " + strconv.Itoa(i))
		second.PublishTopic("This is topic two: " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
