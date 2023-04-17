package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

// defind a behavior
type Stringer interface {
	String() string
}

// defind a method with that method behavior, meaning you can
// reuse this method via the Stringer interface without rewriting
// your function
func (a *Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.\n", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Page   int
}

func (b *Book) String() string {
	return fmt.Sprintf("the %q book was written by %s.\n", b.Title, b.Author)
}

func main() {
	a := &Article{
		Title:  "Understanding Interface in Go",
		Author: "Sammy Shark",
	}
	Print(a)

	b := &Book{
		Title:  "The godfather",
		Author: " Mario Puzo",
		Page:   50,
	}
	Print(b)

}

// defind a function accept the interface, then pass the interface method
func Print(s Stringer) {
	fmt.Println(s.String())
}
