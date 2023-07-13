package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type employee struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	emp := employee{
		firstName: "Salman",
		lastName:  "Barani",
		contactInfo: contactInfo{
			email: "salmanAndB@outlook.com",
			phone: "00905539343945",
		},
	}
	emp.updateFirstName("Johnny")
	emp.print()
}

func (e employee) updateFirstName(firstName string) {
	e.firstName = firstName
}

func (e employee) print() {
	fmt.Printf("%+v", e)
}
