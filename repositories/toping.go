package repositories

import (
	"BE-S2-B41/models"

	"gorm.io/gorm"
)

type TopingRepository interface {
	FindTopings() ([]models.Toping, error)
	GetToping(ID int) (models.Toping, error)
	CreateToping(Toping models.Toping) (models.Toping, error)
	UpdateToping(Toping models.Toping) (models.Toping, error)
	DeleteToping(Toping models.Toping) (models.Toping, error)
}

func RepositoryToping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTopings() ([]models.Toping, error) {
	var Topings []models.Toping
	err := r.db.Find(&Topings).Error

	return Topings, err
}

func (r *repository) GetToping(ID int) (models.Toping, error) {
	var Toping models.Toping
	err := r.db.First(&Toping, ID).Error

	return Toping, err
}

func (r *repository) CreateToping(Toping models.Toping) (models.Toping, error) {
	err := r.db.Create(&Toping).Error

	return Toping, err
}

func (r *repository) UpdateToping(Toping models.Toping) (models.Toping, error) {
	err := r.db.Save(&Toping).Error

	return Toping, err
}

func (r *repository) DeleteToping(Toping models.Toping) (models.Toping, error) {
	err := r.db.Delete(&Toping).Error

	return Toping, err
}
