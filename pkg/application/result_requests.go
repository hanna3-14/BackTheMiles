package application

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
)

type ResultRequestService struct {
	repo   IResultRepository
	erRepo IEventResultRepository
}

func NewResultRequestService(resultRepo IResultRepository, eventResultRepo IEventResultRepository) (ResultRequestService, error) {
	return ResultRequestService{repo: resultRepo, erRepo: eventResultRepo}, nil
}

func (r *ResultRequestService) GetResults() ([]domain.Result, error) {
	results, err := r.repo.FindAllResults()
	if err != nil {
		return []domain.Result{}, err
	}
	for i := range results {
		results[i] = calculateRelativePlaces(results[i])
	}
	return results, nil
}

func (r *ResultRequestService) GetResult(id int) (domain.Result, error) {
	result, err := r.repo.FindResultByID(id)
	if err != nil {
		return domain.Result{}, err
	}
	result = calculateRelativePlaces(result)
	return result, nil
}

func (r *ResultRequestService) PostResult(result domain.Result) error {
	r.repo.CreateResult(result)
	return r.erRepo.CreateEventResult(result.EventID)
}

func (r *ResultRequestService) PatchResult(id int, modifiedResult domain.Result) error {
	storedResult, err := r.repo.FindResultByID(id)
	if err != nil {
		return err
	}
	return r.repo.UpdateResult(id, storedResult, modifiedResult)
}

func (r *ResultRequestService) DeleteResult(id int) error {
	r.erRepo.DeleteEventResult(id)
	return r.repo.DeleteResult(id)
}
