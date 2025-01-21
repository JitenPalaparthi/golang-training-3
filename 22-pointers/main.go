package main

import "fmt"

func main() {
	n := 100
	Incr(n)
	println(n)

	IncrP(&n)
	println(n)

	ptr := Sq3(100)
	println(*ptr)

	ptr2 := new(int)
	Sq4(100, ptr2)
	println(*ptr2)

	println("End of main")

	fmt.Println(*ptr, *ptr2)

	slice1 := []int{10, 11, 12, 13, 143, 15, 18} // stack
	slice2 := make([]int, 10000)                 // heap memory
	slice1 = append(slice1, slice2...)

	var arr1 [10000]int
	println(&slice1)
	println(&slice2)
	println(&arr1)

}

func Incr(n int) {
	n++
}

func IncrP(p *int) {
	if p != nil {
		*p++
	}
}

func Sq1(n int) int {
	return n * n
}

func Sq2(p *int) int {
	if p == nil {
		return 0
	}
	return *p * *p
}

func Sq3(n int) *int {
	p := new(int) // creating a new pointer variable
	*p = n * n    // deferencing the varialbe
	return p      // returning it
}

func Sq4(n int, p *int) {
	// creating a new pointer variable
	if p == nil {
		return
	}
	*p = n * n // deferencing the varialbe
}

// Heap leaks // no chance of heap leak
// dangling pointer // this avoided by complier
// dereference the nil pointer
// double free // avoided in Go

// ICICI Java Based Server

// Tomcap
