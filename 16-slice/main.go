package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []int // declartion of slice
	//slice2 := []int{10, 12, 13, 14, 15}
	if slice1 == nil {
		println("nil slice")
	}
	slice1 = make([]int, 5) //
	// fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}
	fmt.Println(slice1)

	fmt.Println("Slice1")
	fmt.Println(slice1)
	fmt.Printf("Address of Slice:%p\n", &slice1)
	fmt.Printf("Len:%d\n", len(slice1))
	fmt.Printf("Cap:%d\n", cap(slice1))
	if len(slice1) > 0 {
		fmt.Printf("Ptr:%p\n", &slice1[0])
	}
	println("------------")

	slice1 = append(slice1, 9999) // add new element to the slice
	fmt.Println("Slice1")
	fmt.Println(slice1)
	fmt.Printf("Address of Slice:%p\n", &slice1)
	fmt.Printf("Len:%d\n", len(slice1))
	fmt.Printf("Cap:%d\n", cap(slice1))
	if len(slice1) > 0 {
		fmt.Printf("Ptr:%p\n", &slice1[0])
	}
	println("------------")

	slice1 = append(slice1, 1111)
	fmt.Println("Slice1")
	fmt.Println(slice1)
	fmt.Printf("Address of Slice:%p\n", &slice1)
	fmt.Printf("Len:%d\n", len(slice1))
	fmt.Printf("Cap:%d\n", cap(slice1))
	if len(slice1) > 0 {
		fmt.Printf("Ptr:%p\n", &slice1[0])
	}
	println("------------")

	slice1 = append(slice1, 2222, 3333, 4444, 5555)
	fmt.Println("Slice1")
	fmt.Println(slice1)
	fmt.Printf("Address of Slice:%p\n", &slice1)
	fmt.Printf("Len:%d\n", len(slice1))
	fmt.Printf("Cap:%d\n", cap(slice1))
	if len(slice1) > 0 {
		fmt.Printf("Ptr:%p\n", &slice1[0])
	}
	println("------------")

	slice2 := []int{109, 300, 450, 600, 24, 24, 56, 38, 89, 200} // shorthand declaraion
	slice2 = append(slice2, 10, 20, 30)
	s1 := SumOf(slice1)
	s2 := SumOf(slice2)

	arr1 := [5]int{10, 11, 12, 13, 14}
	s3 := SumOf(arr1[:])

	println(s1, s2, s3)
}

/*
Slice Header &0x11110011
-------------
Ptr : nil
Len : 0
Cap : 0
*/

/*
Slice Header &0x11110011
-------------
Ptr : 0x1122
Len : 5
Cap : 5
*/

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
