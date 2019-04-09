package usecase

import (
	"errors"
	"fmt"
	"number-server/app/domain/model"
	"number-server/app/domain/service"
)

type NumberUseCase interface {
	ReadNumber(number *model.Number) error
	Store(numebr *model.Number) error
	GetReport() string
}

type numberUseCase struct {
	numberService service.NumberService
	pipeline      chan *model.Number
}

func NewNumberUseCase(numberService service.NumberService, pipeline chan *model.Number) NumberUseCase {
	return &numberUseCase{
		numberService,
		pipeline,
	}
}

func (u *numberUseCase) ReadNumber(number *model.Number) error {
	if !u.numberService.IsValidNumber(number) {
		return errors.New("Invalid number")
	}

	u.pipeline <- number

	return nil
}

func (u *numberUseCase) Store(number *model.Number) error {
	err := u.numberService.Store(number)
	if err != nil {
		return err
	}
	return nil
}

func (u *numberUseCase) GetReport() string {
	report := u.numberService.GetCounters()

	result := fmt.Sprintf("%d new numbers, %d duplicated entries, %d total numbers", report.Adds, report.Duplicats, report.Total)
	u.numberService.ResetCounters()

	return result
}
