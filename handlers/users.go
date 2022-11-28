package handlers

import (
	dto "BE-S2-B41/dto/result"
	usersdto "BE-S2-B41/dto/users"
	"BE-S2-B41/models"
	"BE-S2-B41/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: user}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	request := new(usersdto.CreateUserRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	user := models.User{
// 		Fullname: request.Fullname,
// 		Email:    request.Email,
// 		Password: request.Password,
// 	}
// 	data, err := h.UserRepository.CreateUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	request := new(usersdto.UpdateUserRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	user, err := h.UserRepository.GetUser(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	if request.Fullname != "" {
// 		user.Fullname = request.Fullname
// 	}
// 	if request.Email != "" {
// 		user.Email = request.Email
// 	}
// 	if request.Password != "" {
// 		user.Password = request.Password
// 	}
// 	data, err := h.UserRepository.UpdateUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "you're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Image:    u.Image,
	}
}
