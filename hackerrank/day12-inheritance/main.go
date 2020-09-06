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

type Person struct {
	id       string
	irstName string
	lastName string
}

func NewPerson(id, firstName, lastName string) *Person {
	return &Person{
		id,
		firstName,
		lastName,
	}
}

type Student struct {
	*Person
	scores []int
}

func NewStudent(p *Person, scores []int) *Student {
	return &Student{
		p,
		scores,
	}
}

func (s *Student) calcule() string {
	if len(s.scores) < 1 {
		return "T"
	}

	acc := 0

	for _, score := range s.scores {
		acc += score
	}

	acc = acc / len(s.scores)

	if acc >= 90 && acc <= 100 {
		return "O"
	}

	if acc >= 80 && acc < 90 {
		return "E"
	}

	if acc >= 70 && acc < 80 {
		return "A"
	}

	if acc >= 55 && acc < 70 {
		return "P"
	}

	if acc >= 40 && acc < 55 {
		return "D"
	}

	return "T"
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	personLine := readLine(reader)
	personData := strings.Split(personLine, " ")

	person := NewPerson(personData[2], personData[0], personData[1])

	readLine(reader)

	scoresLine := readLine(reader)
	scoresString := strings.Split(scoresLine, " ")

	scores := make([]int, 0)

	for _, strScore := range scoresString {
		score, _ := strconv.Atoi(strScore)
		scores = append(scores, score)
	}

	student := NewStudent(person, scores)

	fmt.Println(student.calcule())
}
