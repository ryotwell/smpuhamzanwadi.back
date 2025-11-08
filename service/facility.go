package service

import (
	"project_sdu/model"
	"project_sdu/repository"
)

type FacilityService interface {
	Create(facility *model.Facility) error
	Update(id int, facility *model.Facility) error
	Delete(id int) error
	GetByID(id int) (*model.Facility, error)
	GetAll(limit, offset int) ([]model.Facility, error)
}

type facilityService struct {
	facilityRepo repository.FacilityRepository
}

func NewfacilityService(facilityRepo repository.FacilityRepository) FacilityService {
	return &facilityService{facilityRepo}
}

func (s *facilityService) Create(facility *model.Facility) error {
	if err := s.facilityRepo.Create(facility); err != nil {
		return err
	}

	return nil
}

func (s *facilityService) Update(id int, facility *model.Facility) error {
	if err := s.facilityRepo.Update(id, facility); err != nil {
		return err
	}

	return nil
}

func (s *facilityService) Delete(id int) error {
	if err := s.facilityRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *facilityService) GetByID(id int) (*model.Facility, error) {
	facility, err := s.facilityRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return facility, nil
}

func (s *facilityService) GetAll(limit, page int) ([]model.Facility, error) {
	facilities, err := s.facilityRepo.GetAll(limit, page)
	if err != nil {
		return nil, err
	}
	
	return facilities, nil
}
