package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	auxiliar := *a
	*a = *b
	*b = auxiliar
}

func partition(arr []int, low int, high int) int {
	var pivotElement int = arr[high]
	var smallElementIndex int = low - 1

	for controll := low; controll <= high-1; controll++ {
		if arr[controll] < pivotElement {
			smallElementIndex++
			swap(&arr[smallElementIndex], &arr[controll])
		}
	}

	swap(&arr[smallElementIndex+1], &arr[high])
	return (smallElementIndex + 1)
}

func quicksort(arr []int, start int, end int) {
	if start < end {
		partition := partition(arr, start, end)

		quicksort(arr, start, partition-1)
		quicksort(arr, partition+1, end)
	}
}

func main() {
	nums := []int{165, 4, 78, 90, 2, 5, 7, 1, 0, 87, 24, 32, 10}
	listSize := len(nums)

	quicksort(nums, 0, listSize-1)

	fmt.Println(nums)
}
