package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := Generator(time.Second*1, "Chan-1")
	// sig := Receive(ch1, ch2, ch3, ch4, ch5)
	// <-sig
	<-Receive(ch1)

}

func Generator(duration time.Duration, msg string) chan string {
	ch := make(chan string)
	//chTime := time.After(duration)
	timeout := TimeAfter(duration)
	done := false
	go func() {
		go func() {
			<-timeout
			done = true
			//runtime.Goexit()
		}()
		i := 1
		for {
			time.Sleep(time.Millisecond * 100)
			ch <- fmt.Sprint(msg, "-->", i*i)
			i++
			if done {
				close(ch)
				break
				//runtime.Goexit()
			}
		}
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

func TimeAfter(d time.Duration) chan struct{} {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(d)
		timeout <- struct{}{}
		close(timeout)
	}()
	return timeout
}
