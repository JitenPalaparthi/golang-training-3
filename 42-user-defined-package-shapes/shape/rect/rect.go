package rectangle

func init() {
	println("rect packaged is called-1")
}
func init() {
	println("rect packaged is called-2")
}

var (
	something string
	SomeThing string
)

type Cuboid struct {
	l, b, H float32 // l and b are restricted but H is unrestricted
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) What() string {
	return what()
}

func what() string {
	return "Rect"
}

// unrestricted or exported (~public)
// restricted or unexported (~private)
