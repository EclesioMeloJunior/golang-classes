package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')

	return strings.Replace(line, "\n", "", -1)
}

func main() {
	arr := make([][]int, 6)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 6; i++ {
		line := readLine(reader)

		numbersInLine := strings.Split(line, " ")

		column := make([]int, 6)

		for j := 0; j < 6; j++ {
			number, _ := strconv.Atoi(numbersInLine[j])
			column[j] = number
		}

		arr[i] = column
	}

	var max int
	var lineindex int

	for lineindex = 0; lineindex < 4; lineindex++ {
		head := arr[lineindex]
		middle := arr[lineindex+1]
		bottom := arr[lineindex+2]

		for colindex := 0; colindex < 4; colindex++ {
			headValue := head[colindex] + head[colindex+1] + head[colindex+2]
			midValue := middle[colindex+1]
			bottom := bottom[colindex] + bottom[colindex+1] + bottom[colindex+2]

			sum := headValue + midValue + bottom

			if lineindex == 0 && colindex == 0 {
				max = sum
			} else if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
