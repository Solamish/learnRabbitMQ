package main

import "rabbitmq/RabbitMQ"

func main() {
	rabbitMq := RabbitMQ.NewSimplePattern("testSimple")
	rabbitMq.ConsumeSimple()
}

 

