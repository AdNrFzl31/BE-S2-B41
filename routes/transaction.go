package routes

import (
	"BE-S2-B41/handlers"
	"BE-S2-B41/pkg/middleware"
	"BE-S2-B41/pkg/mysql"
	"BE-S2-B41/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepoTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", middleware.Auth(h.Checkout)).Methods("POST")
	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/my-order", middleware.Auth(h.GetOrderByID)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/canceltrans/{id}", middleware.Auth(h.CancelTransaction)).Methods("PATCH")
	r.HandleFunc("/accepttrans/{id}", middleware.Auth(h.AcceptTransaction)).Methods("PATCH")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
