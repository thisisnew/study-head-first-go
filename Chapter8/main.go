package main

import (
	"fmt"
	magazine "study-head-first-go/Chapter8/magazine"
)

type myStruct struct {
	myField int
}

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active:", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}
