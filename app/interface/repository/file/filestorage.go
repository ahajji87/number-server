package file

import (
	"number-server/app/domain/model"
	"number-server/app/domain/repository"
	"number-server/app/infrastructure/storage"
)

var numbers = make(map[int]bool)
var report model.Report

type numberRepository struct {
	storage storage.FileStorage
}

func NewFileNumberRepository(storage storage.FileStorage) repository.NumberRepository {
	return &numberRepository{
		storage,
	}
}

func (r *numberRepository) Exists(number *model.Number) (bool, error) {
	var num, err = number.ToInt()
	if err != nil {
		return false, err
	}

	if _, ok := numbers[num]; ok {
		report.Duplicats++
		return true, nil
	}

	return false, nil
}

func (r *numberRepository) Save(number *model.Number) error {
	var num, err = number.ToInt()
	if err != nil {
		return err
	}
	numbers[num] = true

	report.Adds++
	report.Total = report.Total + num

	return r.storage.Save(number.Value)
}

func (r *numberRepository) GetReport() *model.Report {
	return &report
}

func (r *numberRepository) DeleteReport() {
	report.Adds = 0
	report.Duplicats = 0
	report.Total = 0
}
