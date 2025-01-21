package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	slice2 := slice1 // header copy
	PrintSliceHeader(slice1, "slice1")
	PrintSliceHeader(slice2, "slice2")
	slice3 := slice1[:5] // from 0th index to the 4th index but not 5
	PrintSliceHeader(slice3, "slice3")
	slice4 := slice1[2:8]
	PrintSliceHeader(slice4, "slice4")
	// slice4[0] = 9999
	// fmt.Println(slice1)
	slice5 := slice1[5:]
	PrintSliceHeader(slice5, "slice5")
	slice5 = append(slice5, 9999, 1111)
	slice5[0] = 7777
	fmt.Println(slice1)
	println("delete an element from the slice")
	slice6 := append(slice1[0:3], slice1[4:]...)
	fmt.Println(slice6)

}

func PrintSliceHeader(slice []int, name string) {
	fmt.Println(name)
	fmt.Println(slice)
	fmt.Printf("Address of Slice:%p\n", &slice)
	fmt.Printf("Len:%d\n", len(slice))
	fmt.Printf("Cap:%d\n", cap(slice))
	if len(slice) > 0 {
		fmt.Printf("Ptr:%p\n", &slice[0])
	}
	println("------------")
}
