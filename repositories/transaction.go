package repositories

import (
	"BE-S2-B41/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	Checkout(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetOrderByUser(ID int) ([]models.Order, error)
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").First(&transaction, ID).Error
	return transaction, err
}

func (r *repository) Checkout(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Order").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetOrderByUser(ID int) ([]models.Order, error) {
	var cart []models.Order
	err := r.db.Preload("Buyyer").Where("buyyer_id = ?", ID).Find(&cart).Error

	return cart, err
}
