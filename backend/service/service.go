package service

import (
	"backend/models"
	"backend/repository"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	_ "log"
	"os"
	"time"
)

type Service struct {
	Repo *repository.Repository
}

func (s *Service) GetContainers() ([]models.Container, error) {
	return s.Repo.GetContainers()
}

func (s *Service) StartConsume() {
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

	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	go func() {
		for message := range messages {
			messageBody := message.Body
			log.Printf("received a message: %s", messageBody)
			err := s.Repo.AddContainer(messageBody)
			if err != nil {
				log.Fatalf("recieved error during consuming messages: %s\n", err)
			}
		}
	}()
	select {}
}
