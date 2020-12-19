package routes

import (
	"gocommerce/controllers"
	"net/http"
)

// GetRoutes provides the application routes
func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/insert", controllers.Insert)
}
