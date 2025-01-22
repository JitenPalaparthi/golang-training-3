package main

import "fmt"

func main() {
	var ishape IShape

	ishape = NewRect(100.4, 201)
	a1 := ishape.Area()
	p1 := ishape.Perimeter()

	ishape = NewSquare(25.3)
	a2 := ishape.Area()
	p2 := ishape.Perimeter()

	fmt.Println("Area of Rect:", a1)
	fmt.Println("Perimeter of Rect:", p1)

	fmt.Println("Area of Square:", a2)
	fmt.Println("Perimeter of Square:", p2)

	ishape = &Rect{L: 101.4, B: 124.5}
	a1 = ishape.Area()
	p1 = ishape.Perimeter()

	fmt.Println("Area of Rect:", a1)
	fmt.Println("Perimeter of Rect:", p1)

}

type IShape interface {
	Area() float32
	Perimeter() float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) Display() {
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.B)
}

type Square float32

func NewSquare(s float32) Square {
	return Square(s)
}

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}
