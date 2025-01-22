package main

import "fmt"

func main() {
	//	num := 0
	//	fmt.Println(100 / num)

	// var ptr *int
	// *ptr = 100

	// arr := [1]int{100}
	// for i := 0; i <= len(arr); i++ {
	// 	println(arr[i])
	// }
	func() { // func-1
		func(u int) { // func-2
			println(100 / u) // know that this is going to be a panic
		}(0)
		fmt.Println("func-1 ends")
	}()

	func() { //func-3 // 1.22
		for i := range 10 {
			fmt.Println(i)
		}
	}()

	fmt.Println("main ends")

}

// errors
// panic --> panic is a runtime error
// fatal
