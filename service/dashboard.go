package service

import (
	"project_sdu/model"
	"project_sdu/repository"
)

type DashboardService interface {
	GetTotalStudents() (int, error)
	GetTotalPosts() (int, error)
	GetTotalStudentsActiveBatch() (int, error)
	GetTotalBatch() (int, error)
	GetActiveBatch() (*model.Batch, error)
}

type dashboardService struct {
	studentRepo repository.StudentRepository
	postRepo    repository.PostRepository
	batchRepo   repository.BatchRepository
}

func NewDashboardService(
	studentRepo repository.StudentRepository,
	postRepo repository.PostRepository,
	batchRepo repository.BatchRepository,
) DashboardService {
	return &dashboardService{
		studentRepo: studentRepo,
		postRepo:    postRepo,
		batchRepo:   batchRepo,
	}
}

func (s *dashboardService) GetTotalStudents() (int, error) {
	total, err := s.studentRepo.CountAll()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *dashboardService) GetTotalPosts() (int, error) {
	total, err := s.postRepo.CountAll()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *dashboardService) GetTotalStudentsActiveBatch() (int, error) {
	batch, err := s.batchRepo.GetActiveBatch()
	if err != nil {
		return 0, err
	}

	if batch == nil {
		return 0, nil
	}

	total, err := s.studentRepo.CountByBatchID(batch.ID)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *dashboardService) GetTotalBatch() (int, error) {
	total, err := s.batchRepo.CountAll()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *dashboardService) GetActiveBatch() (*model.Batch, error) {
	batch, err := s.batchRepo.GetActiveBatch()
	if err != nil {
		return nil, err
	}
	return batch, nil
}
