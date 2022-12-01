package topingdto

type AddToping struct {
	Nametoping string `json:"nametoping" gorm:"type:text"`
	Price      int    `json:"price" gorm:"type:int"`
	Image      string `json:"image" gorm:"type:text"`
}
type UpdateToping struct {
	Nametoping string `json:"nametoping" gorm:"type:text"`
	Price      int    `json:"price" gorm:"type:int"`
	Image      string `json:"image" gorm:"type:text"`
}
