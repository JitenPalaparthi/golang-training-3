package main

import (
	"time"
)

func main() {

	ch, sig := make(chan int), make(chan struct{})
	go Sender(ch, 20)
	go Receive(ch, sig)
	<-sig
}

func Sender(ch chan int, r int) {
	for i := 1; i <= r; i++ {
		time.Sleep(time.Millisecond * 110)
		ch <- i * i
	}
	close(ch) // closing the channel, closeing the channel is not about making the channel nil
}

func Receive(ch chan int, sig chan struct{}) {
	for v := range ch { // iterates until the channel is closed
		println(v)
	}
	sig <- struct{}{}
}
