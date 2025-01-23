package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for c := range ch {
			println(c)
		}
	}()

	wg := new(sync.WaitGroup)
	wg.Add(4)
	go senderSq(wg, ch, 10, "sender-1")
	go senderSq(wg, ch, 15, "sender-2")
	go senderCube(wg, ch, 20, "sender-3")
	go senderCube(wg, ch, 5, "sender-4")
	wg.Wait()
	close(ch)
}

func senderSq(wg *sync.WaitGroup, ch chan string, n uint, name string) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- fmt.Sprint(name, " Sq-->", i*i)
	}
	wg.Done()
}

func senderCube(wg *sync.WaitGroup, ch chan string, n uint, name string) {
	for i := 0; i <= int(n); i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- fmt.Sprint(name, " Cube-->", i*i*i)
	}
	wg.Done()
}
