package main

import (
	"fmt"
)

type Polygon interface {
	getArea() float64
	getShape() string
}

type Triangle struct {
	shape  string
	base   float64
	height float64
}

type Square struct {
	shape string
	side  float64
}

func main() {
	t := Triangle{base: 3, height: 4}
	s := Square{side: 5}

	printArea(t)
	printArea(s)
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (t Triangle) getShape() string {
	return "Triangle"
}

func (s Square) getShape() string {
	return "Square"
}

func printArea(p Polygon) {
	fmt.Println("Area of", p.getShape(), "is", p.getArea())
}
