package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func factorial(n int) int {
	if n <= 1 {
		return n
	}

	return n * factorial(n-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.Replace(line, "\n", "", -1)
	number, _ := strconv.Atoi(line)

	fact := factorial(number)

	fmt.Print(fact)
}
