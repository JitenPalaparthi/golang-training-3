package main

import "fmt"

func main() {

	str1 := "Hello World"
	// str1 also contains a pointer
	/*
		String Header 0x14000104020
		------------
		Ptr: 0x14000102020
		Len: 11
	*/

	fmt.Printf("Length:%d\nAddress of str1:%p\naddress of 0th elenment:%p\n", len(str1), &str1, &([]byte(str1)[0]))
	str1 = "Hello ICICI"
	fmt.Printf("Length:%d\nAddress of str1:%p\naddress of 0th elenment:%p\n", len(str1), &str1, &([]byte(str1)[0]))
	// {
	// 	a, b := 10, 20
	// } // drop(a) drop(b)
	// prinltn(a, b)
}
