package repository

import (
	"errors"
	"fmt"
	"project_sdu/model"

	"gorm.io/gorm"
)

type BatchRepository interface {
	GetByYear(year int) (*model.Batch, error)
	Create(batch *model.Batch) error
	GetOrCreateByYear(year int) (*model.Batch, error)
}

type batchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) BatchRepository {
	return &batchRepository{db}
}

func (r *batchRepository) GetByYear(year int) (*model.Batch, error) {
	var batch model.Batch

	err := r.db.Where("year = ?", year).First(&batch).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &batch, nil
}

func (r *batchRepository) Create(batch *model.Batch) error {
	return r.db.Create(batch).Error
}

func (r *batchRepository) GetOrCreateByYear(year int) (*model.Batch, error) {
	batch := model.Batch{}
	err := r.db.Where("year = ?", year).First(&batch).Error

	if err == gorm.ErrRecordNotFound {
		name := fmt.Sprintf("Tahun Ajaran %d/%d", year, year+1)
		newBatch := model.Batch{
			Name: name,
			Year: year,
		}
		if err := r.db.Create(&newBatch).Error; err != nil {
			return nil, err
		}
		return &newBatch, nil
	}
	if err != nil {
		return nil, err
	}
	return &batch, nil
}
