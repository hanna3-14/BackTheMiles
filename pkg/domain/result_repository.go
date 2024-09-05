package domain

type IResultRepository interface {
	CreateResult(result Result) error
	FindAllResults() ([]Result, error)
	FindResultByID(id int) (Result, error)
	UpdateResult(id int, result, modifiedResult Result) error
	DeleteResult(id int) error
}
