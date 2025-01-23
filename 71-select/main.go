package main

import (
	"fmt"
	"time"
)

func main() {

	ch1, ch2, ch3, timeout := make(chan string), make(chan string), make(chan string), time.After(time.Millisecond*100)
	go Send(ch1, "Server-1")
	go Send(ch2, "Server-2")
	go Send(ch3, "Server-3")

	select {
	case v := <-ch1:
		println(v)
	case v := <-ch2:
		println(v)
	case v := <-ch3:
		println(v)
	case t := <-timeout:
		fmt.Println("timed out-->", t)
	}
}

func Send(ch chan<- string, name string) {
	time.Sleep(time.Millisecond * 100)
	ch <- name
}
