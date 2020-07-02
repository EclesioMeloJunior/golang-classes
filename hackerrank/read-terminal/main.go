package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func onlyWord(receivedFromTerminal string) string {
	return strings.Replace(receivedFromTerminal, "\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	total, _ := reader.ReadString('\n')
	total = onlyWord(total)

	fmt.Println(total)
}
