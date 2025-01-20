package main

func main() {

	var day uint8 = 4

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	num := -49
	switch { // empty switch
	case num >= 0 && num <= 50:
		println(num, "is 0-50")
	case num > 50 && num <= 100:
		println(num, "is 51-100")
	case num >= 101:
		println(num, "is more than 100")
	default:
		println(num, "is negative number")
	}

	char := 'A'
	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is vovel")
	default:
		println(string(char), "is either a consonent or a special char")
	}

	num = 48

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("false negative")
	num = 2
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	switch day1 := 5; day1 {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

}
