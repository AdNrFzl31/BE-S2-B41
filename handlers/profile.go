package handlers

// import (
// 	profiledto "BE-S2-B41/dto/profile"
// 	dto "BE-S2-B41/dto/result"
// 	"BE-S2-B41/models"
// 	"BE-S2-B41/repositories"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type handlerProfile struct {
// 	ProfileRepository repositories.ProfileRepository
// }

// func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
// 	return &handlerProfile{ProfileRepository}
// }

// func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var profile models.Profile
// 	profile, err := h.ProfileRepository.GetProfile(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponseProfile(profile)}
// 	json.NewEncoder(w).Encode(response)
// }

// func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
// 	return profiledto.ProfileResponse{
// 		ID:      u.ID,
// 		Phone:   u.Phone,
// 		Gender:  u.Gender,
// 		Address: u.Address,
// 		UserID:  u.UserID,
// 		User:    u.User,
// 	}
// }