package main

import "fmt"

func main() {

	var num1 MyInt1 = 123
	var num2 int = 123
	var num3 Myint2 = 123
	var num4 MyInt3 = 123
	var float1 float32 = 12312.123

	str1 := num1.ToString()
	s1 := Myint2(num1).Sq()
	c1 := MyInt3(num1).Cube()
	fmt.Println(str1, s1, c1)

	str2 := MyInt1(num2).ToString()
	s2 := Myint2(num2).Sq()
	c2 := MyInt3(num2).Cube()
	fmt.Println(str2, s2, c2)

	str3 := MyInt1(float1).ToString()
	s3 := Myint2(float1).Sq()
	c3 := MyInt3(float1).Cube()
	fmt.Println(str3, s3, c3)

	str4 := MyInt1(num3).ToString()
	s4 := num3.Sq()
	c4 := MyInt3(num3).Cube()
	fmt.Println(str4, s4, c4)

	str5 := MyInt1(num4).ToString()
	s5 := Myint2(num4).Sq()
	c5 := num4.Cube()
	fmt.Println(str5, s5, c5)

}

type Integer = int // this is an alias not a new type

type MyInt1 int // this is a new type

func (m MyInt1) ToString() string {
	return fmt.Sprint(m)
}

type Myint2 int

func (m Myint2) Sq() Myint2 {
	return m * m
}

type MyInt3 Myint2

func (m MyInt3) Cube() MyInt3 {
	return m * m * m
}
