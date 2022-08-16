package go_rabbitmq

import (
	"github.com/wagslane/go-rabbitmq"
	"log"
)

func consumerDefault() {
	consumer, err := rabbitmq.NewConsumer(
		"amqp://guest:guest@localhost", rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		"my_queue",
		[]string{"routing_key1", "routing_key2"},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func producer() {
	publisher, err := rabbitmq.NewPublisher("amqp://user:pass@localhost", rabbitmq.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()
	err = publisher.Publish([]byte("hello, world"), []string{"routing_key"})
	if err != nil {
		log.Fatal(err)
	}
}
