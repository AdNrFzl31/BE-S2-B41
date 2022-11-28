package routes

import (
	"BE-S2-B41/handlers"
	"BE-S2-B41/pkg/middleware"
	"BE-S2-B41/pkg/mysql"
	"BE-S2-B41/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/order/{id}", middleware.Auth(h.AddOrder)).Methods("POST")
	r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/orders", h.FindOrders).Methods("GET")

	// r.HandleFunc("/order/{id}", h.GetOrder).Methods("GET")
	// // r.HandleFunc("/orders-id", h.FindOrdersByID).Methods("GET")
	// r.HandleFunc("/order", middleware.Auth(h.CreateOrder)).Methods("POST")
	// r.HandleFunc("/order/{id}", middleware.Auth(h.UpdateOrder)).Methods("PATCH")
	// r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
}
