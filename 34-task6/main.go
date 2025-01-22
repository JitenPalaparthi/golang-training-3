package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	s1 := make(IntSlice, 5)

	for i := range s1 {
		s1[i] = rand.IntN(100)
	}
	fmt.Println(s1)

	mx, _ := s1.GetMax()
	fmt.Println("GetMax of IntSlice s1", mx)
	mi, _ := s1.GetMin()
	fmt.Println("GetMin of IntSlice s1", mi)

	arr1 := [5]int{10, 45, 832, 3, 59}

	mx, _ = IntSlice(arr1[:]).GetMax()
	mi, _ = IntSlice(arr1[:]).GetMin()
	fmt.Println("GetMax of arr1", mx)
	fmt.Println("GetMin of arr1", mi)

}

type IntSlice []int

func (is IntSlice) GetMax() (int, error) {
	if is == nil {
		return 0, errors.New("nil slice")
	}
	if len(is) == 1 {
		return is[0], nil
	}
	mx := is[0]
	for i := 1; i < len(is); i++ {
		mx = max(mx, is[i])
	}
	return mx, nil
}

func (is IntSlice) GetMin() (int, error) {
	if is == nil {
		return 0, errors.New("nil slice")
	}
	if len(is) == 1 {
		return is[0], nil
	}
	mi := is[0]
	for i := 1; i < len(is); i++ {
		mi = min(mi, is[i])
	}
	return mi, nil
}
