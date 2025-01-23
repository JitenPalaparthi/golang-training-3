package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello World")
	}()

	go func() {
		for i := 1; i <= 20; i++ {
			time.Sleep(time.Second * 1)
			print(i, "  ")
		}
	}()
	go func() {
		defer println("Hello Deffering this call")
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				fmt.Println("recovering this")
			}
		}()
		println("Hello ICICI")
		runtime.Goexit()
		fmt.Println("Can It be recovered")
	}()
	fmt.Println("end of main")

	runtime.Goexit()
	//time.Sleep(time.Second * 1)
	//os.Exit(1) // fatal
}

// 1. main is also a goroutine
// 2. By default no goroutine waits for other goroutine to complete its execution
// 3. Cannot decide the order of execution
