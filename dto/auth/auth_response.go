package authdto

type RegisterResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
}

type LoginResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
}
type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Role     string `gorm:"type: varchar(50)"  json:"role"`
}
