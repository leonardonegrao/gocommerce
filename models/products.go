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

// QueryProduct runs database query to fetch product with specified ID
func QueryProduct(productID string) Product {
	db := database.ConnectWithDatabase()
	defer db.Close()

	productQuery, err := db.Query("select * from products where id = $1", productID)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := productQuery.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	return product
}

// QueryAllProducts runs database query to fetch all products saved
func QueryAllProducts() []Product {
	db := database.ConnectWithDatabase()
	defer db.Close()
	selectProductsQuery, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	dbProduct := Product{}
	products := []Product{}

	for selectProductsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProductsQuery.Scan(&id, &name, &description, &price, &quantity)

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

	insertDataQuery, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataQuery.Exec(name, description, price, quantity)
}

// RemoveProduct removes a product from the database
func RemoveProduct(productID string) {
	db := database.ConnectWithDatabase()
	defer db.Close()

	removeProductQuery, err := db.Prepare("delete from products where id = $1")

	if err != nil {
		panic(err.Error())
	}

	removeProductQuery.Exec(productID)
}

// UpdateProduct updates a product in the database
func UpdateProduct(productID int, name, description string, price float64, quantity int) {
	db := database.ConnectWithDatabase()
	defer db.Close()

	updateProductQuery, err := db.Prepare("update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5")

	if err != nil {
		panic(err.Error())
	}

	updateProductQuery.Exec(name, description, price, quantity, productID)
}
