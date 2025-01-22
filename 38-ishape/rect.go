package main

import "fmt"

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

func (r *Rect) What() string {
	return "Rect"
}

func (r *Rect) Display() {
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.B)
}
