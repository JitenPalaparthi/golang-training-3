package main

import "fmt"

func main() {
	var e1 Employee
	e1.Id = 100
	e1.Name = "Jiten"
	e1.Email = "Jitenp@outlook.com"
	fmt.Println(e1)

	var e2 Employee = Employee{Id: 102, Name: "Ravi", Email: "Ravi@Gmail.Com"}
	e3 := Emp{Id: 102, Name: "Ravi", Email: "Ravi@Gmail.Com"}
	fmt.Println(e2, e3)
	//var i Integer = 100
	// rune, byte , any

	cc1 := ColourCode{int: 100, string: "Red", rgbString: "90 90 90"}
	fmt.Println("Code", cc1.int)
	fmt.Println("Name", cc1.string)
	fmt.Println("RGB Ratio", cc1.rgbString)

}

// struct
// other user defiend types

type Integer = int // not creating a new type , just alias
type Emp = Employee

type Employee struct {
	Id    int
	Name  string
	Email string
}

type rgbString = string

type ColourCode struct { // struct without field names, anonymous fields
	int
	string
	rgbString
}
