package main

func main() {

	var age uint = 17
	var gender rune = 'F'
	//println('m')
	//println(age >= 18 && (gender == 'f' || gender == 70) && true && ((false || true) || false))
	if age >= 18 && (gender == 'f' || gender == 70) && true && ((false || true) || false) {
		// age > =18 && true && true && (true || flase)
		// false && true && true && true
		// false && true && true
		// false && true
		// false
		println("She is eligible for marriage becase of age is ", age)
	} else if age >= 21 && (gender == 'M' || gender == 109) {
		println("He is eligible for marriage becase of age is ", age)
	} else {
		println("not eligible for marriage")
	}

}
