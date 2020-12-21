package controllers

import (
	"gocommerce/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// loads templates from the templates folder
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// Index loads the product page with the proper data and proper template
func Index(w http.ResponseWriter, r *http.Request) {
	products := models.QueryAllProducts()

	tmpl.ExecuteTemplate(w, "index", products)
}

// Create loads the product creation page with the proper data and proper template
func Create(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "create", nil)
}

// Insert receives the data of the form, and if it's a POST, sends to models.CreateProduct
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceToFloat := convertToFloat("price", price)
		quantityToInt := convertToInt("quantity", quantity)

		models.CreateProduct(name, description, priceToFloat, quantityToInt)
	}

	code := 301

	http.Redirect(w, r, "/", code)
}

// Delete receives the ID of the product to delete, calls the model's method and redirects
func Delete(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	models.RemoveProduct(productID)

	http.Redirect(w, r, "/", 301)
}

// Edit receives the ID of the product to update, and applies info into the edit page
func Edit(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	product := models.QueryProduct(productID)

	tmpl.ExecuteTemplate(w, "edit", product)
}

// Update receives the ID of the product to update, calls the model's method and redirects
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idToInt := convertToInt("id", id)
		priceToFloat := convertToFloat("price", price)
		quantityToInt := convertToInt("quantity", quantity)

		models.UpdateProduct(idToInt, name, description, priceToFloat, quantityToInt)
	}

	http.Redirect(w, r, "/", 301)
}

func convertToFloat(field string, value string) float64 {
	valueToFloat, err := strconv.ParseFloat(value, 64)

	if err != nil {
		log.Println("Error while converting", field, "to float:", err)
	}

	return valueToFloat
}

func convertToInt(field string, value string) int {
	valueToInt, err := strconv.Atoi(value)

	if err != nil {
		log.Println("Error while converting", field, "to integer:", err)
	}

	return valueToInt
}
