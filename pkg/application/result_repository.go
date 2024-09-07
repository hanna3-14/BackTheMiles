package application

import "github.com/hanna3-14/BackTheMiles/pkg/domain"

type IResultRepository interface {
	CreateResult(result domain.Result) error
	FindAllResults() ([]domain.Result, error)
	FindResultByID(id int) (domain.Result, error)
	UpdateResult(id int, result, modifiedResult domain.Result) error
	DeleteResult(id int) error
}
