package shape

import "fmt"

func init() {
	println("shape packaged is called-1")
}

type IShape interface {
	Area() float32
	Perimeter() float32
	IWhat
}

type IWhat interface {
	What() string
}

func PrintShape(ishape IShape) {
	fmt.Println("Area of ", ishape.What(), ":", ishape.Area())
	fmt.Println("Perimeter of ", ishape.What(), ":", ishape.Perimeter())
	println()
}
