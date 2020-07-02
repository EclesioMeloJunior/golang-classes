package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func separeWord(word string) {
	oddchars := []string{}
	evenchars := []string{}

	for letterindex, letter := range word {
		if letterindex%2 == 0 {
			evenchars = append(evenchars, string(letter))
			continue
		}

		oddchars = append(oddchars, string(letter))
	}

	leftside := strings.Join(evenchars, "")
	rightside := strings.Join(oddchars, "")

	fmt.Printf("%s\n", fmt.Sprintf("%s %s", leftside, rightside))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	total, _ := reader.ReadString('\n')
	total = strings.Replace(total, "\n", "", -1)

	totalwords, _ := strconv.Atoi(total)

	for i := 0; i < totalwords; i++ {
		word, _ := reader.ReadString('\n')
		word = strings.Replace(word, "\n", "", -1)

		separeWord(word)
	}
}
