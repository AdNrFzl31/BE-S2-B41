package handlers

import (
	dto "BE-S2-B41/dto/result"
	transactiondto "BE-S2-B41/dto/transaction"
	"BE-S2-B41/models"
	"BE-S2-B41/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
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
	AccountID := int(userInfo["id"].(float64))

	UserCart, err := h.TransactionRepository.GetOrderByUser(AccountID)
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
		AccountID: AccountID,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		PosCode:   request.Poscode,
		Address:   request.Address,
		Order:     UserCart,
		Subtotal:  Total,
		Status:    "Waiting",
	}

	transaction, err := h.TransactionRepository.Checkout(dataTransaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: "Transaction Failed!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactions, _ := h.TransactionRepository.GetTransaction(transaction.ID)

	data := models.Transaction{
		ID:       transactions.ID,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		PosCode:  request.Poscode,
		Address:  request.Address,
		Order:    UserCart,
		Subtotal: Total,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CancelTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction.Status = "Cancel"
	data, err := h.TransactionRepository.CancelTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) AcceptTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Cek id Transaction => " + err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// fmt.Println(transaction.ID)
	// fmt.Println(transaction.Status)

	transaction.Status = "Success"
	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(transactiondto.UpdateTransaction)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransaction(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	orderTrans, err := h.TransactionRepository.GetOrderByUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var Total = 0
	for _, i := range orderTrans {
		Total += i.Subtotal
	}

	// fmt.Println(orderTrans)
	// fmt.Println(transaction.ID)
	// fmt.Println(Total)

	if request.Name != "" {
		transaction.Name = request.Name
	}
	if request.Email != "" {
		transaction.Email = request.Email
	}
	if request.Phone != "" {
		transaction.Phone = request.Phone
	}
	if request.Poscode != "" {
		transaction.PosCode = request.Poscode
	}
	if request.Address != "" {
		transaction.Address = request.Address
	}
	transaction.Status = "Payment"
	transaction.Subtotal = Total

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trans, _ := h.TransactionRepository.GetTransaction(data.ID)
	orderUser, err := h.TransactionRepository.GetOrderByUser(data.AccountID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataUpdate := models.Transaction{
		ID:        trans.ID,
		Name:      trans.Name,
		Address:   trans.Address,
		Status:    trans.Status,
		Order:     orderUser,
		Subtotal:  trans.Subtotal,
		AccountID: trans.AccountID,
		Account:   trans.Account,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: dataUpdate}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.GetOrderByID()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: transactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	orders, err := h.TransactionRepository.FindTransactionID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: orders}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trans, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// order, err := h.TransactionRepository.GetOrderByID(trans.UserID)

	dataTransactions := models.Transaction{
		ID:      trans.ID,
		Name:    trans.Name,
		Email:   trans.Email,
		Phone:   trans.Phone,
		Address: trans.Address,

		Subtotal:  trans.Subtotal,
		AccountID: trans.AccountID,
		Status:    trans.Status,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: dataTransactions}
	json.NewEncoder(w).Encode(response)
}
