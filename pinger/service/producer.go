package service

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
	"log"
	"os"
	"pinger/models"
	"time"
)

func SendPingResult(container *models.Container) error {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	var (
		conn *amqp.Connection
		err  error
	)
	for {
		conn, err = amqp.Dial(rabbitMQURL)
		if err == nil {
			log.Println("Successful connection to RabbitMQ server")
			break
		}
		time.Sleep(5 * time.Second)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("unable to open a channel. Error: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"containers-info",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	body, err := json.Marshal(container)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("failed to publish a container: %s. Error: %s", container.IP, err)
	}

	log.Printf("container with IP=%s have been published", container.IP)

	return err
}
