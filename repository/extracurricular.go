package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type ExtracurricularRepository interface {
	Create(ex *model.Extracurricular) error
	Update(id int, ex *model.Extracurricular) error
	Delete(id int) error
	GetByID(id int) (*model.Extracurricular, error)
	GetAll(limit, offset int) ([]model.Extracurricular, error)
}

type extracurricularRepository struct {
	db *gorm.DB
}

func NewExtracurricularRepository(db *gorm.DB) ExtracurricularRepository {
	return &extracurricularRepository{db: db}
}

func (r *extracurricularRepository) Create(ex *model.Extracurricular) error {
	return r.db.Create(ex).Error
}

func (r *extracurricularRepository) Update(id int, ex *model.Extracurricular) error {
	return r.db.Model(&model.Extracurricular{}).
		Where("id = ?", id).
		Updates(ex).
		Error
}

func (r *extracurricularRepository) Delete(id int) error {
	return r.db.Delete(&model.Extracurricular{}, id).Error
}

func (r *extracurricularRepository) GetByID(id int) (*model.Extracurricular, error) {
	var ex model.Extracurricular
	err := r.db.Where("id = ?", id).First(&ex).Error
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

func (r *extracurricularRepository) GetAll(limit, page int) ([]model.Extracurricular, error) {
	var ex []model.Extracurricular

	offset := (page - 1) * limit

	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&ex).Error
	return ex, err
}
