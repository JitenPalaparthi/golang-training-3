package main

import (
	"github.com/JitenPalaparthi/icici-shapes-package3/shape"
	"github.com/JitenPalaparthi/icici-shapes-package3/shape/rect"
)

func main() {
	r1 := rect.NewRect(10.23, 12.23)
	shape.PrintShape(r1)
	shape.Greet()
}
