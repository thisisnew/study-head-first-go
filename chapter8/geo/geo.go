package geo

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type Landmark struct {
	Name string
	Coordinate
}
