package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitMq := RabbitMQ.NewSimplePattern("testSimple")
	rabbitMq.PublicSimple("hello, world")
	fmt.Println("发送成功")
}

 

