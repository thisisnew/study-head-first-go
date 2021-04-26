package main

import (
	"fmt"
	magazine "study-head-first-go/Chapter8/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("PostalCode:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "4586 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"

	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("PostalCode:", employee.PostalCode)
}
