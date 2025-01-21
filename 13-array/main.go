package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var arr1 [5]int

	var arr2 [5]any

	// What are zero values?
	// What is the type of an array
	fmt.Println(arr1)
	fmt.Println(arr2)

	for i := range arr1 {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println(arr1)

	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	println(sum)

	sum = 0

	for _, v := range arr1 {
		sum += v
	}
	println(sum)

	// arr2[0] = int(100)
	// arr2[1] = float32(123.23)
	// arr2[2] = float64(123.23)
	// arr2[3] = int64(123123)
	// arr2[4] = uint8(255)
	// Find out the sum of elements
}

// array
// fixed length

// string header
// Ptr --> Pointer to the actual store
// Len
