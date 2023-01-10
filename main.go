package main

import (
	"github.com/ogbenioye/assignments/hols_Challenge/handler"
	"github.com/ogbenioye/assignments/hols_Challenge/handler/product"
	"github.com/ogbenioye/assignments/hols_Challenge/handler/store"
)

func main() {

	go func() {
		firstBatch := &store.Store{
			SupplierName: "Joseph Eziwanne",
			Products: product.Products{
				Car: product.Car{
					"Toyota",
					"Corolla",
					"automatic",
					"petrol",
					8227.85,
					7,
				},
			},
		}

		secondBatch := &store.Store{
			"Dele Bishop",
			product.Products{
				product.Car{
					"Ford",
					"Explorer",
					"automatic",
					"petrol",
					15063.3,
					3,
				},
			},
		}

		thirdBatch := &store.Store{
			"Akanbi & JOhnsons ltd",
			product.Products{
				product.Car{
					"Honda",
					"CRV",
					"Automatic",
					"Petrol",
					6329.1,
					12,
				},
			},
		}

		// //*************** DEMO ***************

		// //adding items to store
		firstBatch.AddItem("Toyota Corolla")
		secondBatch.AddItem("Ford Explorer")
		thirdBatch.AddItem("Honda CRV")

	}()
	handler.Menu()
}
