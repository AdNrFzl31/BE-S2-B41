package authdto

type RegisterResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
}

type LoginResponse struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
