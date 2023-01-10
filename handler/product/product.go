package product

import (
	"fmt"
)

type Car struct {
	Make            string
	Model           string
	Transmission    string
	FuelType        string
	Price           float32
	QuantityInStock int
}

type Products struct {
	Car
}

func (p *Products) DisplayProduct() {
	fmt.Println(p)
}

func (p *Products) ProductStatus() {

	if p.Car.QuantityInStock > 0 {
		fmt.Printf("We have %v %v %v's available\n", p.Car.QuantityInStock, p.Car.Make, p.Car.Model)
	} else if p.Car.QuantityInStock == 0 {
		fmt.Printf("There are no %v %vs available\n", p.Car.Make, p.Car.Model)
	}
}
