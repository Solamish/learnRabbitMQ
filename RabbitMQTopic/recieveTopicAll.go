package main

import "rabbitmq/RabbitMQ"

func main()  {
	testOne:=RabbitMQ.NewRabbitMQTopic("exTopic","#")
	testOne.ReceiveTopic()
}
