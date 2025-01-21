package main

import "fmt"

func main() {

	fmt.Println("hello", "Wrold", true, (100 + 200))
	var slice1 []int // nil slice
	slice1 = append(slice1, 100, 101, 102, 103, 104, 105)
	fmt.Println(slice1)

	s1 := SumOf()
	s2 := SumOf(10, 20)
	s3 := SumOf(10, 20, 30, 40, 50, 60, 70, 123, 123, 12, 5, 45, 45, 45, 34, 46, 34, 7, 8, 9)
	fmt.Println(s1, s2, s3)
	s4 := SumOf(slice1...) // This is understood that slice is converted and passed as a variadic argument
	arr1 := [5]int{10, 20, 30, 40, 50}
	s5 := SumOf(arr1[:]...)
	fmt.Println(s4, s5)

	// v1 := arr1[:] // slice , ptr, len and cap
	// v2 := arr1    // array

}

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// variadic parameter must be the last parameter in a func or method
// other than functions or methods, cannot use variadic parameter
// use range loop to iterate variadic argument
// slice can be passed as an argument as a variadic variable
