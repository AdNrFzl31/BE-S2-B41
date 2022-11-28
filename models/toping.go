package models

import "time"

type Toping struct {
	ID         int       `json:"id" gorm:"primary_key:auto_increment"`
	Nametoping string    `json:"nametoping" gorm:"type: varchar(255)"`
	Price      int       `json:"price" gorm:"type: int"`
	Image      string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type TopingResponse struct {
	ID         int    `json:"id"`
	Nametoping string `json:"nametoping"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
}

type TopingOrderResponse struct {
	ID         int    `json:"id"`
	Nametoping string `json:"nametoping"`
}

func (TopingResponse) TableName() string {
	return "topings"
}

func (TopingOrderResponse) TableName() string {
	return "topings"
}
