package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

/**
订阅模式
*/

func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, "")
}

// 订阅模式生产
func (m *RabbitMQ) PublishPub(message string) {
	// 1.建立交换机
	err := m.channel.ExchangeDeclare(m.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	m.failOnErr(err, "failed to declare an exchange")
	// 2. 发送消息
	err = m.channel.Publish(
		m.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 订阅模式消费者代码
func (m *RabbitMQ) ReceievePub() {
	err := m.channel.ExchangeDeclare(m.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	q, err := m.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	m.failOnErr(err, "fail to declare a queue")

	// 绑定队列到交换机中
	err = m.channel.QueueBind(
		q.Name,
		// 在pub/sub模式下，这里的key为空
		"",
		m.Exchange,
		false,
		nil)

	// 消费消息
	messages, err := m.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,)

	forever := make(chan int)

	go func() {
		for d := range messages {
			log.Printf("receieved message :%s",d.Body)
		}
	}()
	fmt.Println("等待消息")
	<- forever
}
