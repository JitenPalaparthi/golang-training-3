package main

import (
	"fmt"
	"unsafe"
)

func main() {
	t1 := T1{}
	fmt.Println(unsafe.Sizeof(t1))
	fmt.Println(unsafe.Sizeof(T2{}))

	// builtin functions real,imag, max and min
	c1 := 123.3 + 1.1i
	r, im := real(c1), imag(c1)
	mx := max(10123, 123)
	mi := min(123.13, 23)
	fmt.Println(r, im, mx, mi)

	i := 0
loop:
	if i > 10 {
		goto exit
	}
	i++
	if i%2 == 0 {
		goto even
	} else {
		goto odd
	}
even:
	fmt.Println("even:", i)
	goto loop
odd:
	fmt.Println("odd:", i)
	goto loop
exit:
}

type T1 struct {
	B1 bool    //1
	I1 int32   //4
	B2 bool    //1
	I2 int64   //8
	B3 bool    //1
	F1 float64 //8
} // 23 bytes

// B1 I1
// B1---- --- B2------- I2 B3------ F1 --> 40
type T2 struct {
	I2 int64   //8
	F1 float64 //8
	I1 int32   //4
	B1 bool    //1
	B2 bool    //1
	B3 bool    //1
} // 23 bytes
