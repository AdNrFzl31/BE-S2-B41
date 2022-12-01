package routes

import (
	"BE-S2-B41/handlers"
	"BE-S2-B41/pkg/middleware"
	"BE-S2-B41/pkg/mysql"
	"BE-S2-B41/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	ProductRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(ProductRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")

}
