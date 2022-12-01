package models

import "time"

type Order struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int                  `json:"qty" gorm:"type:int"`
	Subtotal  int                  `json:"subtotal" gorm:"type: int"`
	ProductID int                  `json:"product_id"`
	Product   ProductResponse      `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Toping    []Toping             `json:"toppings" gorm:"many2many:order_toppings; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BuyyerID  int                  `json:"buyyer_id"`
	Buyyer    UsersProfileResponse `json:"buyyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type OrderResponse struct {
	Qty       int `json:"qty"`
	Subtotal  int `json:"subtotal"`
	BuyyerID  int `json:"buyyer_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
