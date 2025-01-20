package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	var any1 any = 12312

	println(IsNumber(any1))
	any1 = 12312.123
	println(IsNumber(any1))
	any1 = true
	println(IsNumber(any1))

	var any2, any3 any = 123.123, 123.23
	r1, err := Add(any2, any3)
	if err != nil {
		println(err.Error())
	} else {
		println("add r1:", r1)
	}

	r2, err := Add(true, false)
	if err != nil {
		println(err.Error())
	} else {
		println("add r2:", r2)
	}

}

func Add(a, b any) (any, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, fmt.Errorf("a and b must be of same type")
	}
	if !IsNumber(a) {
		return nil, errors.New("not a number type.can perform arithmetic operation only on Number types")
	}
	switch v := a.(type) {
	case int:
		return a.(int) + b.(int), nil
	case uint:
		return a.(uint) + b.(uint), nil
	case uint8:
		return uint16(a.(uint8) + b.(uint8)), nil
	case uint16:
		return a.(uint16) + b.(uint16), nil
	case uint32:
		return a.(uint32) + b.(uint32), nil
	case uint64:
		return a.(uint64) + b.(uint64), nil
	case int8:
		return int16(a.(int8) + b.(int8)), nil
	case int16:
		return a.(int16) + b.(int16), nil
	case int32:
		return a.(int32) + b.(int32), nil
	case int64:
		return a.(int64) + b.(int64), nil
	case float32:
		return v + b.(float32), nil //' can use the value v as well
	case float64:
		return v + b.(float64), nil
	}

	return nil, nil
}

func IsNumber(a any) bool {
	switch a.(type) {
	case int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}
