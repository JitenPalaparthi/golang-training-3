package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	fmt.Println("using goto")

	i := 1
loop:
	func() {
		k := i * 10
		fmt.Println(k)
	}()
	i++
	if i <= 10 {
		goto loop
	}

}
