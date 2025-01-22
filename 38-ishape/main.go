package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var ishapes []IShape
	ishapes = append(ishapes, NewRect(10, 13), NewSquare(123.2), NewRect(123.123, 123.2), NewSquare(123.23), NewRect(123.23, 43.3))
	c1 := NewCuboid(12.3, 23.3, 43.5)
	ishapes = append(ishapes, c1)
	ishapes = append(ishapes, NewCircle(123.23))

	for _, v := range ishapes {
		Shape(v)
	}

	for i := 1; i <= 10; i++ {
		r := rand.IntN(8)
		Shape(ishapes[r])
	}

	var ishape1 IShape = NewRect(10, 20)
	ishape1.Area()
	var r Rect
	r.Area()

}

func Shape(ishape IShape) {
	fmt.Println("Area of ", ishape.What(), ":", ishape.Area())
	fmt.Println("Perimeter of ", ishape.What(), ":", ishape.Perimeter())
	println()
}
