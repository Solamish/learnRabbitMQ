package main

import "rabbitmq/RabbitMQ"

func main() {
	testTwo:=RabbitMQ.NewRabbitMQRouting("test","test_two")
	testTwo.ReceiveRouting()
}