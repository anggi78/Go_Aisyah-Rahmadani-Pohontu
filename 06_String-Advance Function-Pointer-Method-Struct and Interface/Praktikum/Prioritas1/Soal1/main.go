package main

import "fmt"

type Car struct {
	fuelin   float64
	carType  string
}

func (e Car) EstDistance() float64 {
	return e.fuelin / 1.5
}

func main() {
	estimatedDist := Car{10.5, "SUV"}.EstDistance()
	
	fmt.Println("Car type: SUV")
	fmt.Println("Est. distance:", estimatedDist, "mill")
}