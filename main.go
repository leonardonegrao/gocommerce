package main

import (
	"gocommerce/routes"
	"net/http"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
