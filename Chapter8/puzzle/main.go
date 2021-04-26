package main

import (
	"fmt"
	"study-head-first-go/Chapter8/geo"
)

func main() {
	location := geo.Coordinate{
		Latitude:  37.42,
		Longitude: -122.08,
	}

	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
