package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	wg.Add(1)
	go Incr(wg)
	go Decr(wg)

	fmt.Println(counter)
	wg.Wait()
	//time.Sleep(time.Millisecond * 10)
	fmt.Println(counter)
}

func Incr(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				counter++
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Done()
}

func Decr(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				counter--
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Done()
}
