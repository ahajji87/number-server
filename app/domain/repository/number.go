package repository

import (
	"number-server/app/domain/model"
)

type NumberRepository interface {
	Exists(number *model.Number) (bool, error)
	Save(number *model.Number) error
	GetReport() *model.Report
	DeleteReport()
}
