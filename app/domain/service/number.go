package service

import (
	"number-server/app/domain/model"
	"number-server/app/domain/repository"
)

const MAX_LEN = 9

type NumberService interface {
	Store(number *model.Number) error
	GetCounters() *model.Report
	ResetCounters()
	IsValidNumber(number *model.Number) bool
}

type numberService struct {
	repo repository.NumberRepository
}

func NewNumberService(repository repository.NumberRepository) NumberService {
	return &numberService{
		repo: repository,
	}
}

func (s *numberService) Store(number *model.Number) error {
	var numExists, err = s.repo.Exists(number)
	if err != nil {
		return err
	}
	if numExists {
		return nil
	}

	if err := s.repo.Save(number); err != nil {
		return err
	}

	return nil
}

func (s *numberService) GetCounters() *model.Report {
	return s.repo.GetReport()
}

func (s *numberService) ResetCounters() {
	s.repo.DeleteReport()
}

func (s *numberService) IsValidNumber(number *model.Number) bool {
	if len(number.Value) != MAX_LEN {
		return false
	}
	_, err := number.ToInt()
	if err != nil {
		return false
	}

	return true
}
