package handlers

import (
	orderdto "BE-S2-B41/dto/order"
	dto "BE-S2-B41/dto/result"
	"BE-S2-B41/models"
	"BE-S2-B41/repositories"
	"encoding/json"
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

func (h *handlerOrder) AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(orderdto.CreateOrder)
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
	buyyerID := int(userInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.OrderRepository.GetProductOrder(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Product Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	topings, err := h.OrderRepository.GetTopingOrder(request.TopingID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "Failed", Message: "Toping Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var priceTopings = 0
	for _, i := range topings {
		priceTopings += i.Price
	}
	var subTotal = request.Qty * (product.Price + priceTopings)

	dataOrder := models.Order{
		Qty:       request.Qty,
		BuyyerID:  buyyerID,
		ProductID: product.ID,
		Toping:    topings,
		Subtotal:  subTotal,
	}

	cart, err := h.OrderRepository.AddOrder(dataOrder)

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

// func (h *handlersOrder) GetOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	order, err := h.OrderRepository.GetOrder(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: " Server Error ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: convertResponseOrder(order)}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlersOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	// idTrans := int(userInfo["time"].(float64))
// 	request := new(orderdto.CreateOrder)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: " Server Error ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}
// 	validate := validator.New()
// 	err := validate.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: " Failed ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	// var toppingId []int
// 	// for _, r := range r.FormValue("toppingId") {
// 	// 	if int(r-'0') >= 0 {
// 	// 		toppingId = append(toppingId, int(r-'0'))
// 	// 	}
// 	// }
// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	requestForm := models.Order{
// 		ProductID: id,
// 		// TransactionID: idTrans,
// 		Qty:       request.QTY,
// 		Subtotal:  request.SubTotal,
// 		ToppingID: request.ToppingID,
// 	}
// 	validatee := validator.New()
// 	errr := validatee.Struct(requestForm)
// 	if errr != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: " Failed ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	// topping, _ := h.OrderRepository.FindToppingsID(request.ToppingID)
// 	order := models.Order{
// 		ProductID: id,
// 		// TransactionID: idTrans,
// 		Qty:      request.QTY,
// 		Subtotal: request.SubTotal,
// 		// Topping:       topping,
// 	}
// 	data, err := h.OrderRepository.CreateOrder(order)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: " Server Error ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handlersOrder) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	request := new(orderdto.UpdateOrder)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: " Failed ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	order, err := h.OrderRepository.GetOrder(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Status: " Failed ", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}
// 	// len > 0
// 	if request.ProductID != 0 {
// 		order.ProductID = request.ProductID
// 	}
// 	data, err := h.OrderRepository.UpdateOrder(order)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Status: "Server Error", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

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
