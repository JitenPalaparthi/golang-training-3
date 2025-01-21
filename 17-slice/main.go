package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	slice1 := make([]int, 5)
	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}
	num := 100
	fmt.Printf("address of num:%p\n", &num)
	Incr(num) // pass by value
	println(num)
	fmt.Println(slice1)
	//DoubleSlice(slice1)
	//fmt.Println(slice1)
	DoubleSliceAndAdd(slice1, 9999)
	fmt.Println(slice1)

	slice2 := make([]int, 3, 5)
	slice2[0], slice2[1], slice2[2] = 10, 11, 12
	fmt.Println(slice2)
	DoubleSliceAndAdd(slice2, 9999)
	fmt.Println(slice2)
	slice2 = DoubleSliceAndAddR(slice2, 1111)
	fmt.Println(slice2)
}

/*
Slice1 : 0x11ff
Ptr:0x11
Len:5
Cap:5
*/

/*
Slice : 0x11ffbb
Ptr:0x11
Len:5
Cap:5
*/

/*
Slice : 0x11ffbb
Ptr:0x332
Len:6
Cap:10
*/

func DoubleSlice(slice []int) {
	for i, v := range slice {
		slice[i] = v * 2
	}
}

func DoubleSliceAndAdd(slice []int, num int) {
	slice = append(slice, num) // ptr len cap
	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println("Inside the function")
	fmt.Println(slice)
}

func DoubleSliceAndAddR(slice []int, num int) []int {
	slice = append(slice, num) // ptr len cap
	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println("Inside the function")
	fmt.Println(slice)
	return slice
}

func PrintSliceHeader(slice []int, name string) {
	fmt.Println(name)
	fmt.Println(slice)
	fmt.Printf("Address of Slice:%p\n", &slice)
	fmt.Printf("Len:%d\n", len(slice))
	fmt.Printf("Cap:%d\n", cap(slice))
	if len(slice) > 0 {
		fmt.Printf("Ptr:%p\n", &slice[0])
	}
	println("------------")
}

func Incr(n int) {
	fmt.Printf("address of n:%p\n", &n)
	n++
}
