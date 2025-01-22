package main

import "fmt"

func main() {

	c1 := Company{Name: "ICICI Securities", Website: "icicibank.com",
		Address: struct {
			City    string
			Pincode string
		}{City: "Mumbai", Pincode: "404040"}}

	c1.PrintDetails()
}

type Company struct {
	Name, Website string
	Address       struct { // emmbedded struct
		City, Pincode string
	}
}

func (c *Company) PrintDetails() {
	fmt.Println("Name:", c.Name)
	fmt.Println("Website:", c.Website)

	fmt.Println("City:", c.Address.City)
	fmt.Println("Pincode:", c.Address.Pincode)

}
