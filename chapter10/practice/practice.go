package main

import (
	"fmt"
	"log"
	"study-head-first-go/chapter10/geo"
)

func main() {
	coordinates := geo.Landmark{}
	err := coordinates.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Name())
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
