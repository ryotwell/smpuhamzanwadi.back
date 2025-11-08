package service

import (
	"project_sdu/model"
	"project_sdu/repository"
)

type ExtracurricularService interface {
	Create(ex *model.Extracurricular) error
	Update(id int, ex *model.Extracurricular) error
	Delete(id int) error
	GetByID(id int) (*model.Extracurricular, error)
	GetAll(limit, offset int) ([]model.Extracurricular, error)
}

type extracurricularService struct {
	exRepo repository.ExtracurricularRepository
}

func NewExtracurricularService(exRepo repository.ExtracurricularRepository) ExtracurricularService {
	return &extracurricularService{exRepo}
}

func (s *extracurricularService) Create(ex *model.Extracurricular) error {
	if err := s.exRepo.Create(ex); err != nil {
		return err
	}

	return nil
}

func (s *extracurricularService) Update(id int, ex *model.Extracurricular) error {
	if err := s.exRepo.Update(id, ex); err != nil {
		return err
	}

	return nil
}

func (s *extracurricularService) Delete(id int) error {
	if err := s.exRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *extracurricularService) GetByID(id int) (*model.Extracurricular, error) {
	extra, err := s.exRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return extra, nil
}

func (s *extracurricularService) GetAll(limit, page int) ([]model.Extracurricular, error) {
	extras, err := s.exRepo.GetAll(limit, page)
	if err != nil {
		return nil, err
	}
	
	return extras, nil
}
