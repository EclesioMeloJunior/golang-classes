package main

import (
	"fmt"
	"time"
)

func say(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(word)
	}
}

func main() {
	go say("FODASE")
	say("ABS")
}
