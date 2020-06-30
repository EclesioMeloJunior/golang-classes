package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadNumber() (numbers []int) {
	defer func() {
		fmt.Printf("Total Numbers: %v\n", len(numbers))
	}()

	file, err := os.Open("numbers.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	fmt.Println(len(byteValue))

	json.Unmarshal(byteValue, &numbers)

	return numbers
}

func swap(a *int, b *int) {
	auxiliar := *a
	*a = *b
	*b = auxiliar
}

func bubblesort(numbers []int) []int {
	var i, j int

	for i = 0; i < len(numbers)-1; i++ {
		for j = 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(&numbers[j], &numbers[j+1])
			}
		}
	}

	return numbers
}

func main() {
	n := loadNumber()
	order := bubblesort(n)

	fmt.Println(order[0])
	fmt.Println(order[len(order)-1])
	fmt.Println(order)
}
