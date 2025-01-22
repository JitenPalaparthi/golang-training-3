package main

import "fmt"

func main() {
	r1 := Rect{L: 12.2, B: 14.5}
	a1 := Area(r1)
	p1 := Perimeter(r1)
	// What are Area and Perimeter?
	fmt.Println("area:", a1)
	fmt.Println("perimeter:", p1)

	r2 := Rect{L: 12.2, B: 14.5}
	fmt.Printf("Address of r2:%p\n", &r2)

	a2 := (&r2).Area() // call by ref
	p2 := r2.Perimeter()
	fmt.Println("area:", a2)
	fmt.Println("perimeter:", p2)
	fmt.Println("Area in r2:", r2.A)      // ???
	fmt.Println("Perimeter in r2:", r2.P) // ???

	r3 := &Rect{L: 12.2, B: 34.3}
	a3 := r3.Area() // call by ref
	p3 := r3.Perimeter()
	fmt.Println("area:", a3)
	fmt.Println("perimeter:", p3)
	fmt.Println("Area in r3:", r3.A)      // ???
	fmt.Println("Perimeter in r3:", r3.P) // ???

	//r4 := new(Rect)

	//var ptr *int

	var r4 *Rect
	if r4 == nil {
		println("nil")
	}
	r5 := new(Rect)
	fmt.Println(r5)
	(*r5).L = 123.23
	r5.B = 12.4
	a5 := r5.Area()
	fmt.Println("Area of r5:", a5)
	p5 := r5.Perimeter()
	fmt.Println("Perimeter of r5:", p5)

	r6 := New(12, 45)
	a6, p6 := r6.Area(), r6.Perimeter()
	fmt.Println("area and perimeter", a6, p6)

}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func NewDefault() *Rect {
	return &Rect{L: 1, B: 1}
}

// func NewSquare(s float32) *Square {
// 	return &Square{S: s}
// }

// type Square struct {
// 	S float32
// }

type Rect struct {
	//L,B,A,P float32
	L, B float32
	A    float32
	P    float32
	//Ok   bool
}

func Area(r Rect) float32 {
	r.A = r.L * r.B
	return r.A
}

func Perimeter(r Rect) float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) Area() float32 {
	fmt.Printf("Address of r:%p\n", &r)
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}
