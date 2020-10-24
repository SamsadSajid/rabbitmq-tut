package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
        conn, err := amqp.Dial("amqp://localhost:5672/")
		FailOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
		FailOnError(err, "Failed to open a channel")
        defer ch.Close()

        q, err := ch.QueueDeclare(
                "task_queue", // name
                true,         // durable
                true,        // delete when unused
                false,        // exclusive
                false,        // no-wait
                nil,          // arguments
        )
		FailOnError(err, "Failed to declare a queue")

        body := toJson("01722222222", "120.00", "12.33", "API", "POSTPAID")
		
        err = ch.Publish(
			"",           // exchange
			q.Name,       // routing key
			false,        // mandatory
			false,
			amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "application/json",
					Body:         body,
        })
		FailOnError(err, "Failed to publish a message")
        log.Printf(" [x] Sent %s", body)
}
