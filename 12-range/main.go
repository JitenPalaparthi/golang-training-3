package main

func main() {

	var str1 = "Hello 世界"
	var str2 = "Hello World"

	// for i, v := range str1 {

	// }

	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), " ")
	}
	println()
	for i := 0; i < len(str1); i++ {
		print(str1[i], "-->", string(str1[i]), " ")
	}

	println()
	for i, v := range str1 {
		println("i", i, "-->", string(v))

}

// range loop on string, array, slice --> index and value
// range loop on map returns key and value
// range loop on channel returns a value from the channel
