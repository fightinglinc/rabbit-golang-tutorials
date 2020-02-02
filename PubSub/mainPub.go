package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishPub("Number " + strconv.Itoa(i))
		fmt.Printf("This is number %v pieces of information "+
			"in the Publish/Subscribe Pattern\n", strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
