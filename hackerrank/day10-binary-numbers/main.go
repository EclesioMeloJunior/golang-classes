package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToBinary(n string) string {
	number, _ := strconv.ParseInt(n, 10, 64)
	return strconv.FormatInt(number, 2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.Replace(line, "\n", "", -1)

	binaryNumber := convertToBinary(line)

	maxsequency := 0
	sequency := 0

	for _, value := range binaryNumber {
		number := string(value)

		if number == "1" {
			sequency++

			if sequency > maxsequency {
				maxsequency = sequency
			}

			continue
		}

		sequency = 0
	}

	fmt.Print(maxsequency)
}
