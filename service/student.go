package service

import (
	"project_sdu/model"
	"project_sdu/repository"
	"time"
)

type StudentService interface {
	CreateStudent(student *model.Student) error
	GetStudentByID(id int) (*model.Student, error)
	GetStudentsByBatchYear(year, limit, page int, q string) ([]model.Student, error)
	GetAllStudents(limit int, page int, q string) ([]model.Student, error)
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
	var year int
	if student.Parent != nil {
		if err := s.parentRepo.Create(student.Parent); err != nil {
			return err
		}
		parentCreated = true
		student.ParentId = &student.Parent.ID
	}

	if student.Batch != nil {
		year = student.Batch.Year
	} else {
		year = time.Now().Year()
	}

	batch, err := s.batchRepo.GetOrCreateByYear(year)
	if err != nil {
		return err
	}

	student.BatchId = &batch.ID
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

// 	// save parent if exist
// 	if student.Parent != nil {
// 		if err := s.parentRepo.Create(student.Parent); err != nil {
// 			return err
// 		}
// 		parentCreated = true
// 		student.ParentId = &student.Parent.ID
// 	}

// 	if student.Batch != nil {
// 		year := student.Batch.Year

// 		// cari batch berdasarkan tahun
// 		batch, err := s.batchRepo.GetByYear(year)
// 		if err != nil {
// 			return err
// 		}

// 		// jika tidak ada → buat baru
// 		if batch == nil {
// 			newBatch := model.Batch{
// 				Name: fmt.Sprintf("Batch %d", year),
// 				Year: year,
// 			}
// 			if err := s.batchRepo.Create(&newBatch); err != nil {
// 				return err
// 			}
// 			batch = &newBatch
// 		}

// 		student.BatchId = &batch.ID
// 		student.Batch = nil

// 		// simpan student
// 		if err := s.studentRepo.Create(student); err != nil {
// 			if parentCreated && student.Parent != nil {
// 				_ = s.parentRepo.Delete(student.Parent.ID)
// 			}
// 			return err
// 		}

// 		return nil
// 	}

// 	currentYear := time.Now().Year()

// 	batch, err := s.batchRepo.GetByYear(currentYear)
// 	if err != nil {
// 		// Jika batch tidak ditemukan -> buat baru
// 		newBatch := model.Batch{
// 			Name: fmt.Sprintf("Batch %d", currentYear),
// 			Year: currentYear,
// 		}

// 		if err := s.batchRepo.Create(&newBatch); err != nil {
// 			return err
// 		}

// 		batch = &newBatch
// 	}

// 	student.BatchId = &batch.ID

// 	if err := s.studentRepo.Create(student); err != nil {
// 		// delete parent if fail to save student
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

func (s *studentService) GetAllStudents(limit int, page int, q string) ([]model.Student, error) {
	students, err := s.studentRepo.GetAll(limit, page, q)
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
