package service

import (
	"project_sdu/model"
	"project_sdu/repository"
)

type CurriculumService interface {
	Create(ex *model.Curriculum) error
	Update(id int, ex *model.Curriculum) error
	Delete(id int) error
	GetByID(id int) (*model.Curriculum, error)
	GetAll(limit, offset int) ([]model.Curriculum, error)
	GetByCategory(category string, limit, page int) ([]model.Curriculum, error)
}

type curriculumService struct {
	curriculumRepo repository.CurriculumRepository
}

func NewCurriculumService(curriculumRepo repository.CurriculumRepository) CurriculumService {
	return &curriculumService{curriculumRepo}
}

func (s *curriculumService) Create(ex *model.Curriculum) error {
	if err := s.curriculumRepo.Create(ex); err != nil {
		return err
	}

	return nil
}

func (s *curriculumService) Update(id int, ex *model.Curriculum) error {
	if err := s.curriculumRepo.Update(id, ex); err != nil {
		return err
	}

	return nil
}

func (s *curriculumService) Delete(id int) error {
	if err := s.curriculumRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *curriculumService) GetByID(id int) (*model.Curriculum, error) {
	curriculum, err := s.curriculumRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return curriculum, nil
}

func (s *curriculumService) GetAll(limit, page int) ([]model.Curriculum, error) {
	curriculum, err := s.curriculumRepo.GetAll(limit, page)
	if err != nil {
		return nil, err
	}
	
	return curriculum, nil
}

func (s *curriculumService) GetByCategory(category string, limit, page int) ([]model.Curriculum, error) {
	return s.curriculumRepo.GetByCategory(category, limit, page)
}