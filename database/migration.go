package database

import (
	"BE-S2-B41/models"
	"BE-S2-B41/pkg/mysql"
	"fmt"
)

// gunakan ini untuk otomatis membuat model ke database
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Toping{},
		&models.Product{},
		&models.Order{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration Success")
}
