package main

import (
	"fmt"
	"reflect"
)

func main() {

	arr1 := [5]int{10, 11, 12, 13, 14} //shorthand declaration
	arr2 := [...]int{10, 11, 12, 13}   // shorthand declaration
	fmt.Println("Arr1:", reflect.TypeOf(arr1))
	fmt.Println("Arr2:", reflect.TypeOf(arr2))

	SumOf(arr1)
	_ = SumOf4(arr2)

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	println("2d array")
	for _, a1 := range arr2d {
		for _, v := range a1 {
			print(v, " ")
		}
		println()
	}

	println("arr3d")

	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				print(v, " ")
			}
			println()
		}
	}

}

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf4(arr [4]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
