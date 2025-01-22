package main

import "fmt"

func main() {
	var p1 struct {
		Name, Email string
		Address     struct{ City string }
	} = struct {
		Name    string
		Email   string
		Address struct{ City string }
	}{Name: "Jiten", Email: "Jitenp@outlook.com",
		Address: struct{ City string }{City: "Mumbai"}}

	PrintDetails(p1)

	var p2 struct {
		Name, Email string
		Address     struct{ City string }
	}

	p2.Name = "Priya"
	p2.Email = "Priya@Outlook.com"
	p2.Address.City = "Bengalore"
	PrintDetails(p2)
}

func PrintDetails(p struct {
	Name, Email string
	Address     struct{ City string }
}) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("City:", p.Address.City)
}

// type Person struct {
// 	Name, Email string
// }
