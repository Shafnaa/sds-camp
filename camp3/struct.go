package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var john Customer
	john.Name = "John Doe"
	john.Address = "Kalimanah"
	john.Age = 19

	fmt.Println(john.Name);
	fmt.Println(john.Address);
	fmt.Println(john.Age);
}