package service

import (
	"errors"
	"project_sdu/model"
	"project_sdu/repository"
)

type StudentService interface {
	CreateStudent(student *model.Student) error
	GetStudentByID(id int) (*model.Student, error)
	GetAllStudents(limit int, offset int) ([]model.Student, error)
	UpdateStudent(id int, student *model.Student) error
	DeleteStudent(id int) error
}

type studentService struct {
	studentRepo repository.StudentRepository
	parentRepo  repository.ParentRepository
}

func NewStudentService(studentRepo repository.StudentRepository, parentRepo repository.ParentRepository) StudentService {
	return &studentService{
		studentRepo : studentRepo,
		parentRepo : parentRepo,
	}
}

// CreateStudent menyimpan student beserta parent-nya jika ada
func (s *studentService) CreateStudent(student *model.Student) error {
	// Simpan parent jika ada
	if student.Parent != nil {
		if err := s.parentRepo.Create(student.Parent); err != nil {
			return errors.New("failed to create parent: " + err.Error())
		}
		// Set ParentID di student setelah parent tersimpan
		student.ParentId = &student.Parent.ID
	}

	// Simpan student
	if err := s.studentRepo.Create(student); err != nil {
		return errors.New("failed to create student: " + err.Error())
	}

	return nil
}

func (s *studentService) GetStudentByID(id int) (*model.Student, error) {
	student, err := s.studentRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("student not found")
	}
	return student, nil
}

func (s *studentService) GetAllStudents(limit int, offset int) ([]model.Student, error) {
	students, err := s.studentRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentService) UpdateStudent(id int, student *model.Student) error {
	// Update parent first if exists
	// if student.Parent != nil {
	// 	if student.Parent.ID != 0 {
	// 		// Update existing parent
	// 		if err := s.parentRepo.Update(student.Parent); err != nil {
	// 			return errors.New("failed to update parent: " + err.Error())
	// 		}
	// 	} else {
	// 		// Create parent if new
	// 		if err := s.parentRepo.Create(student.Parent); err != nil {
	// 			return errors.New("failed to create parent: " + err.Error())
	// 		}
	// 		student.ParentId = &student.Parent.ID
	// 	}
	// }

	// Update student
	if err := s.studentRepo.Update(id, student); err != nil {
		return err
	}

	return nil
}

func (s *studentService) DeleteStudent(id int) error {
	if err := s.studentRepo.Delete(id); err != nil {
		return err
	}
	return nil
}
