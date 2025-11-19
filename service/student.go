package service

import (
	"errors"
	"project_sdu/model"
	"project_sdu/repository"
	"time"
)

type StudentService interface {
	CreateStudent(student *model.Student) error
	GetStudentByID(id int) (*model.Student, error)
	GetStudentsByBatchYear(year, limit, page int, q string) ([]model.Student, error)
	GetAllStudents(limit int, page int, q string, batchID *int) ([]model.Student, error)
	UpdateStudent(id int, student *model.Student) error
	DeleteStudent(id int) error
}

type studentService struct {
	studentRepo repository.StudentRepository
	parentRepo  repository.ParentRepository
	batchRepo   repository.BatchRepository
}

func NewStudentService(studentRepo repository.StudentRepository, parentRepo repository.ParentRepository, batchRepo repository.BatchRepository) StudentService {
	return &studentService{
		studentRepo: studentRepo,
		parentRepo:  parentRepo,
		batchRepo:   batchRepo,
	}
}

// CreateStudent menyimpan student beserta parent-nya jika ada
func (s *studentService) CreateStudent(student *model.Student) error {
	var parentCreated bool
	if student.Parent != nil {
		if err := s.parentRepo.Create(student.Parent); err != nil {
			return err
		}
		parentCreated = true
		student.ParentId = &student.Parent.ID
	}

	// check if there is an active batch
	activeBatch, err := s.batchRepo.GetActiveBatch()
	if err != nil {
		if parentCreated && student.Parent != nil {
			_ = s.parentRepo.Delete(student.Parent.ID)
		}
		return errors.New("currently there is no batch active")
	}

	// validate the date
	now := time.Now()

	if activeBatch.StartDate == nil || activeBatch.EndDate == nil {
		return errors.New("batch has invalid start_date or end_date")
	}

	if now.Before(*activeBatch.StartDate) || now.After(*activeBatch.EndDate) {
		return errors.New("the registration period has ended")
	}

	student.BatchId = &activeBatch.ID
	student.Batch = nil

	if err := s.studentRepo.Create(student); err != nil {
		if parentCreated && student.Parent != nil {
			_ = s.parentRepo.Delete(student.Parent.ID)
		}
		return err
	}

	return nil
}

// func (s *studentService) CreateStudent(student *model.Student) error {
// 	var parentCreated bool
// 	var year int
// 	if student.Parent != nil {
// 		if err := s.parentRepo.Create(student.Parent); err != nil {
// 			return err
// 		}
// 		parentCreated = true
// 		student.ParentId = &student.Parent.ID
// 	}

// 	if student.Batch != nil {
// 		year = student.Batch.Year
// 	} else {
// 		year = time.Now().Year()
// 	}

// 	batch, err := s.batchRepo.GetOrCreateByYear(year)
// 	if err != nil {
// 		return err
// 	}

// 	student.BatchId = &batch.ID
// 	student.Batch = nil

// 	if err := s.studentRepo.Create(student); err != nil {
// 		if parentCreated && student.Parent != nil {
// 			_ = s.parentRepo.Delete(student.Parent.ID)
// 		}
// 		return err
// 	}

// 	return nil
// }

func (s *studentService) GetStudentsByBatchYear(year, limit, page int, q string) ([]model.Student, error) {
	batch, err := s.batchRepo.GetByYear(year)
	if err != nil {
		return []model.Student{}, err
	}

	return s.studentRepo.GetStudentsByBatchID(batch.ID, limit, page, q)
}

func (s *studentService) GetStudentByID(id int) (*model.Student, error) {
	student, err := s.studentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentService) GetAllStudents(limit int, page int, q string, batchID *int) ([]model.Student, error) {
	return s.studentRepo.GetAll(limit, page, q, batchID)
}

func (s *studentService) UpdateStudent(id int, student *model.Student) error {
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
