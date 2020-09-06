package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')

	return strings.Replace(line, "\r\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < n; i++ {
		nmrs := strings.Split(readLine(reader), " ")

		n1, _ := strconv.Atoi(nmrs[0])
		n2, _ := strconv.Atoi(nmrs[1])

		n3 := (int)(math.Abs((float64)(n1 - n2)))

		result := (n3 + 9) / 10

		fmt.Println(result)
	}
}
