package handlers

import (
	orderdto "BE-S2-B41/dto/order"
	dto "BE-S2-B41/dto/result"
	"BE-S2-B41/models"
	"BE-S2-B41/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// var path_file_order = os.Getenv("PATH_FILE")

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) FindOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.OrderRepository.FindOrder()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: " Server Error ", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: orders}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(orderdto.CreateOrder)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "cek dto"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println(request.BuyyerID)
	// fmt.Println(request.ProductID)
	// fmt.Println(request.Qty)
	// fmt.Println(request.SubTotal)
	// fmt.Println(request.TopingID)

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "error validation"}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	product, err := h.OrderRepository.GetProductOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Product Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println("products = ", product)

	topings, err := h.OrderRepository.GetTopingOrder(request.TopingID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Toping Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// fmt.Println("cek = ", topings)

	var priceTopings = 0
	for _, i := range topings {
		priceTopings += i.Price
	}
	var subTotal = request.Qty * (product.Price + priceTopings)
	fmt.Println("cek = ", subTotal)

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	buyerID := int(userInfo["id"].(float64))

	CekRequestTrans, _ := h.OrderRepository.GetTransactionID(buyerID)
	// fmt.Println("cek buyerid = ", CekRequestTrans)

	var transID int
	if CekRequestTrans.ID != 0 {
		transID = CekRequestTrans.ID
	} else {
		requestTrans := models.Transaction{
			Name:     "-",
			Email:    "-",
			Phone:    "-",
			Poscode:  "-",
			Address:  "-",
			Status:   "Waiting",
			Subtotal: 0,
			BuyyerID:  buyerID,
		}
		transOrder, _ := h.OrderRepository.RequestTransaction(requestTrans)
		transID = transOrder.ID
	}
	// fmt.Println("cek = ", transID)

	dataOrder := models.Order{
		Qty:           request.Qty,
		Subtotal:      subTotal,
		TransactionID: transID,
		ProductID:     product.ID,
		Toping:        topings,
		BuyyerID:      buyerID,
	}
	// fmt.Println("dataOrder = ", dataOrder)

	cart, err := h.OrderRepository.AddOrder(dataOrder)

	fmt.Println("Cart = ", cart.Qty)
	fmt.Println("Subtotal = ", cart.Subtotal)
	fmt.Println("TransactionID = ", cart.TransactionID)
	fmt.Println("ProductID = ", cart.ProductID)
	fmt.Println("Toping = ", cart.Toping)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: "Order Failed!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	order, _ := h.OrderRepository.GetOrder(cart.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: order}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) GetOrdersByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int64(userInfo["id"].(float64))

	transaction, _ := h.OrderRepository.GetTransactionID(int(userID))

	orders, err := h.OrderRepository.GetOrdersByID(int(transaction.ID))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// for i, p := range orders {
	// 	orders[i].Product.Image = os.Getenv("PATH_FILE") + p.Product.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: orders}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerOrder) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(orderdto.UpdateOrder)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	order, err := h.OrderRepository.GetOrder(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	priceItem := order.Subtotal / order.Qty
	order.Qty = request.Qty
	order.Subtotal = priceItem * order.Qty

	data, err := h.OrderRepository.UpdateOrder(order)
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

func (h *handlerOrder) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: " Failed ", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data, err := h.OrderRepository.DelOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponseOrder(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseOrder(u models.Order) models.Order {
	return models.Order{
		ID:       u.ID,
		Qty:      u.Qty,
		Subtotal: u.Subtotal,
		Product:  u.Product,
		Toping:   u.Toping,
	}
}
