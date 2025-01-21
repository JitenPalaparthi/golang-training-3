package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice1 := make([]int, 10)
	for i := range slice1 {
		slice1[i] = rand.Intn(10)
	}

	PrintSliceHeader(&slice1, "slice1")
	DoubleSlice(&slice1)
	PrintSliceHeader(&slice1, "slice1")
	DoubleSliceAdd(&slice1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	PrintSliceHeader(&slice1, "slice1")
}

func DoubleSlice(slice *[]int) {
	for i, v := range *slice {
		(*slice)[i] = v * v
		//*slice[i] = v * v
	}
}

func DoubleSlice2(slice []int) {
	for i, v := range slice {
		slice[i] = v * v
	}
}

func DoubleSliceAdd(slice *[]int, nums ...int) {
	*slice = append(*slice, nums...)
	for i, v := range *slice {
		(*slice)[i] = v * v
	}
}

func PrintSliceHeader(slice *[]int, name string) { //slice []*int
	fmt.Println(name)
	fmt.Println(*slice)
	fmt.Printf("Address of Slice:%p\n", slice)
	fmt.Printf("Len:%d\n", len(*slice))
	fmt.Printf("Cap:%d\n", cap(*slice))
	if len(*slice) > 0 {
		fmt.Printf("Ptr:%p\n", &(*slice)[0])
	}
	println("------------")
}
