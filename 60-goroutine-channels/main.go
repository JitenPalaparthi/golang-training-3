package main

func main() {
	ch := make(chan int, 2) // shorthand declartion
	//go func() {
	//time.Sleep(time.Second * 5)
	ch <- 100 // sending a value to the channel
	ch <- 200
	//ch <- 300
	//}()
	v := <-ch // receiving a value form the channel
	println(v)
	println("Hello main")
}

// make function is used to instantiate a channel
// bufferd and unbuffered channels.
// unbufferd channel means it can store one value and once the value is reveived from the channel , can send another valuer
// the sender is blocked until the receiver receives the value
// the receiver is blocked until the sender sends the value
// the block is based on the buffer.. if the buffer is full then sender and receiver are blocked

// live packet capture
