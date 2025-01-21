package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	slice1[0], slice1[1], slice1[2] = 10, 11, 12
	//fmt.Println("Slice1", slice1)
	slice2 := slice1 // header copy
	fmt.Printf("Address of slice1:%p\n", &slice1)
	fmt.Printf("Address of slice2:%p\n", &slice2)

	slice2[0] = 1111
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	slice2 = append(slice2, 2222)
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)
	slice2[1] = 8888
	slice2 = append(slice2, 3333)
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	slice2 = append(slice2, 4444)
	slice2[2] = 0
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)
	slice1 = slice2 // header copy
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)
	//clear(slice1)
	// fmt.Println("Slice1", slice1)
	// fmt.Println("Slice2", slice2)

	slice3 := make([]int, 3)
	copy(slice3, slice1) // deep copy, only the values are copied
	fmt.Println(slice3)
	slice3[0] = 9999
	fmt.Println("Slcie3:", slice3)
	fmt.Println("slice1", slice1)
}
