package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle {Radius : %.2f}", c.radius)
}

type Square struct {
	Width  float64
	Height float64
}

func (s *Square) Area() float64 {
	return s.Height * s.Width
}

func (s *Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func main() {

	c := &Circle{5}
	s := &Square{4, 3}
	a := Less(c, s)
	fmt.Printf("less is %+v\n", a)
	PrintArea(s)
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() > s2.Area() {
		return s2
	}
	return s1
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
