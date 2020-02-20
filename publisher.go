package main

import (
	"fmt"
	"rabbitmq/mq"
)

func main() {
	rabbitMq := mq.NewSimplePattern("testSimple")
	rabbitMq.PublicSimple("hello, world")
	fmt.Println("发送成功")
}

 

