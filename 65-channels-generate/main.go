package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- Add(10, 20)
	}()
	println(<-ch)
	ch1 := Generator(20, "Gen-1")
	ch2 := Generator(10, "Gen-2")
	done1, done2 := false, false
	sig := make(chan struct{})
	go func() {

		for {
			if done1 && done2 {
				sig <- struct{}{}
				runtime.Goexit()
			}
			select {
			case v1, ok1 := <-ch1:
				if ok1 {
					println(v1)
				} else {
					done1 = true
				}
			case v2, ok2 := <-ch2:
				if ok2 {
					println(v2)
				} else {
					done2 = true
				}
			}
		}
	}()
	<-sig
}

func Add(a, b int) int {
	return a + b
}

func Generator(n uint, msg string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i <= int(n); i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- fmt.Sprint(msg, "-->", i*i)
		}
		close(ch)
	}()
	return ch
}
