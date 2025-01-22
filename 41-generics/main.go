package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	r1 := AddG(100, 200)     // go does not create but c++ does a method for AddG1(int,int)
	r2 := AddG(10.34, 34.34) // go does not create but c++ doesa method for AddG2(float64,float64)

	var num1, num2 Myint = 10, 20
	r3 := AddG(num1, num2)    // // go does not create but c++ does a method for AddG3(Myint,Myint)
	r4 := AddG(uint8(10), 20) // both shjould be same type , if not Go tries to infer as same type
	fmt.Println(r1, r2, r3, r4)
	// AddG(true, true)
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

type Myint int

// func AddG[T 	~int | ~float64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~float32](a, b T) T {
func AddG[T Number](a, b T) T {
	return a + b
}

type Number interface {
	~int | ~float64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~float32
}

func AddI(a, b int) int {
	return a + b
}
func AddF(a, b float64) float64 {
	return a + b
}

// type Calc[T Number] struct {
// 	A, B T
// }
