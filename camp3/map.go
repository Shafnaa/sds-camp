package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "John",
		"age": "16",
	}

	company := map[string]string{
		"name": "John Company",
		"industry": "FnB",
	}

	fmt.Println(person["name"]);
	fmt.Println(person["age"]);

	fmt.Println(company["name"]);
	fmt.Println(company["industry"]);
}