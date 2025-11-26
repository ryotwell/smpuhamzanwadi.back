package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type CurriculumRepository interface {
	Create(curriculum *model.Curriculum) error
	Update(id int, curriculum *model.Curriculum) error
	Delete(id int) error
	GetByID(id int) (*model.Curriculum, error)
	GetAll(limit, offset int) ([]model.Curriculum, error)
	GetByCategory(category string, limit, page int) ([]model.Curriculum, error)
}

type curriculumRepository struct {
	db *gorm.DB
}

func NewCurriculumRepository(db *gorm.DB) CurriculumRepository {
	return &curriculumRepository{db: db}
}

func (r *curriculumRepository) Create(curriculum *model.Curriculum) error {
	return r.db.Create(curriculum).Error
}

func (r *curriculumRepository) Update(id int, curriculum *model.Curriculum) error {
	return r.db.Model(&model.Curriculum{}).
		Where("id = ?", id).
		Updates(curriculum).
		Error
}

func (r *curriculumRepository) Delete(id int) error {
	return r.db.Delete(&model.Curriculum{}, id).Error
}

func (r *curriculumRepository) GetByID(id int) (*model.Curriculum, error) {
	var curriculum model.Curriculum
	err := r.db.Where("id = ?", id).First(&curriculum).Error
	if err != nil {
		return nil, err
	}
	return &curriculum, nil
}

func (r *curriculumRepository) GetAll(limit, page int) ([]model.Curriculum, error) {
	var curriculum []model.Curriculum

	offset := (page - 1) * limit

	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&curriculum).Error
	return curriculum, err
}

func (r *curriculumRepository) GetByCategory(category string, limit, page int) ([]model.Curriculum, error) {
	var curriculum []model.Curriculum

	offset := (page - 1) * limit

	err := r.db.
		Where("category ILIKE ?", category).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&curriculum).Error

	return curriculum, err
}