package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := Generator(20, "Chan-1")
	ch2 := Generator(10, "Chan-2")
	ch3 := Generator(30, "Chan-3")
	ch4 := Generator(30, "Chan-4")
	ch5 := Generator(30, "Chan-5")
	// sig := Receive(ch1, ch2, ch3, ch4, ch5)
	// <-sig
	<-Receive(ch1, ch2, ch3, ch4, ch5)
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

func Receive(chs ...chan string) chan struct{} {
	wg := new(sync.WaitGroup)
	sig := make(chan struct{})
	go func() {
		for _, ch := range chs {
			wg.Add(1)
			go func() {
				for c := range ch {
					println(c)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
