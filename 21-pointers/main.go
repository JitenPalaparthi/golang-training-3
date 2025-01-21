package main

func main() {
	var num int = 100
	var ptr1 *int = &num
	*ptr1 = 200
	// ptr1++ // pointer arithmetic is not allowed go
	// there is away to do that
	var ok bool = true
	var ptrOk *bool = &ok

	*ptrOk = false
	*(&num) = 300

	var ptr2 *int // what is zero value of pointer?
	if ptr2 == nil {
		println("nil pointer")
	}

	ptr2 = new(int) // built in function, 8 bytes is allocarted on 64 bit machines
	// it allocates memory based on the type provided and assignes zero value to that
	*ptr2 = 100
	println(*ptr2)
	// if Greet == nil {
	// 	println("not nil")
	// }

	ptr3 := new(string)
	// 16 bytes
	// String Header
	// Ptr: nil
	// Len: 0

	*ptr3 = "Hello ICICI Securities"

	var ptr4 **int = &ptr1
	var ptr5 ***int = &ptr4

	**ptr4 = 1000
	***ptr5 = 2000
	println(num)

	n := 100
	Incr(n)
	println(n)
	IncrP(&n)
	println(n)

}

// func Greet() {
// 	println("hello ICICI Direct")
// }

// string -> Ptr, Len

// can check nil
// slice --> Ptr, Len, Cap
// any/interface --> DataPtr, TypePtr
// func -> Ptr,
// map -> Ptr
// chan -> Ptr
