package main

import "rabbitmq/mq"

func main() {
	rabbitMq := mq.NewSimplePattern("testSimple")
	rabbitMq.ConsumeSimple()
}

 

