package rabbitmq

import (
	"dropx/pkg/config"
	"dropx/pkg/logger"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
	mu         sync.Mutex
	closeChan  chan *amqp.Error
}

func InitRabbitMQ() *RabbitMQ {
	r := &RabbitMQ{}
	if err := r.connect(); err != nil {
		logger.FatalLog(fmt.Errorf("failed to initialize rabbitmq: %v", err), nil)
	}
	return r
}

func (r *RabbitMQ) connect() error {
	var err error
	r.Connection, err = amqp.Dial(config.Global.RabbitMqUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	r.Channel, err = r.Connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	logger.InfoLog("RabbitMQ channel opened successfully", nil)

	// Listen close signal
	r.closeChan = make(chan *amqp.Error)
	r.Connection.NotifyClose(r.closeChan)

	// Handle reconnect
	go r.handleReconnect()

	return nil
}

func (r *RabbitMQ) handleReconnect() {
	for err := range r.closeChan {
		if err != nil {
			logger.ErrorLog(fmt.Errorf("rabbitmq connection closed, reason: %v", err), nil)
		}

		for {
			time.Sleep(2 * time.Second)
			logger.InfoLog("Attempting to reconnect to RabbitMQ...", nil)

			r.mu.Lock()
			if err := r.connect(); err != nil {
				logger.ErrorLog(fmt.Errorf("reconnect to rabbitmq failed: %v", err), nil)
				r.mu.Unlock()
				continue
			}
			r.mu.Unlock()

			logger.InfoLog("RabbitMQ reconnected successfully", nil)
			break
		}
	}
}

func (r *RabbitMQ) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			logger.ErrorLog(fmt.Errorf("failed to close channel: %v", err), nil)
		}
	}

	if r.Connection != nil {
		if err := r.Connection.Close(); err != nil {
			logger.ErrorLog(fmt.Errorf("failed to close connection: %v", err), nil)
		}
	}
}

func (r *RabbitMQ) Publish(queue string, data interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal publish data: %w", err)
	}

	_, err = r.Channel.QueueDeclare(
		queue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = r.Channel.Publish(
		"",    // default exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish: %w", err)
	}

	return nil
}

func (r *RabbitMQ) Consume(queue string, handler func([]byte)) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.Channel.QueueDeclare(
		queue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		logger.FatalLog(fmt.Errorf("failed to declare queue: %v", err), nil)
	}

	msgs, err := r.Channel.Consume(
		queue,
		"",    // consumer name
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		logger.FatalLog(fmt.Errorf("failed to consume queue: %v", err), nil)
	}

	go func() {
		for msg := range msgs {
			handler(msg.Body)
		}
	}()
}
