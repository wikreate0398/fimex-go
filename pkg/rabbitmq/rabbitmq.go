package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"wikreate/fimex/pkg/failed"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

type ListnerInput struct {
	Ctx        context.Context
	Exchange   string
	QueueName  string
	RoutingKey string
	Resolver   Resolver
	Wg         *sync.WaitGroup
}

type Resolver interface {
	Handle()
	ToStruct(result []byte)
}

type Credentials struct {
	Host     string
	Port     int
	User     string
	Password string
}

func NewRabbitMQ(conn *amqp.Connection, ch *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{conn, ch}
}

func (r *RabbitMQ) exchangeDeclare(exchange string) {
	err := r.ch.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
	)
	failed.PanicOnError(err, "Failed to declare a exchange")
}

func (r *RabbitMQ) queueDeclare(exchange string, queueName string, routingKey string) amqp.Queue {
	q, err := r.ch.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	failed.PanicOnError(err, fmt.Sprintf("Failed to declare a consumers %s", queueName))

	err = r.ch.QueueBind(queueName, routingKey, exchange, false, nil)
	failed.PanicOnError(
		err,
		fmt.Sprintf("Failed to bind a consumers %s %s with routing key: %s", queueName, exchange, routingKey),
	)

	return q
}

func (r *RabbitMQ) Listen(input ListnerInput) {
	input.Wg.Add(1)
	go func() {
		defer input.Wg.Done()
		r.exchangeDeclare(input.Exchange)
		q := r.queueDeclare(input.Exchange, input.QueueName, input.RoutingKey)

		msgs, err := r.ch.Consume(
			q.Name, "", true, false, false, false, nil,
		)
		failed.PanicOnError(err, fmt.Sprintf("Failed to register a consumers %s", q.Name))

		for {
			select {
			case msg := <-msgs:
				input.Resolver.ToStruct(msg.Body)
				input.Resolver.Handle()
				continue
			case <-input.Ctx.Done():
				fmt.Println("Consumer stopped")
				return
			}
		}
	}()
}

func (r *RabbitMQ) Close() {
	r.conn.Close()
	r.ch.Close()
	fmt.Println("RabbitMQ closed")
}

func InitRabbitMQ(c Credentials) *RabbitMQ {
	url := fmt.Sprintf("amqp://%s:%s@%s:%v/", c.User, c.Password, c.Host, c.Port)
	conn, err := amqp.Dial(url)
	failed.PanicOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failed.PanicOnError(err, "Failed to open a channel")

	return NewRabbitMQ(conn, ch)
}
