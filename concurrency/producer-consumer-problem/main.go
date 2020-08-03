package main

import "fmt"

type Consumer struct {
	messages *chan int
}

func NewConsumer(messages *chan int) *Consumer {
	return &Consumer{
		messages: messages,
	}
}

func (c *Consumer) consume() {
	fmt.Println("Consumer: Started")

	for {
		message := <-*c.messages
		fmt.Println("Consumer: Received", message)
	}
}

type Producer struct {
	messages *chan int
	done     *chan bool
}

func NewProducer(messages *chan int, done *chan bool) *Producer {
	return &Producer{
		messages: messages,
		done:     done,
	}
}

func (p *Producer) produce(max int) {
	fmt.Println("Producer: Started")

	for i := 0; i < max; i++ {
		fmt.Println("Producer: Sending", i)
		*p.messages <- i
	}

	*p.done <- true

	fmt.Println("Producer: Done")
}

func main() {

	messages := make(chan int, 1)
	done := make(chan bool)

	go NewProducer(&messages, &done).produce(5)
	go NewConsumer(&messages).consume()

	<-done
}
