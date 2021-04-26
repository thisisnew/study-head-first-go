package main

import (
	"fmt"
	magazine "study-head-first-go/Chapter8/magazine"
)

func main() {
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "6811",
	}

	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
}
