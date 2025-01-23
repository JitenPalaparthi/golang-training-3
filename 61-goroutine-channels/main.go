package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	wg.Add(1)
	go func() {
		send(ch)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		receive((ch))
		wg.Done()
	}()

}

func send(ch chan int) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 200)
		ch <- i
	}
}

func receive(ch chan int) {
	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
	}
}
