package main

import (
	"fmt"
	"math/rand/v2"
)

var G = rand.IntN(100)

var G1 = func(i int) int {
	return i * i
}(10)

var G2 int

func init() {
	G2 = func(i int) int {
		return i * i
	}(10)
	println(G1, G2)
}

func main() {
	fmt.Println("Hey")
	defer println("end of main")
	//defer recoverThis()
	func() { // func-1
		defer println("end of func-1-1")
		//defer recoverThis()
		func(u int) { // func-2
			defer println("end of func-2-1")
			defer recoverThis()
			println(100 / u) // know that this is going to be a panic
			fmt.Println("after panic")
		}(0)
		fmt.Println("func-1 ends")
	}()

	func() { //func-3 // 1.22
		for i := range 10 {
			fmt.Println(i)
		}
	}()
	fmt.Println("main ends-2")
}

// errors
// panic --> panic is a runtime error
// fatal

func recoverThis() {
	if v := recover(); v != nil {
		fmt.Println("Recovering from the following panic", v)
	}
}
