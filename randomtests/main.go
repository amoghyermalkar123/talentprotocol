package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	openai "github.com/sashabaranov/go-openai"
)

func umain() {
	client := openai.NewClient("")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()
	// Create a topic name
	topic := "evaluation-queue"

	// Publish a message to the topic
	message := []byte("Hello, NATS!")
	err = nc.Publish(topic, message)
	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	fmt.Printf("Published message to topic '%s': %s\n", topic, message)
	for {
	}
}

func sub() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	topic := "evaluation-queue"

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(topic, ch)
	// handle err
	for msg := range ch {
		// do something to the nats.Msg object
		fmt.Println("recv", msg)
	}
	// Unsubscribe if needed
	sub.Unsubscribe()
}
