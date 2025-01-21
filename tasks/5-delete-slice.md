func DeleteElem(slice []int, index uint) ([]int, error) {
	if slice == nil {
		return nil, errors.New("nil slice")
	}
	if int(index) >= len(slice) {
		return nil, errors.New("out of bounds")
	}

	// implement logic to delete an element from the slice

	// // what happens when the slice is nil
	// for _, v := range slice {
	// 	println(v)
	// }
	// // what happens when the slice is nil
	// slice = append(slice, 100) // slice is allocated and insert an element in the slice

	// slice[0] = 100 // when happen if the slice is nil
	// // actual implementation to remove an element from the slice

	return slice, nil

}

implement delete an element from the slice