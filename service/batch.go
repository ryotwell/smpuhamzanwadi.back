package service

import (
	"errors"
	"project_sdu/model"
	"project_sdu/repository"
)

type BatchService interface {
	Create(batch *model.Batch) error
	Update(id int, batch *model.Batch) error
	Delete(id int) error
	GetByID(id int) (*model.Batch, error)
	GetAll(limit, page int, q string) ([]model.Batch, error)
}

type batchService struct {
	batchRepo repository.BatchRepository
}

func NewBatchService(batchRepo repository.BatchRepository) BatchService {
	return &batchService{batchRepo}
}

func (s *batchService) Create(batch *model.Batch) error {
	if err := s.batchRepo.Create(batch); err != nil {
		return err
	}
	return nil
}

func (s *batchService) Update(id int, batch *model.Batch) error {
	batchExist, _ := s.batchRepo.GetActiveBatch()

	if batchExist != nil && batch.IsActive != nil && *batch.IsActive {
		return errors.New("there is already an active batch exist")
	}

	if err := s.batchRepo.Update(id, batch); err != nil {
		return err
	}
	return nil
}

func (s *batchService) Delete(id int) error {
	if err := s.batchRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *batchService) GetByID(id int) (*model.Batch, error) {
	batch, err := s.batchRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func (s *batchService) GetAll(limit, page int, q string) ([]model.Batch, error) {
	batches, err := s.batchRepo.GetAll(limit, page, q)
	if err != nil {
		return nil, err
	}

	return batches, nil
}
