package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var num1 int
	var ok1 bool    // zero value is false
	var str1 string // zero value is ""

	println(num1, ok1, str1)

	var (
		num2   uint8   = 100
		num3   uint16  = 12312
		float1 float32 = 12312.123
		float2 float64 = 123.12312312
	)
	println(num2, num3, float1, float2)

	var num4 = 12321312 // type inference
	// num4 is int
	var age uint8 = 39 // age is inferred as int, uint8
	float3 := 123.23   // short hand declaration

	ok2 := true // bool

	str2 := "Hello world" // automatically created as string based on the value
	println(num4, age, float3, ok2, str2)

	var char1 rune = 'A' + 1
	var char2 int32 = 'B'
	var char3 uint8 = 'C'
	var char4 rune = '你'
	fmt.Println(char1, char2, char3, char4)

	var str5 = "Hello World"                  // 11 chars , 11 bytes
	var str6 = "Hello 世界"                     // 8 chars , 12 bytes
	str7 := "Hello World, How are you doing?" // 31 charss, 31 bytes

	fmt.Println("Size of str5:", unsafe.Sizeof(str5), "Len:", len(str5))
	fmt.Println("Size of str6:", unsafe.Sizeof(str6), "Len:", len(str6))
	fmt.Println("Size of str7:", unsafe.Sizeof(str7), "Len:", len(str7))

	fmt.Println(str5, str6, str7)

	str5 = "Hello Universe" // mutate

	// 你好，世界

	c1 := 22 + 1.1i        // complex number
	c2 := complex(22, 1.1) // complex number

	var r1, im1 float32 = 12.1, 1.1
	c3 := complex(r1, im1)

	fmt.Println(c1, c2, c3)
}

// 0-255 -- 1 byte
// 256- 16k -2
//

type integer = int // alias for int

/* numbers
int, int8,int16,int32,int64
uint,uint8,uint16,uint32,uint64
float32, float64
complex64, complex128
rune, byte // these are not new types just alias for int32 and uint8
*/

// boolean --> bool
// strings --> string

// zero value-->
// type inference -->type is inferred based on the value that is assigned

// string header
// Ptr -> 8 bytes
// Len -> 8 bytes
