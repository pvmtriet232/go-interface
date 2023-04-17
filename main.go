package main

import "fmt"

func main() {
	An := &human{"Hello", 70}
	Kiki := &animal{"gau gau", 4}
	SaySomeThing(An)
	SaySomeThing(Kiki)
}
func SaySomeThing(l listener) {
	l.listen()
}

type listener interface {
	listen()
}

type animal struct {
	sound  string
	maxAge int
}

func (s *animal) listen() {
	fmt.Println(s.sound)
}

type human struct {
	sound  string
	maxAge int
}

func (s *human) listen() {
	fmt.Println(s.sound)
}
