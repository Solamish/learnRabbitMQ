package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func NewRabbitMQTopic(exchangeName string, key string) *RabbitMQ {
	return NewRabbitMQ("",exchangeName, key)
}

// 生产
func (m *RabbitMQ)PublishTopic(message string) {
	err := m.channel.ExchangeDeclare(
		m.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	m.failOnErr(err, "failed to declare an exchange")
	//2.发送消息
	err = m.channel.Publish(
		m.Exchange,
		//要设置
		m.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 路由模式接受消息
func (m *RabbitMQ) ReceiveTopic() {
	//1.试探性创建交换机
	err := m.channel.ExchangeDeclare(
		m.Exchange,
		//交换机类型
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	m.failOnErr(err, "Failed to declare an exch"+
		"ange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := m.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	m.failOnErr(err, "Failed to declare a queue")

	//绑定队列到 exchange 中
	err = m.channel.QueueBind(
		q.Name,
		//需要绑定key
		m.Key,
		m.Exchange,
		false,
		nil)

	//消费消息
	messages, err := m.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("退出请按 CTRL+C\n")
	<-forever
}

