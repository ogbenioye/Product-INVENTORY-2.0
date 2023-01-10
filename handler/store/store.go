package store

import (
	"fmt"
	"strings"

	"github.com/ogbenioye/assignments/hols_Challenge/handler/product"
)

var StoreItems = make(map[string]product.Products)

var soldItems = make(map[string]product.Products)

type Store struct {
	SupplierName string
	product.Products
}

func (s *Store) Up4Sale() {
	sum := 0
	carQuantity := 0
	motocycleQuantity := 0

	for _, v := range StoreItems {
		quantity := v.QuantityInStock
		carQuantity += v.Car.QuantityInStock
		sum += quantity
	}
	fmt.Printf("There are %v products available for sale. %v cars and %v motocycles\n", sum, carQuantity, motocycleQuantity)
}

func (s *Store) AddItem(productMakeAndModel string) {
	StoreItems[strings.ToLower(productMakeAndModel)] = s.Products
}

func (s *Store) ListProducts() {
	if len(StoreItems) != 0 {
		fmt.Println("Here's the list of all products in the store:")

		for _, v := range StoreItems {
			fmt.Println(v)
		}
	} else {
		fmt.Println("**There are no products in the store at the moment**")
	}

}

func (s *Store) SellItem(productMakeAndModel string, quantity int) {
	//use cli

	if quantity <= s.Products.QuantityInStock {
		newQuantity := s.Products.QuantityInStock - quantity
		StoreItems[productMakeAndModel] = product.Products{
			product.Car{
				s.Make,
				s.Model,
				s.Transmission,
				s.FuelType,
				s.Price,
				newQuantity,
			},
		}

		soldItems[productMakeAndModel] = product.Products{
			product.Car{
				s.Make,
				s.Model,
				s.Transmission,
				s.FuelType,
				s.Price,
				quantity,
			},
		}
	} else {
		fmt.Printf("Not enough products. Only %v %v %v's available\n", s.QuantityInStock, s.Make, s.Model)
	}
}

func (s *Store) ListSoldItems() {
	var totalPrice float32
	if len(soldItems) == 0 {
		fmt.Println("*No item has been sold*")
	} else {
		fmt.Println("Here's the list of all sold items:")

		for _, v := range soldItems {
			fmt.Println(v)
			price := v.Price * float32(v.QuantityInStock)
			totalPrice += price
		}
		fmt.Printf("The total price of these items is $%v\n", totalPrice)
	}

	// userconf.EOF()
}
