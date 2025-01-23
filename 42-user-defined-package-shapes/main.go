package main

import (
	. "fmt"
	a "shapes/shape"
	rectangle "shapes/shape/rect"
	c "shapes/shape/square"
	. "time"
)

func init() {
	println("main package is called-1")
}

func main() {
	r1 := rectangle.NewRect(10.12, 123.23)
	s1 := c.NewSquare(100.23)
	a.PrintShape(r1)
	a.PrintShape(s1)

	var ishape a.IShape = r1

	var r2 *rectangle.Rect = ishape.(*rectangle.Rect) // type assertion
	Println(r2)
	ishape = c.NewSquare(123.23)
	switch s := ishape.(type) {
	case *rectangle.Rect:
		Println("captured in Rect case", s)
	case *c.Square:
		Println("captured in *Square case", s)
	case c.Square:
		Println("captured in Square case", s)
	}
	Println(Now())
}

// to create a package , it should be inside a directory
