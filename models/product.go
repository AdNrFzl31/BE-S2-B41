package models

import "time"

type Product struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Nameproduct string    `json:"nameproduct" gorm:"type: varchar(255)"`
	Price       int       `json:"price" gorm:"type: int"`
	Image       string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Nameproduct string `json:"nameproduct"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}

func (ProductResponse) TableName() string {
	return "products"
}
