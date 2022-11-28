package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	TopingRoutes(r)
	ProductRoutes(r)
	OrderRoutes(r)
	AuthRoutes(r)
	TransactionRoutes(r)
}
