package repositories

import (
	"BE-S2-B41/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	Checkout(transaction models.Transaction) (models.Transaction, error)
	GetOrderByUser(ID int) ([]models.Order, error)
	GetTransaction(ID int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	CancelTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *repository) Checkout(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) GetOrderByUser(ID int) ([]models.Order, error) {
	var cart []models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").Where("buyyer_id = ?", ID).Find(&cart).Error
	return cart, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").Preload("Buyyer").First(&transaction, ID).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) CancelTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}
