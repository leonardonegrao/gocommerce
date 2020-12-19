package models

import (
	"gocommerce/database"
)

// Product is the model of the products that are received from the database and displayed
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

// QueryAllProducts runs database query to fetch all products saved
func QueryAllProducts() []Product {
	db := database.ConnectWithDatabase()
	defer db.Close()
	selectProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	dbProduct := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		dbProduct.ID = id
		dbProduct.Name = name
		dbProduct.Description = description
		dbProduct.Price = price
		dbProduct.Quantity = quantity

		products = append(products, dbProduct)
	}

	return products
}

// CreateProduct creates a product in the database
func CreateProduct(name, description string, price float64, quantity int) {
	db := database.ConnectWithDatabase()
	defer db.Close()

	insertData, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
}
