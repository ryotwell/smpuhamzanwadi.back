package service

import (
	// "errors"
	"project_sdu/model"
	"project_sdu/repository"
)

type ParentService interface {
	CreateParent(parent *model.Parent) error
	GetAllParents(limit int, offset int) ([]model.Parent, error)
	GetParentByID(id int) (*model.Parent, error)
	UpdateParent(id int, parent *model.Parent) error
	DeleteParent(id int) error
}

type parentService struct {
	parentRepo repository.ParentRepository
}

func NewParentService(parentRepo repository.ParentRepository) ParentService {
	return &parentService{parentRepo}
}

func (s *parentService) CreateParent(parent *model.Parent) error {
	if err := s.parentRepo.Create(parent); err != nil {
		return err
	}
	return nil
}

func (s *parentService) GetAllParents(limit int, offset int) ([]model.Parent, error) {
	parents, err := s.parentRepo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func (s *parentService) GetParentByID(id int) (*model.Parent, error) {
	parent, err := s.parentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return parent, nil
}

func (s *parentService) UpdateParent(id int, parent *model.Parent) error {
	_, err := s.parentRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.parentRepo.Update(id, parent)
}

func (s *parentService) DeleteParent(id int) error {
	_, err := s.parentRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.parentRepo.Delete(id)
}
