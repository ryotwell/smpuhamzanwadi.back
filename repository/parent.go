package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type ParentRepository interface {
	Create(parent *model.Parent) error
	GetAll(limit int, offset int) ([]model.Parent, error)
	GetByID(id int) (*model.Parent, error)
	Update(id int, parent *model.Parent) error
	Delete(id int) error
}

type parentRepository struct {
	db *gorm.DB
}

func NewParentRepo(db *gorm.DB) ParentRepository {
	return &parentRepository{db}
}

// Create a new parent
func (r *parentRepository) Create(parent *model.Parent) error {
	return r.db.Create(parent).Error
}

// Get all parent with pagination support and sort the names in alphabetical order
func (r *parentRepository) GetAll(limit int, offset int) ([]model.Parent, error) {
	var parents []model.Parent

	err := r.db.
		Limit(limit).
		Offset(offset).
		Find(&parents).
		Error

	if err != nil {
		return nil, err
	}

	return parents, nil
}

// Get parent by ID
func (r *parentRepository) GetByID(id int) (*model.Parent, error) {
	var parent model.Parent
	err := r.db.First(&parent, id).Error
	if err != nil {
		return nil, err
	}
	return &parent, nil
}

// Update parent data by ID
func (r *parentRepository) Update(id int, parent *model.Parent) error {
		if err := r.db.Model(&model.Parent{}).
		Where("id = ?", id).
		Updates(parent).
		Error; err != nil {
		return err
	}

	return nil
}

// Delete parent by ID
func (r *parentRepository) Delete(id int) error {
	return r.db.Delete(&model.Parent{}, id).Error
}
