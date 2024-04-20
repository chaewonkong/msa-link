package queue

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ represents the message queue
type RabbitMQ struct {
	// Conn is the connection
	Conn *amqp.Connection

	// Ch is the channel
	Ch *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ
func NewRabbitMQ(connURL string) *RabbitMQ {
	conn, err := amqp.Dial(connURL)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &RabbitMQ{conn, ch}
}

// Close closes the connection and channel
func (mq *RabbitMQ) Close() {
	mq.Conn.Close()
	mq.Ch.Close()
}

// URL returns a formatted string for a RabbitMQ connection.
func URL(user, password, host, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
}
