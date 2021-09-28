package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	// look at queue messages
	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)

	forever := make(chan bool)
	// anonymous function
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ instance")
	fmt.Println(" [+] - waiting for messages")
	// receive from the channel, blocks main.go from exiting
	<-forever


}