package main

type IShape interface {
	Area() float32
	Perimeter() float32
	IWhat
}

type IWhat interface {
	What() string
}
