package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type Logger interface {
	Error(args ...interface{})

	PanicOnErr(err error, args ...interface{})
	FatalOnErr(err error, args ...interface{})
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	log  Logger

	register []RegisterDto
}

type RegisterDto struct {
	Exchange   string
	QueueName  string
	RoutingKey string
	Resolver   Resolver
}

type Resolver interface {
	Handle(result []byte) error
}

type Credentials struct {
	Host     string
	Port     int
	User     string
	Password string
}

func newRabbitMQ(conn *amqp.Connection, ch *amqp.Channel, log Logger) *RabbitMQ {
	return &RabbitMQ{conn: conn, ch: ch, log: log}
}

func (r *RabbitMQ) exchangeDeclare(exchange string) {
	err := r.ch.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
	)
	r.log.PanicOnErr(err, "Failed to declare a exchange")
}

func (r *RabbitMQ) queueDeclare(exchange string, queueName string, routingKey string) {
	_, err := r.ch.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	r.log.PanicOnErr(err, fmt.Sprintf("Failed to declare a consumers %s", queueName))

	err = r.ch.QueueBind(queueName, routingKey, exchange, false, nil)

	r.log.PanicOnErr(
		err,
		fmt.Sprintf("Failed to bind a consumers %s %s with routing key: %s", queueName, exchange, routingKey),
	)
}

func (r *RabbitMQ) Register(input RegisterDto) {
	r.exchangeDeclare(input.Exchange)
	r.queueDeclare(input.Exchange, input.QueueName, input.RoutingKey)

	r.register = append(r.register, input)
}

func (r *RabbitMQ) Listen(ctx context.Context, wg *sync.WaitGroup) {

	grouped := make(map[string]map[string]RegisterDto)
	for _, input := range r.register {
		if _, exists := grouped[input.QueueName]; !exists {
			grouped[input.QueueName] = make(map[string]RegisterDto)
		}

		grouped[input.QueueName][input.RoutingKey] = input
	}

	for queueName, items := range grouped {
		wg.Add(1)
		go func(queueName string, items map[string]RegisterDto) {
			defer wg.Done()

			defer func() {
				if r := recover(); r != nil {
				}
			}()

			msgs, err := r.ch.Consume(
				queueName, "", true, false, false, false, nil,
			)
			r.log.PanicOnErr(err, fmt.Sprintf("Failed to register a consumers %s", queueName))

			for {
				select {
				case msg := <-msgs:
					if result, exists := items[msg.RoutingKey]; exists {
						fmt.Println(queueName, msg.RoutingKey, msg.Exchange, string(msg.Body))
						if err := result.Resolver.Handle(msg.Body); err != nil {
							r.log.Error(err)
						}
					}
					continue
				case <-ctx.Done():
					fmt.Println("Consumer stopped")
					return
				}
			}
		}(queueName, items)
	}
}

func (r *RabbitMQ) Close() {
	r.log.PanicOnErr(r.conn.Close(), "Failed to close connection")
	r.log.PanicOnErr(r.ch.Close(), "Failed to close channel")
	fmt.Println("RabbitMQ closed")
}

func InitRabbitMQ(c Credentials, log Logger) *RabbitMQ {
	url := fmt.Sprintf("amqp://%s:%s@%s:%v/", c.User, c.Password, c.Host, c.Port)
	conn, err := amqp.Dial(url)
	log.FatalOnErr(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	log.FatalOnErr(err, "Failed to open a channel")

	return newRabbitMQ(conn, ch, log)
}
