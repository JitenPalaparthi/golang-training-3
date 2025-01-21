package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {

	var mymap map[string]any // declartion of the map
	if mymap == nil {
		println("map is nil:", unsafe.Sizeof(mymap))
	}

	mymap = make(map[string]any)

	mymap["560086"] = "Bangalore-1"
	mymap["560096"] = "Bangalore-2"
	mymap["560034"] = "Bangalore-3"

	v := mymap["560086"]
	println(v)

	for key, value := range mymap {
		fmt.Println(key, "-->", value)
	}

	//var mymap2 map[string]int = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3}
	mymap2 := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3}

	v1, ok := mymap2["12312312"]
	if !ok {
		println("key does not exist")
	} else {
		fmt.Println(v1)
	}

	_, ok1 := mymap2["12312312"]
	if !ok1 {
		println("key does not exist")
	}
	if keys, err := GetKeys(mymap); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Keys", keys)
	}

	if values, err := GetValues(mymap); err != nil {
		println(err.Error())
	} else {
		fmt.Println("Values", values)
	}

	//delete(mymap, "123123")
	if err := Delete(mymap, "12321"); err != nil {
		println(err.Error())
	} else {
		println("element successfully deleted")
	}

	// copy --> only slice
	// delete --> only with map
	// make --> only with slice,map and chan
	// append -> only with slice
	// clear --> slice and map
	// range-> keyword -> array, string , slice , map , chan , range of values and functions

}

// e08d99312033ce18299e 10c77822ec28a2bf7560
// 1							2

// map is created , some memory is allocated

func GetKeys(m map[string]any) ([]string, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	keys := make([]string, len(m)) // 3 elements
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys, nil
}

func GetValues(m map[string]any) ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	var values []any // nil ?
	for _, v := range m {
		values = append(values, v)
	}
	return values, nil
}

func Delete(m map[string]any, key string) error {
	if m == nil {
		return errors.New("nil map")
	}
	_, ok := m[key]
	if !ok {
		return fmt.Errorf("no such key as specified %v", key)
	}
	delete(m, key)
	return nil
}
