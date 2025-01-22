package main

import "fmt"

type Fn func(int, int) int

func (f Fn) Sq(n int) int {
	return n * n
}

func main() {

	func() {
		fmt.Println("hello ICICI")
	}() // this executes

	func(msg string) {
		fmt.Println(msg)
	}("Hello ICICI Direct") // this executes

	r := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(r)

	f := func(a, b int) int {
		return a - b
	}

	r2 := f(10, 20)

	fmt.Println(r2)

	a1, b1 := 12, 30
	r3 := Calc(a1, b1, func(i1, i2 int) int {
		return i1 + i2
	})

	fmt.Println("Sum:", r3)

	r4 := Calc(10, 2, Mul)
	fmt.Println("Mu.:", r4)

	fn1 := func(a, b int) int {
		return a - b
	}
	r5 := Calc2(a1, b1, Fn(fn1))
	fmt.Println("Sub:", r5)
}

func Calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func Calc2(a, b int, fn Fn) int {
	return fn(a, b)
}

func Mul(a, b int) int {
	return a * b
}
