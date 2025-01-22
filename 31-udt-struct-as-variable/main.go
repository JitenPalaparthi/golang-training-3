package main

import "fmt"

func main() {
	var p1 struct{ Name, Email string } = struct {
		Name  string
		Email string
	}{Name: "Jiten", Email: "Jitenp@outlook.com"}

	p1.Email = "Jiten.Palaparthi@Gmail.Com"
	fmt.Println(p1)
}

// func PrintDetails(p )
