package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name      string               `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email     string               `json:"email" form:"email" gorm:"type: varchar(255)"`
	Phone     string               `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	PosCode   string               `json:"pos_code" form:"pos_code" gorm:"type: varchar(255)"`
	Address   string               `json:"address" form:"address" gorm:"type : varchar(255)"`
	Subtotal  int                  `json:"price"`
	Status    string               `json:"status"  gorm:"type:varchar(25)"`
	AccountID int                  `json:"accountid"`
	Account   UsersProfileResponse `json:"account"`
	OrderID   []int                `json:"orderid" gorm:"-"`
	Order     []Order              `json:"order" gorm:"-"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

// type TransactionResponse struct {
// 	ID     int `json:"id"`
// 	UserID int `json:"user_id"`
// }

// func (TransactionResponse) TableName() string {
// 	return "transactions"
// }
