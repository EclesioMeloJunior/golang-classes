package main

import (
	"fmt"
	"strings"
)

func main() {
	arrlen := 6
	myarr := "1 2 3 4 5 100"

	removespaces := strings.Split(myarr, " ")

	for i := arrlen - 1; i >= 1; i-- {
		fmt.Printf("%s ", string(removespaces[i]))
	}

	fmt.Printf("%s", string(removespaces[0]))
}
