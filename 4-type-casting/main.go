package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 uint8 = 100
	var num2 int64 = int64(num1)

	// var ok1 = true

	// var num3 uint8 = uint8(ok1)

	// bool cannot be casted to any type

	var float1 float32 = 123.23

	var num3 int64 = int64(float1)

	println(num1, num2, num3)

	var num4 int = 12312123123
	var num5 uint8 = uint8(num4)
	fmt.Printf("%b\n", num4)
	fmt.Println(num5)

	var num6 uint16 = uint16(num4)
	fmt.Println(num6)
	// 101101110111011100 0001011011110011
	// 11110011
	// 243
	// 5875

	var char1 rune = 65
	var char2 int32 = 'A'
	var char3 = 22000 // int
	println(string(char1))
	println(string(char2))
	println(string(char3))

	// str1 := strconv.Itoa(char3) // number as it is converted as string

	str1 := fmt.Sprint(char3)
	println(str1)
	str2 := "12323"
	num, err := strconv.Atoi(str2)
	if err != nil { // that means there is an error
		println(err.Error(), num)
	} else {
		fmt.Printf("atoi:%d Type:%T\n", num, num)
	}

	str2 = "1231.2334324"

	float2, err := strconv.ParseFloat(str2, 64)

	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(float2)
		fmt.Printf("%.3f\n", float2)
	}

	a1, p1 := Rect(12.12, 11.89)
	fmt.Println("Area:", a1, "perimeter", p1)
	//a, b, c, d := Calc(100.5, 200.3)

	a2, _ := Rect(12.12, 11.89)
	fmt.Println("Area:", a2)
	_, p2 := Rect(12.12, 11.89)
	fmt.Println("perimeter", p2)

	//strconv.ParseBool()
	//ParseBool returns the boolean value represented by the string.
	//It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
}

func Rect(l, b float32) (float32, float32) {
	add := l * b // creating new variable called add
	return add, 2 * (l + b)
}

func Calc(l, b float32) (add float32, sub float32, mul float32, div float32) {
	add = l + b
	sub = l - b
	mul = l * b
	div = l / b
	return add, sub, mul, div
}

// casting -> implicit, explicit
// conversion -> converting using functions or methods
// assertion -> it works with interfaces

// go does not have implicit conversion
