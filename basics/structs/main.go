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

	teams := "SAlman"
	ch(teams)
	fmt.Println(teams)

}

func (poniterToEmployer *employee) updateFirstName(firstName string) {
	(*poniterToEmployer).firstName = firstName
}

func ch(teams string) {
	teams = "1"
}

func (e employee) print() {
	fmt.Printf("%+v\n", e)
}
