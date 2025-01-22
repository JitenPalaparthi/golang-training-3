package main

import (
	"errors"
	"fmt"
)

func main() {

	var m MyMap
	m = m.Init()

	m["560086"] = "Bangalore-1"
	m["560096"] = "Bangalore-2"
	m["560034"] = "Bangalore-3"
	m.Display()

	keys, err := m.GetKeys()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(keys)
	}

	if values, err := m.GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(values)
	}

	var m1 map[string]any
	m1 = make(map[string]any)

	m1["560086"] = "Bangalore-1"
	m1["560096"] = "Bangalore-2"
	m1["560034"] = "Bangalore-3"

	fmt.Println()
	MyMap(m1).Display()
	if keys, err := MyMap(m1).GetKeys(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keys)
	}

	if values, err := MyMap(m1).GetValues(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(values)
	}

}

type MyMap map[string]any

func (m MyMap) Init() MyMap {
	if m == nil {
		m = make(MyMap)
		return m
	}
	return m
}

func (m MyMap) GetKeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("nil MyMap")
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys, nil
}

func (m MyMap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil MyMap")
	}
	var values []any
	for _, v := range m {
		values = append(values, v)
	}
	return values, nil
}

func (m MyMap) Display() {
	if m == nil {
		fmt.Println("nil MyMap, no record")
		return
	}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}
