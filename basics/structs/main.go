package main

import "fmt"

type employee struct {
	firstName  string
	lastName   string
	age        int
	married    bool
	salary     int
	profession string
}

func main() {
	emp := employee{"Salman", "Barani", 29, false, 2500, "Software engineer"}
	fmt.Println(emp)
}
