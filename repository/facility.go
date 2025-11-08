package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type FacilityRepository interface {
	Create(facility *model.Facility) error
	Update(id int, facility *model.Facility) error
	Delete(id int) error
	GetByID(id int) (*model.Facility, error)
	GetAll(limit, offset int) ([]model.Facility, error)
}

type facilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) FacilityRepository {
	return &facilityRepository{db: db}
}

func (r *facilityRepository) Create(facility *model.Facility) error {
	return r.db.Create(facility).Error
}

func (r *facilityRepository) Update(id int, facility *model.Facility) error {
	return r.db.Model(&model.Facility{}).
		Where("id = ?", id).
		Updates(facility).
		Error
}

func (r *facilityRepository) Delete(id int) error {
	return r.db.Delete(&model.Facility{}, id).Error
}

func (r *facilityRepository) GetByID(id int) (*model.Facility, error) {
	var facility model.Facility
	err := r.db.Where("id = ?", id).First(&facility).Error
	if err != nil {
		return nil, err
	}
	return &facility, nil
}

func (r *facilityRepository) GetAll(limit, page int) ([]model.Facility, error) {
	var facilities []model.Facility

	offset := (page - 1) * limit

	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&facilities).Error
	return facilities, err
}
