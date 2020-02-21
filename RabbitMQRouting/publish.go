package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	exchangeOne:=RabbitMQ.NewRabbitMQRouting("test","test_one")
	exchangeTwo:=RabbitMQ.NewRabbitMQRouting("test","test_two")
	for i := 0; i <= 10; i++ {
		exchangeOne.PublishRouting("Hello test one!" + strconv.Itoa(i))
		exchangeTwo.PublishRouting("Hello test Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

