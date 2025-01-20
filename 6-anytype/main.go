package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		num1   int     = 1231
		float1 float32 = 1231.123
		any1   any     = 123123.123
		any2   any     = float1
		str1   string  = "12312.223"
		str2   string  = "123123"
		sum    float64
	)

	float2, _ := strconv.ParseFloat(str1, 64)
	num2, _ := strconv.Atoi(str2)
	sum = float64(num1) + float64(float1) + any1.(float64) + float64(any2.(float32)) + float2 + float64(num2)
	fmt.Println(sum)
}
