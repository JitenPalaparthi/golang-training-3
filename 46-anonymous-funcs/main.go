package main

import (
	"fmt"
)

type Fn func(int, int) int

func (f Fn) Caller(s string) {
	fmt.Println("calling a caller", s)
}

func main() {

	map1 := make(map[string]any)

	map1["add"] = func(a, b int) int {
		return a + b
	}
	map1["sub"] = func(a, b int) int {
		return a - b
	}
	map1["mul"] = Mul

	var fn1 Fn = func(i1, i2 int) int {
		return i1 / i2
	}

	map1["div"] = fn1

	map1["greet"] = func() {
		fmt.Println("hello world")
	}

	a1, b1 := 10, 20

	for k, v := range map1 {
		switch v.(type) {
		case func():
			fmt.Println(k, "func()")
			(v.(func()))()
		case Fn:
			r := (v.(Fn))(a1, b1)
			//r1 := ((func(int, int) int)(v.(Fn)))(a1, b1)
			fmt.Println(k, "Fn -->", r)
		case func(int, int) int:
			r := (v.(func(int, int) int))(a1, b1)
			fmt.Println(k, "func(int,int)int -->", r)
			Fn(v.(func(int, int) int)).Caller(k)
		}
		//fmt.Println("Key:", k, "V of type:", reflect.TypeOf(v))
	}

	fn2 := Calc(100, 20)
	r := fn2()
	fmt.Println(r)

}

func Mul(a, b int) int {
	return a * b
}

func Calc(a, b int) func() int {
	return func() int {
		return a + b
	}
}
