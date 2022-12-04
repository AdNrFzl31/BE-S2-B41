package handlers

import (
	dto "BE-S2-B41/dto/result"
	transactiondto "BE-S2-B41/dto/transaction"
	"BE-S2-B41/models"
	"BE-S2-B41/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey: os.Getenv("CLIENT_KEY"),
}

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

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// BuyyerID := int(userInfo["id"].(float64))

	// UserCart, err := h.TransactionRepository.GetOrderByUser(BuyyerID)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Status: "Failed", Message: "User Account not make a Purchase!"}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// var Total = 0
	// for _, i := range UserCart {
	// 	Total += i.Subtotal
	// }

	// dataTransaction := models.Transaction{
	// 	BuyyerID: BuyyerID,
	// 	Name:     request.Name,
	// 	Email:    request.Email,
	// 	Phone:    request.Phone,
	// 	Poscode:  request.Poscode,
	// 	Address:  request.Address,
	// 	Order:    UserCart,
	// 	Subtotal: Total,
	// 	Status:   "Waiting",
	// }

	// transaction, err := h.TransactionRepository.Checkout(dataTransaction)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Status: "Server Error", Message: "Transaction Failed!"}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// transactions, _ := h.TransactionRepository.GetTransaction(transaction.ID)

	// data := models.Transaction{
	// 	ID:       transactions.ID,
	// 	Name:     request.Name,
	// 	Email:    request.Email,
	// 	Phone:    request.Phone,
	// 	Poscode:  request.Poscode,
	// 	Address:  request.Address,
	// 	Order:    UserCart,
	// 	Subtotal: Total,
	// }

	// w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Status: "success", Data: data}
	// json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	IDtrans, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetTransaction(IDtrans)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.Update("pending", transaction.ID)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			h.TransactionRepository.Update("success", transaction.ID)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		h.TransactionRepository.Update("success", transaction.ID)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.Update("failed", transaction.ID)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.Update("failed", transaction.ID)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.Update("pending", transaction.ID)
	}

	w.WriteHeader(http.StatusOK)
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
		transaction.Poscode = request.Poscode
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
	orderUser, err := h.TransactionRepository.GetOrderByUser(data.BuyyerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataUpdate := models.Transaction{
		ID:       trans.ID,
		Name:     trans.Name,
		Address:  trans.Address,
		Status:   trans.Status,
		Order:    orderUser,
		Subtotal: trans.Subtotal,
		BuyyerID: trans.BuyyerID,
		Buyyer:   trans.Buyyer,
	}
	fmt.Println(dataUpdate)

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(trans.ID),
			GrossAmt: int64(trans.Subtotal),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: trans.Buyyer.Fullname,
			Email: trans.Buyyer.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: snapResp}
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

		Subtotal: trans.Subtotal,
		BuyyerID: trans.BuyyerID,
		Status:   trans.Status,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: dataTransactions}
	json.NewEncoder(w).Encode(response)
}
