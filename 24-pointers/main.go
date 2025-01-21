package main

import "unsafe"

func main() {

	arr := [5]int{10, 11, 12, 13, 14}
	// ptr += 8 // pointer arithmetic
	var uintptr1 uintptr

	uintptr1 = uintptr(unsafe.Pointer(&arr[0]))

	for i := 0; i < len(arr); i++ {
		v := *(*int)(unsafe.Pointer(uintptr1))
		println(v)
		uintptr1 += 8
	}

	//v3 := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 8))
	//println(v3)

	str1 := "Hello ICICI"

	uintptr2 := uintptr(unsafe.Pointer(&str1))
	uintptr2 += 8
	v := (*int)(unsafe.Pointer(uintptr2))
	*v = 20
	println(*v)
	println(len(str1))

	for i, v := range str1 {
		println(i, "-->", string(v))
	}
}

/*
A pointer value of any type can be converted to a Pointer.
A Pointer can be converted to a pointer value of any type.
A uintptr can be converted to a Pointer.
A Pointer can be converted to a uintptr.
*/
