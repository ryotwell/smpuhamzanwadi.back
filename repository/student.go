package repository

import (
	"errors"
	"project_sdu/model"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrNIKExists  = errors.New("NIK sudah terdaftar")
	ErrNISNExists = errors.New("NISN sudah terdaftar")
)

type StudentRepository interface {
	Create(student *model.Student) error
	GetByID(id int) (*model.Student, error)
	GetAll(limit int, offset int) ([]model.Student, error)
	Update(id int, student *model.Student) error
	Delete(id int) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

// Create a new student
func (r *studentRepository) Create(student *model.Student) error {
	err := r.db.Create(student).Error
	if err != nil {
		if strings.Contains(err.Error(), "idx_students_nik") {
			return ErrNIKExists
		}
		if strings.Contains(err.Error(), "idx_students_nisn") {
			return ErrNISNExists
		}
		return err
	}
	return nil
	// if err != nil {
	// 	// Tangkap error PostgreSQL
	// 	var pqErr *pq.Error
	// 	if errors.As(err, &pqErr) {
	// 		if pqErr.Code == "23505" { // unique_violation
	// 			if pqErr.Constraint == "idx_students_nik" {
	// 				return errors.New("NIK sudah terdaftar")
	// 			}
	// 			if pqErr.Constraint == "idx_students_nisn" {
	// 				return errors.New("NISN sudah terdaftar")
	// 			}
	// 		}
	// 	}
	// 	return err
	// }
	// return nil
}

// Get single student by ID with Parent relation
func (r *studentRepository) GetByID(id int) (*model.Student, error) {
	var student model.Student
	err := r.db.
		Preload("Parent").
		First(&student, id).
		Error

	if err != nil {
		return nil, err
	}

	return &student, nil
}

// Get all students with pagination support
// func (r *studentRepository) GetAll(limit int, offset int) ([]model.Student, error) {
// 	var students []model.Student
// 	err := r.db.
// 		Preload("Parent").
// 		Limit(limit).
// 		Offset(offset).
// 		Find(&students).
// 		Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return students, nil
// }

// Get all students with pagination support and sort the names in alphabetical order
func (r *studentRepository) GetAll(limit int, page int) ([]model.Student, error) {
	var students []model.Student

	offset := (page - 1) * limit

	err := r.db.
		Preload("Parent").
		Order("full_name ASC").
		Limit(limit).
		Offset(offset).
		Find(&students).
		Error

	if err != nil {
		return nil, err
	}

	return students, nil
}

// Update student and parent data by student ID
func (r *studentRepository) Update(id int, student *model.Student) error {
	var existingStudent model.Student
	if err := r.db.Preload("Parent").First(&existingStudent, id).Error; err != nil {
		return err
	}

	if err := r.db.Model(&existingStudent).Updates(student).Error; err != nil {
		return err
	}

	if student.Parent != nil {
		if err := r.db.Model(&existingStudent.Parent).Updates(student.Parent).Error; err != nil {
			return err
		}
	}

	return nil
}

// Delete student by ID
func (r *studentRepository) Delete(id int) error {
	return r.db.Delete(&model.Student{}, id).Error
}
