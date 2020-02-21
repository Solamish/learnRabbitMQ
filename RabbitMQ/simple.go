package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

/**
  simple 模式
*/

// url 格式 amqp://账号：密码@rabbit的服务器地址：端口号/vhost
const MQURL = "amqp://guest:guest@127.0.0.1:5672"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key   string
	MqUrl string
}

// 创建RabbitMq示例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitMq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MqUrl:     MQURL,
	}
	var err error
	rabbitMq.conn, err = amqp.Dial(rabbitMq.MqUrl)
	rabbitMq.failOnErr(err, "创建连接错误")
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	rabbitMq.failOnErr(err, "获取channel失败")
	return rabbitMq
}

func (m *RabbitMQ) Destroy() {
	m.channel.Close()
	m.conn.Close()
}

func (m *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
	return
}

// 创建simple模式的mq
func NewSimplePattern(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

func (m *RabbitMQ) PublishSimple(message string) {
	// 1.申请队列
	_, err := m.channel.QueueDeclare(m.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否排他
		false,
		// 是否阻塞
		false,
		// 额外参数
		nil)
	if err != nil {
		fmt.Println(err)
	}
	// 2.发送消息到队列
	_ = m.channel.Publish(
		m.Exchange,
		m.QueueName,
		// 如果设置为true， 会根据exchange的类型和routekey规则寻找队列，
		// 如果无法找到符合规则的队列，那么发送的消息会返回给发送者
		false,
		// 如果设置为true, 当exchange发送消息到消息队列后发现该队列上没有绑定消费者，则会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (m *RabbitMQ) ConsumeSimple() {
	// 1.声明队列
	q, err := m.channel.QueueDeclare(m.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否排他
		false,
		// 是否阻塞
		false,
		// 额外参数
		nil)
	if err != nil {
		fmt.Println(err)
	}
	// 2.接受消息
	msgs, err := m.channel.Consume(
		q.Name,
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true, 表示不能将同一Connection中发送的消息，不能发给这个Connection中的消费者
		false,
		// 是否阻塞
		false,
		nil,)
	forever := make(chan bool)
	go func() {
		for d := range msgs{
			log.Printf("Received a message: %s",d.Body)
		}
	}()
	log.Println("waiting for message")
	<- forever
}
