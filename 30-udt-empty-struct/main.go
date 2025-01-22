package main

import "fmt"

func main() {
	(&Empty{}).Greet()
	(&Empty1{}).Greet("Iello ICICI")

	e1 := new(Empty)
	e2 := new(Empty1)
	e1.Greet()
	e2.Greet("Hello ICICI Securities")
}

type Empty struct{} // there are many use cases

func (e *Empty) Greet() {
	fmt.Println("Hello World")
}

type Empty1 struct{}

func (e *Empty1) Greet(msg string) {
	fmt.Println(msg)
}
