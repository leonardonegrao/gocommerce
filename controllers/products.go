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

		priceToFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Error while converting price to float:", err)
		}

		quantityToInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Error while converting quantity to integer:", err)
		}

		models.CreateProduct(name, description, priceToFloat, quantityToInt)
	}

	code := 301

	http.Redirect(w, r, "/", code)
}
