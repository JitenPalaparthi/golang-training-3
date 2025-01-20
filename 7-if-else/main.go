package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var age uint8 = 18
	if age >= 18 {
		println("eligible for vote becase of age", age)
	} else {
		println("not eligible for vore because of age", age)
	}

	str1 := "1231er2"

	if num, err := strconv.Atoi(str1); err != nil {
		fmt.Println(err.Error())
	} else {
		println(num)
	}

	fmt.Println(Calc(123, 23))
	fmt.Println(Calc(123.123, 23.232))
	fmt.Println(Calc(true, false))
	fmt.Println(Calc("hello", "World"))
	fmt.Println(Calc(uint8(232), uint8(123)))
}

func Calc(a, b any) (any, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("a and b are two different types")
	}
	if _, ok := a.(int); ok {
		return a.(int) + b.(int), nil
	}
	if _, ok := a.(uint); ok {
		return a.(uint) + b.(uint), nil
	}
	if _, ok := a.(int8); ok {
		return int16(a.(int8)) + int16(b.(int8)), nil
	}
	if _, ok := a.(int16); ok {
		return a.(int16) + b.(int16), nil
	}
	if _, ok := a.(int32); ok {
		return a.(int32) + b.(int32), nil
	}
	if _, ok := a.(int64); ok {
		return a.(int64) + b.(int64), nil
	}
	if _, ok := a.(uint8); ok {
		return uint16(a.(uint8)) + uint16(b.(uint8)), nil
	}
	if _, ok := a.(uint16); ok {
		return a.(uint16) + b.(uint16), nil
	}
	if _, ok := a.(uint32); ok {
		return a.(uint32) + b.(uint32), nil
	}
	if _, ok := a.(uint64); ok {
		return a.(uint64) + b.(uint64), nil
	}
	if _, ok := a.(float32); ok {
		return a.(float32) + b.(float32), nil
	}
	if _, ok := a.(float64); ok {
		return a.(float64) + b.(float64), nil
	}
	return nil, errors.New("invalid types to perorm arithmetic operation")
}
