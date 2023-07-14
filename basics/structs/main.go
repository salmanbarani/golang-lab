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

	empPointer := &emp

	empPointer.updateFirstName("Johnny")
	emp.print()
}

func (poniterToEmployer *employee) updateFirstName(firstName string) {
	(*poniterToEmployer).firstName = firstName
}

func (e employee) print() {
	fmt.Printf("%+v\n", e)
}
