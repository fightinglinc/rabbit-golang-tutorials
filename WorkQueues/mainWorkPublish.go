package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")

	for i := 0; i <= 100; i++ {
		// strconv.Itoa() method can convert int value to string
		rabbitmq.PublishSimple("Hello imooc!" + strconv.Itoa(i))
		// 1 * time.Second => 1s
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
