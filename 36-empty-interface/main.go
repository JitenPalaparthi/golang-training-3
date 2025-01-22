package main

import "fmt"

func main() {
	//var any1 interface{} // empty interface
	var iempty IEmpty = 123
	iempty = "Hello World"
	iempty = 123.123

	var f1 float64
	f1, ok := iempty.(float64)
	if ok {
		fmt.Println(f1)
	} else {
		fmt.Println(ok, "could not be asserted")
	}

	iempty = true

}

type IEmpty interface{}

// interface is contact, agreement
// shared behaviour
