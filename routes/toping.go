package routes

import (
	"BE-S2-B41/handlers"
	"BE-S2-B41/pkg/middleware"
	"BE-S2-B41/pkg/mysql"
	"BE-S2-B41/repositories"

	"github.com/gorilla/mux"
)

func TopingRoutes(r *mux.Router) {
	TopingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(TopingRepository)

	r.HandleFunc("/topings", h.FindTopings).Methods("GET") // add this code
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")
	r.HandleFunc("/toping", middleware.Auth(middleware.UploadFile(h.CreateToping))).Methods("POST")       // add this code
	r.HandleFunc("/toping/{id}", middleware.Auth(h.DeleteToping)).Methods("DELETE")                       // add this code
	r.HandleFunc("/toping/{id}", middleware.Auth(middleware.UploadFile(h.UpdateToping))).Methods("PATCH") // add this code

}
