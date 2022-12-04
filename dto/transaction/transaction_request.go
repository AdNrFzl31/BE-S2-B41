package transactiondto

type Checkout struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Poscode  string `json:"poscode" form:"poscode"`
	Address  string `json:"address" form:"address"`
	OrderID  []int  `json:"order_id" form:"order_id"`
	Subtotal int    `json:"subtotal" form:"subtotal"`
	Status   string `json:"status" form:"status"`
	BuyyerID int    `json:"buyyer_id" form:"buyyer_id"`
}

type UpdateTransaction struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Poscode string `json:"poscode" form:"poscode"`
	Address string `json:"address" form:"address"`
	// Status  string `json:"status" form:"status"`
}
