package orderdto

type CreateOrder struct {
	BuyyerID  int   `json:"buyyer_id" form:"buyyer_id"`
	Qty       int   `json:"qty" form:"qty"`
	SubTotal  int   `json:"subtotal" form:"subtotal"`
	ProductID int   `json:"product_id" form:"product_id"`
	TopingID  []int `json:"toping_id" form:"toping_id"`
}

type UpdateOrder struct {
	Qty int `json:"qty" form:"qty"`
}
