package main

import (
	"fmt"
)

func main() {
	// db.Select("users").Where("id=?",id).Limit(10).Offset(20).Desc().Result
	c1 := New(200)
	r := c1.Add(100).Sub(200).Mul(10).Div(5).Add(100).Get()
	fmt.Println(r)

	cr1 := NewCircle()

	cr1.SetR(100.34).SetBGColour("red").SetBorderColour("blue")
	fmt.Println(cr1)
	cr1.SetBorder(12.2)
	fmt.Println(cr1)
	//println(G)
}

func Sq(n int) int {
	G = n * n
	return G
}

var G int //= Cube() // Global variable

func init() {
	fmt.Println("I am called-1")
}
func Cube() int {
	return 3 * 3 * 3
}

func init() { // Called before main
	G = Sq(100) // the value of G is given in the init ..
	fmt.Println("init -->", G)
}

func init() {
	fmt.Println("I am called-2")
}

func init() {
	fmt.Println("I am called-3")
}

func init() {
	fmt.Println("I am called-4")
}
