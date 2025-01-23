package main

import (
	"fmt"
	"sync"
)

var (
	counter int = 0
	wg          = new(sync.WaitGroup)
	mu          = new(sync.Mutex)
)

func main() {
	wg.Add(1)
	wg.Add(1)
	go Incr(wg, mu)
	go Decr(wg, mu)

	fmt.Println(counter)
	wg.Wait()
	//time.Sleep(time.Millisecond * 10)
	fmt.Println(counter)
}

func Incr(wg *sync.WaitGroup, mu *sync.Mutex) {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				mu.Lock()
				defer mu.Unlock()
				counter++
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Done()
}

func Decr(wg *sync.WaitGroup, mu *sync.Mutex) {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				mu.Lock()
				defer mu.Unlock()
				counter--
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Done()
}
