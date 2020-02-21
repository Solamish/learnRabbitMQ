package main

import "rabbitmq/RabbitMQ"

func main() {
	testOne:=RabbitMQ.NewRabbitMQRouting("test","test_one")
	testOne.ReceiveRouting()
}
