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

func strToNumber(nmrs []string) []int {
	numbers := make([]int, len(nmrs))
	for i, n := range nmrs {
		numbers[i], _ = strconv.Atoi(n)
	}

	return numbers
}

func swap(a, b *int) {
	tmp := *a

	a = b
	b = &tmp
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < t; i++ {
		nmrs := strToNumber(strings.Split(readLine(reader), " "))

		aws := 1e18

		for j := 0; j < 2; j++ {
			da := (int)(math.Min((float64)(nmrs[4]), (float64)(nmrs[0]-nmrs[2])))
			db := (int)(math.Min((float64)(nmrs[4]-da), (float64)(nmrs[1]-nmrs[3])))

			aws = math.Min(aws, (float64)((nmrs[0]-da)*(nmrs[1]-db)))

			swap(&nmrs[1], &nmrs[0])
			swap(&nmrs[2], &nmrs[3])
		}

		fmt.Println((int)(aws))
	}

}
