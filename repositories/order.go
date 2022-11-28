package repositories

import (
	"BE-S2-B41/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	AddOrder(cart models.Order) (models.Order, error)
	GetOrder(ID int) (models.Order, error)
	FindOrder() ([]models.Order, error)
	DelOrder(cart models.Order) (models.Order, error)
	UpdateOrder(cart models.Order) (models.Order, error)
	GetProductOrder(ID int) (models.Product, error)
	GetTopingOrder(ID []int) ([]models.Toping, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddOrder(cart models.Order) (models.Order, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var cart models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").First(&cart, ID).Error

	return cart, err
}

func (r *repository) FindOrder() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").Find(&orders).Error // add this code

	return orders, err
}

func (r *repository) DelOrder(cart models.Order) (models.Order, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) UpdateOrder(cart models.Order) (models.Order, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) GetProductOrder(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repository) GetTopingOrder(ID []int) ([]models.Toping, error) {
	var toping []models.Toping
	err := r.db.Find(&toping, ID).Error

	return toping, err
}

// type OrderRepository interface {
// 	FindOrders() ([]models.Order, error)
// 	GetOrder(ID int) (models.Order, error)
// 	CreateOrder(Order models.Order) (models.Order, error)
// 	UpdateOrder(Order models.Order) (models.Order, error)
// 	DeleteOrder(Order models.Order) (models.Order, error)
// 	// CreateTransactionID(transaction models.Transaction) (models.Transaction, error)
// 	// FindToppingsID(ToppingID []int) ([]models.Topping, error)
// 	// FindOrdersTransaction(TrxID int) ([]models.Order, error)
// }

// func RepositoryOrder(db *gorm.DB) *repository {
// 	return &repository{db}
// }

// func (r *repository) FindOrders() ([]models.Order, error) {
// 	var order []models.Order
// 	err := r.db.Preload("Product").Preload("Topping").Find(&order).Error

// 	return order, err
// }

// func (r *repository) GetOrder(ID int) (models.Order, error) {
// 	var cart models.Order
// 	err := r.db.Preload("Product").Preload("Topping").First(&cart, ID).Error

// 	return cart, err
// }

// func (r *repository) CreateOrder(cart models.Order) (models.Order, error) {
// 	err := r.db.Create(&cart).Error

// 	return cart, err
// }

// func (r *repository) UpdateOrder(cart models.Order) (models.Order, error) {
// 	err := r.db.Save(&cart).Error

// 	return cart, err
// }

// func (r *repository) DeleteOrder(cart models.Order) (models.Order, error) {
// 	err := r.db.Delete(&cart).Error

// 	return cart, err
// }

// // func (r *repository) CreateTransactionID(transaction models.Transaction) (models.Transaction, error) {
// // 	err := r.db.Create(&transaction).Error

// // 	return transaction, err
// // }

// // func (r *repository) FindToppingsID(ToppingID []int) ([]models.Topping, error) {
// // 	var toppings []models.Topping
// // 	err := r.db.Debug().Find(&toppings, ToppingID).Error

// // 	return toppings, err
// // }

// // func (r *repository) FindTransactionID(TransactionID []int) ([]models.Topping, error) {
// // 	var toppings []models.Topping
// // 	err := r.db.Debug().Find(&toppings, TransactionID).Error

// // 	return toppings, err
// // }

// // func (r *repository) FindOrdersTransaction(TrxID int) ([]models.Order, error) {
// // 	var order []models.Order
// // 	err := r.db.Preload("Product").Preload("Topping").Debug().Find(&order, "transaction_id = ?", TrxID).Error

// // 	return order, err
// // }
