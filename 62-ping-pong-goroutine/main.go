package main

import (
	"fmt"
	. "fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	sig := make(chan struct{})
	go Ping(ch1, ch2, sig)
	go Pong(ch1, ch2)
	<-sig // When does it release?
	// future pattern
}

func Ping(chPing chan<- string, chPong <-chan string, sig chan<- struct{}) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 100)
		chPing <- Sprint("ping->", i)
		fmt.Println(<-chPong)
	}
	sig <- struct{}{}
}

func Pong(chPing <-chan string, chPong chan<- string) {
	for i := 1; i <= 100; i++ {
		fmt.Println(<-chPing)
		chPong <- Sprint("pong->", i)
	}
}
