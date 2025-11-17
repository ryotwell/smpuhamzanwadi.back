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
	GetAll(limit, page int) ([]model.Batch, error)
	GetActiveBatch() (*model.Batch, error)
	GetByID(id int) (*model.Batch, error)
	Update(id int, batch *model.Batch) error
	Delete(id int) error
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
	err := r.db.Where("year = ?", batch.Year).First(&batch).Error
	if err == nil {
		return errors.New("Batch already exist")
	}

	if err == gorm.ErrRecordNotFound {
		if err := r.db.Create(batch).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *batchRepository) GetAll(limit, page int) ([]model.Batch, error) {
	var batches []model.Batch

	offset := (page - 1) * limit

	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&batches).Error

	return batches, err
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

func (r *batchRepository) GetByID(id int) (*model.Batch, error) {
	var batch model.Batch
	err := r.db.
		Where("id = ?", id).
		First(&batch).Error

	if err != nil {
		return nil, err
	}

	return &batch, nil
}

func (r *batchRepository) GetActiveBatch() (*model.Batch, error) {
	var batch model.Batch

	err := r.db.
		Where("is_active = ?", true).
		First(&batch).Error

	if err != nil {
		return nil, err
	}

	return &batch, nil
}

func (r *batchRepository) Update(id int, batch *model.Batch) error {
	return r.db.Model(&model.Batch{}).
		Where("id = ?", id).
		Updates(batch).
		Error
}

func (r *batchRepository) Delete(id int) error {
	return r.db.Delete(&model.Batch{}, id).Error
}
