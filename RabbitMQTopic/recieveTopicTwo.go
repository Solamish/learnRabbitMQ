package main

import "rabbitmq/RabbitMQ"

func main()  {
	testOne:=RabbitMQ.NewRabbitMQTopic("exTopic","test.*.two")
	testOne.ReceiveTopic()
}
