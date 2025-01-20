package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	i := 1
	println("Even numbers")
	for i <= 10 {
		if i%2 == 0 {
			print(i, " ")
		}
		i++
	}
	println()

	for num := 1; ; num++ {
		if num >= 10 {
			break
		}
		if num%2 == 0 {
			continue
		}
		print(num, " ")
	}
	println()

	a, b := 0, 1
	println("fib")
	for i := 1; i <= 20; i++ {
		print(a, " ")
		a, b = b, b+a // swap
	}

	num := 1
	for {
		if num == 10 {
			break
		}
		print(num, " ")
		num++
	}

	println()

	for num := 1; ; {
		if num == 10 {
			break
		}
		print(num, " ")
		num++
	}

	println()

	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	println()
	//done := false
	for i := 1; i <= 10; i++ {
		// if done {
		// 	break
		// }
		for j := 3; j <= 10; j++ {
			if i == j {
				//done = true
				break
			}
			fmt.Println("i", i, "-->", "j", j)
		}
	}

	println()
out:
	for i := 1; i <= 10; i++ {
		for j := 3; j <= 10; j++ {
			if i == j {
				break out
			}
			fmt.Println("i", i, "-->", "j", j)
		}
	}
outer:
	for {
		num := rand.IntN(100)
		switch {
		case num%9 == 0:
			println(num, "break now")
			break outer
		default:
			println(num)
		}
	}

	str1 := "hello world"

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), " ")
	}

	//str1 := "hello world"
	println()
	for i := 0; i < len(str1); i++ {
		print(str1[i], " ")
	}

}
