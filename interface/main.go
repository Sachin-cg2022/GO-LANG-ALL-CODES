package main

import "fmt"

func main() {

	var obj Vehicle = Bike{Name: "honda",
		Colour: "red",
		Price:  100000,
	}

	obj.ShowDetails()
	fmt.Println(obj.ShowName())
}

// set of methods(behaviour)
type Vehicle interface {
	ShowDetails()
	ShowName() string
}

// set of variables(field)
type Bike struct {
	Name, Colour string
	Price        float64
}

func (bike Bike) ShowDetails() {
	fmt.Println("BikeName :", bike.Name)
	fmt.Println("Bikecolour :", bike.Colour)
	fmt.Println("Bikeprice :", bike.Price)
}

func (bike Bike) ShowName() string {
	return bike.Name
}
