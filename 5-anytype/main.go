package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var any1 interface{}
	var any1 any // can hold any kind of variable or value
	// what is the zero value of any   --> nil
	// what is the default type of any --> nil

	if any1 == nil {
		println("any1 is nil")
	}

	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))
	str1 := "Hello World"
	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)
	//unsafe.Sizeof()

	any1 = 12312 // int

	var num1 int = any1.(int)
	var float1 float32 // runtime error

	float1, ok := any1.(float32) // runtime error
	if ok {
		println(float1)
	} else {
		println("could not assert to float32")
	}
	println(num1)
	any1 = true //
	//var ok1 = bool(any1)
	any1 = str1
	//

	any1 = 123123.12312 // float64

	float2, ok := any1.(float64) // runtime error
	if ok {
		fmt.Println(float2)
	} else {
		println("could not assert to float64")
	}
	//var float1 float32 = any1.(float32) // runtime error

	//var any2 any = any1

}

// String header
// Ptr
// Len

// Any Header
// DataPtr --> The pointer to the actual data
// TypePtr --> The pointer to the type information
