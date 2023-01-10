package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/ogbenioye/assignments/hols_Challenge/handler/product"
	"github.com/ogbenioye/assignments/hols_Challenge/handler/store"
)

var NewBatch *store.Store

func NewLine(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(" ")
	}
}

func EOF() {
	var ans string
	NewLine(1)
	fmt.Println("Do you want to continue? Y/N")
	fmt.Scan(&ans)

	if strings.ToLower(ans) == "y" {
		Menu()
	} else if strings.ToLower(ans) == "n" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	} else {
		fmt.Println("ERROR: Invalid Input")
		EOF()
	}
}

func Menu() {
	for {
		var menuInput int
		NewLine(1)
		fmt.Println("Choose a menu number\n")
		fmt.Println("1 List all products\t\t2 Add a new product\n")
		fmt.Println("3 Check product status\t\t4 Sell an item\n")
		fmt.Println("5 List all sold items\t\t6 Check available products for sale\n")
		fmt.Println("7 Exit")

		_, err := fmt.Scan(&menuInput)
		if err != nil {
			fmt.Println("Invalid input")
		}

		switch menuInput {
		case 1:
			NewBatch.ListProducts()
			EOF()
		case 2:
			addNewProduct()
		case 3:
			productStats()
		case 4:
			sellprod()
		case 5:
			NewBatch.ListSoldItems()
			EOF()
		case 6:
			NewBatch.Up4Sale()
			EOF()
		case 7:
			fmt.Println("Goodbye")
			os.Exit(0)
		}

	}
}

func addNewProduct() {

	var supplierName string
	var productMake string
	var productModel string
	var transmission string
	var fuelType string
	var price float32
	var quantity int

	fmt.Println("Enter supplier name:")
	_, err := fmt.Scan(&supplierName)
	if err != nil {
		fmt.Printf("An error occured: %v", err)
	}

	fmt.Println("Enter product make name:")
	fmt.Scan(&productMake)

	fmt.Println("Enter product model name:")
	fmt.Scan(&productModel)

	fmt.Println("Enter product tansmission:")
	fmt.Scan(&transmission)

	fmt.Println("Enter fuel type:")
	fmt.Scan(&fuelType)

	fmt.Println("Enter product price:")
	fmt.Scan(&price)

	fmt.Println("Enter product quantity:")
	fmt.Scan(&quantity)

	prod := productMake + " " + productModel

	NewBatch = &store.Store{
		supplierName,
		product.Products{
			product.Car{
				productMake,
				productModel,
				transmission,
				fuelType,
				price,
				quantity,
			},
		},
	}

	store.StoreItems[prod] = NewBatch.Products
	NewLine(1)
	fmt.Println("**Product successfully added**")
	NewLine(1)

	EOF()
}

func productStats() {
	var mk string
	var md string
	NewLine(1)
	fmt.Println("Enter the product make to check status:")
	fmt.Scan(&mk)
	fmt.Println("Enter the product model to check status:")
	fmt.Scan(&md)

	mkMd := mk + " " + md

	v, ok := store.StoreItems[strings.ToLower(mkMd)]
	if ok {
		v.ProductStatus()
	} else {
		fmt.Println("Error: No prodct with that name")
		productStats()
	}
	EOF()
}

func sellprod() {
	var mk string
	var md string
	var quan int
	NewLine(1)
	fmt.Println("Enter the product make:")
	fmt.Scan(&mk)
	fmt.Println("Enter the product model:")
	fmt.Scan(&md)
	fmt.Println("Enter the quantity to sell")
	fmt.Scan(&quan)

	mkMd := mk + " " + md

	NewBatch.SellItem(mkMd, quan)

	NewLine(1)
	fmt.Println("*Product sold*")
	EOF()
}
