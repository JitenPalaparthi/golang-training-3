package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {
	r1 := rect.NewRect(10.12, 123.23)
	s1 := square.NewSquare(100.23)
	shape.PrintShape(r1)
	shape.PrintShape(s1)

	var ishape shape.IShape = r1

	var r2 *rect.Rect = ishape.(*rect.Rect) // type assertion
	fmt.Println(r2)
	ishape = square.NewSquare(123.23)
	switch s := ishape.(type) {
	case *rect.Rect:
		fmt.Println("captured in Rect case", s)
	case *square.Square:
		fmt.Println("captured in *Square case", s)
	case square.Square:
		fmt.Println("captured in Square case", s)
	}

}

// to create a package , it should be inside a directory
