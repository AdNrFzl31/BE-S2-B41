package repositories

import (
	"BE-S2-B41/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	Checkout(transaction models.Transaction) (models.Transaction, error)
	FindTransactions(ID int) ([]models.Transaction, error)
	CancelTransaction(transaction models.Transaction) (models.Transaction, error)

	GetOrderByUser(ID int) ([]models.Order, error)
	GetTransaction(ID int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetOrderByID() ([]models.Transaction, error)
	FindTransactionID(ID int) ([]models.Transaction, error)
	Update(status string, ID int) error
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Checkout(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) FindTransactions(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").Where("BuyyerID = ?", ID).Find(&transactions).Error
	return transactions, err
}

func (r *repository) CancelTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}

func (r *repository) GetOrderByUser(ID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Preload("Toping").Where("transaction_id = ?", ID).Find(&order).Error
	return order, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").Preload("Order").First(&transaction, ID).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) GetOrderByID() ([]models.Transaction, error) {
	var order []models.Transaction
	err := r.db.Preload("Order").Preload("Buyyer").Where("status != ?", "waiting").Find(&order).Error
	return order, err
}

func (r *repository) FindTransactionID(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Buyyer").Preload(clause.Associations).Preload("Order.Product").Preload("Order.Toping").Where("status != ? AND buyyer_id = ?", "waiting", ID).Find(&transaction).Error

	return transaction, err
}
func (r *repository) Update(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("Order.Product").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	// if status != transaction.Status && status == "success" {
	// 	var product models.Product
	// 	r.db.First(&product, transaction.Product.ID)
	// 	product.Qty = product.Qty - 1
	// 	r.db.Save(&product)
	// }

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}
