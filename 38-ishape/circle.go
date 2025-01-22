package main

const PI = 3.14

type Circle float32

func NewCircle(r float32) Circle {
	return Circle(r)
}

func (c Circle) Area() float32 {
	return float32(PI * c * c)
}

func (c Circle) Perimeter() float32 {
	return float32(PI * 2 * c)
}

func (c Circle) What() string {
	return "Circle"
}
