package routes

// import (
// 	"BE-S2-B41/handlers"
// 	"BE-S2-B41/pkg/mysql"
// 	"BE-S2-B41/repositories"

// 	"github.com/gorilla/mux"
// )

// func ProfileRoutes(r *mux.Router) {
// 	profileRepository := repositories.RepositoryProfile(mysql.DB)
// 	h := handlers.HandlerProfile(profileRepository)

// 	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
// }
