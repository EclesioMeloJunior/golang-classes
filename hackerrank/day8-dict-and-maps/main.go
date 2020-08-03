package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func entriesToNumber(firstline string) int {
	toString := strings.Replace(firstline, "\n", "", -1)
	toNumber, _ := strconv.Atoi(toString)

	return toNumber
}

func onlyWord(receivedFromTerminal string) string {
	return strings.Replace(receivedFromTerminal, "\n", "", -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')

	entries := entriesToNumber(firstLine)

	phoneBook := make(map[string]string)

	for i := 0; i < entries; i++ {
		line, _ := reader.ReadString('\n')

		infos := strings.Split(line, " ")

		phoneBook[infos[0]] = infos[1]
	}

	for {
		search, _ := reader.ReadString('\n')
		searchString := onlyWord(search)

		if searchString == "" {
			break
		}

		phone, exists := phoneBook[searchString]

		if !exists {
			fmt.Print("Not found")
			continue
		}

		message := fmt.Sprintf("%s=%s", searchString, phone)
		fmt.Print(message)
	}
}
