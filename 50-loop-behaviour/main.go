package main

func main() {
	funcs1 := make([]func(), 0)

	for i := 1; i <= 5; i++ {
		funcs1 = append(funcs1, func() {
			println(i)
		})
	}

	for _, v := range funcs1 {
		v()
	}

	i := 0
	funcs2 := make([]func(int), 0)
loop:
	i++
	funcs2 = append(funcs2, func(j int) {
		println(j)
	})
	if i <= 5 {
		goto loop
	}
	for i, v := range funcs2 {
		v(i)
	}
}

// <1.22
