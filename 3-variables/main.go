package main

var G int = 12312       // stored in the data segment
const PI float32 = 3.14 // code/text segment
const MAX = 9999        // shorthand declarion of constant
const SEC = 60

func main() {
	var num = 100

	const (
		MIN     = 60 * 60
		HOUR    = 60 * MIN * SEC
		SOMEVAL = MIN/100 + 123 + (MIN + HOUR + MAX) // depends on all the operatnds are constants yes can do expression
	)
	{

		a, b := 10, 20 // stack
		// t := b
		// b = a
		// a = t
		a, b = b, a

		a, b, c := 10, 20, 30

		a, b, c = b, c, a

		println(a, b, c)
		a = 10
		b = 20
		c = 2
		c = 100/a + a*c + a + b - (100 + 100/4) // is it an atomic operation?
		//  10+    20   +30     -      125
		// -65
		println(a, b, c, num)
	}
	//println(a, b, c, num)

}
