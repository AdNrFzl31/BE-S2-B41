package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name      string               `json:"name" form:"name" gorm:"type: text"`
	Email     string               `json:"email" form:"email" gorm:"type: text"`
	Phone     string               `json:"phone" form:"phone" gorm:"type: text"`
	PosCode   string               `json:"pos_code" form:"pos_code" gorm:"type: text"`
	Address   string               `json:"address" form:"address" gorm:"type : text"`
	Subtotal  int                  `json:"price"`
	Status    string               `json:"status"  gorm:"type:varchar(25)"`
	AccountID int                  `json:"accountid"`
	Account   UsersProfileResponse `json:"account"`
	OrderID   []int                `json:"orderid" gorm:"-"`
	Order     []Order              `json:"order" gorm:"-"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}
type TransactionResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
