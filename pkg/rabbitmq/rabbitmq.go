package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	cnfg "github.com/Shakezidin/pkg/config"
	msgg "github.com/Shakezidin/pkg/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeConfirmationMessages(cnfg *cnfg.Conf) {
	// Connect to RabbitMQ
	// val := fmt.Sprint("amqp://guest:guest@localhost:",cnfg.PORT )
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	// Declare the queue
	q, err := ch.QueueDeclare(
		"confirmation_queue", // queue name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process incoming messages
	for msg := range msgs {
		// Extract booking details from the message
		var bookingDetails msgg.Messages
		err := json.Unmarshal(msg.Body, &bookingDetails)
		if err != nil {
			log.Printf("Error decoding message body: %v", err)
			continue
		}
		fmt.Println(bookingDetails)

		// Send confirmation email using booking details
		err = msgg.SendConfirmationEmail(cnfg, bookingDetails)
		if err != nil {
			log.Printf("Error sending confirmation email: %v", err)
			// Handle error (e.g., retry logic)
		}
	}
}
