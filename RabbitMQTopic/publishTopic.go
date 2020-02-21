package main

import (
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
	"fmt"
)

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("exTopic","test.topic.one")
	imoocTwo:=RabbitMQ.NewRabbitMQTopic("exTopic","test.topic.two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishTopic("Hello test topic one!" + strconv.Itoa(i))
		imoocTwo.PublishTopic("Hello test topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	
}
