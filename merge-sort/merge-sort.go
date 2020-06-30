package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func split(arr []int) []int {
	lenarr := len(arr)

	if lenarr == 1 {
		return arr
	}

	med := lenarr / 2

	leftarr := split(arr[0:med])
	rightarr := split(arr[med:lenarr])

	return merge(leftarr, rightarr)
}

func merge(left []int, right []int) []int {
	lenleft := len(left)
	lenright := len(right)

	mergedarr := make([]int, lenleft+lenright)

	pontleft := 0
	pontright := 0

	for index := range mergedarr {

		if pontleft == lenleft {
			rightValue := right[pontright]
			mergedarr[index] = rightValue
			pontright++

			continue
		}

		if pontright == lenright {
			leftValue := left[pontleft]
			mergedarr[index] = leftValue
			pontleft++

			continue
		}

		rightValue := right[pontright]
		leftValue := left[pontleft]

		if leftValue > rightValue {
			mergedarr[index] = rightValue
			pontright++
		} else {
			mergedarr[index] = leftValue
			pontleft++
		}
	}

	return mergedarr
}

func main() {
	file, err := os.Open("numbers.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var numbers []int

	byteValue, _ := ioutil.ReadAll(file)
	fmt.Println(len(byteValue))

	json.Unmarshal(byteValue, &numbers)

	fmt.Println(fmt.Sprintf("Ordering: %v", len(numbers)))
	fmt.Print("===== Merge Sort Algorithm\n")
	split(numbers)

	//fmt.Print("===== Order\n")
	//fmt.Println(ordered)
}
