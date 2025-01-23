package main

import "fmt"

func main() {
	fmt.Println("start main")
	defer println("end of main") // this gets executed towards the end

	defer func() { // func-1
		println("\nstart of func-1")
		defer println("end of func-1") //-1
		for i := range 10 {
			defer println(i + 1)
		}
		//-1
	}()

	str := "Hello World"

	for _, v := range str {
		defer print(string(v))
	}

}
