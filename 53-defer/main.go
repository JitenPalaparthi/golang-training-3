package main

func main() {
	i := 100

	defer func() { // i is global so there is an impact
		println(i)
	}()

	defer func(j int) { // i is copied to j so a separate copy of i is maintained
		println(j)
	}(i)

	func() {
		i = i + 100
		println(i)
	}()
}
