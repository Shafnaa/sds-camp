package main

import "fmt"

type Employee struct {
	name string
	empid int
}

func main() {
	emp := Employee{"ABC", 19078}

	pts := &emp

	fmt.Println(pts)

	fmt.Println(pts.name)

	fmt.Println((*pts).name)

}
