package main

import "fmt"

func main() {
	p1 := Person{Name: "Jiten", Email: "JitenP@Outlook.com", Address: Address{City: "Bengaluru", Pincode: "560086"}}
	p1.PrintDetails()

	c1 := new(Company)
	c1.Name = "ICICI Direct"
	c1.Website = "icicibank.com"
	c1.City = "Mumbai"
	c1.Pincode = "404040"
	c1.Status = "active"
	c1.Address.Status = "active"

	c1.PrintDetails()
	c1.PrintAddress()

}

type Person struct {
	Name, Email string
	Address     Address // composition
	//ShippingDetails ShippingDetails
}

func (p *Person) PrintDetails() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	p.Address.PrintAddress()
}

type Address struct {
	City, Pincode, Status string
}

func (a *Address) PrintAddress() {
	fmt.Println("City:", a.City)
	fmt.Println("Pincode:", a.Pincode)
	fmt.Println("Status:", a.Status)

}

type Company struct {
	Name, Website, Status string
	Address               // promoted filed
}

func (c *Company) PrintDetails() {
	fmt.Println("Name:", c.Name)
	fmt.Println("Website:", c.Website)
	fmt.Println("Status:", c.Status)
}
