package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	first := RabbitMQ.NewRabbitMQRouting("logs", "one")
	second := RabbitMQ.NewRabbitMQRouting("logs", "two")

	for i := 0; i <= 10; i++ {
		first.PublishRouting("Hello, the first one!" + strconv.Itoa(i))
		second.PublishRouting("Hello, the second one!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}