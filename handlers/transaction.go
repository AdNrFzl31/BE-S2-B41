package handlers

import (
	dto "BE-S2-B41/dto/result"
	transactiondto "BE-S2-B41/dto/transaction"
	"BE-S2-B41/models"
	"BE-S2-B41/repositories"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: transactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(transactiondto.Checkout)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "cek dto"}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "error validation"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	accountID := int(userInfo["id"].(float64))

	UserCart, err := h.TransactionRepository.GetOrderByUser(accountID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "User Account not make a Purchase!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var Total = 0
	for _, i := range UserCart {
		Total += i.Subtotal
	}

	dataTransaction := models.Transaction{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		PosCode:  request.Poscode,
		Address:  request.Address,
		Status:   "Waiting",
		Subtotal: Total,
	}

	transaction, err := h.TransactionRepository.Checkout(dataTransaction)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: "Transaction Failed!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactions, _ := h.TransactionRepository.GetTransaction(transaction.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: transactions}
	json.NewEncoder(w).Encode(response)
}
