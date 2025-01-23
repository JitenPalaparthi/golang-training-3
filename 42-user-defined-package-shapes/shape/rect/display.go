package rectangle

import "fmt"

func (r *Rect) Display() {
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.B)
}
