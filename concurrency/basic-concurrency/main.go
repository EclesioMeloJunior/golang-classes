package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("Hello")
		wg.Done()
	}()

	wg.Wait()
}
