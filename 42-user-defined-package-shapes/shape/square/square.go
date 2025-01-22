package square

type Square float32

func NewSquare(s float32) Square {
	return Square(s)
}

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

func (s Square) What() string {
	return "Square"
}
