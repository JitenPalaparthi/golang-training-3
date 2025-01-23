package main

import (
	"sync"
	"time"
)

// wg maintains a state /counter

func main() {
	var wg = new(sync.WaitGroup)
	defer wg.Wait() // wait until? the state becomes 0

	//wg.Add(7)
	wg.Add(1)
	go func() {
		wg.Add(1)
		go func() {
			for i := 1; i <= 10; i++ {
				time.Sleep(time.Millisecond * 100)
				println("Natural Numbers:", i)
			}
			wg.Done()
		}()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				time.Sleep(time.Millisecond * 100)
				println("Even Numbers:", i)
			}
		}
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				time.Sleep(time.Millisecond * 100)
				println("Odd Numbers:", i)
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go Fib1(wg, 20)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		Fib2(20)
		wg.Done()
	}(wg)

	wg.Add(1)
	go func() {
		Fib2(20)
		wg.Done()
	}()

	println("End of main")
}

func Fib1(wg *sync.WaitGroup, r int) {
	a, b := 0, 1
	for i := 1; i <= 10; i++ {
		a, b = b, a+b
		time.Sleep(time.Millisecond * 100)
		println("fib:", a)
	}
	wg.Done()
}

func Fib2(r int) {
	a, b := 0, 1
	for i := 1; i <= 10; i++ {
		a, b = b, a+b
		time.Sleep(time.Millisecond * 100)
		println("fib:", a)
	}
}
