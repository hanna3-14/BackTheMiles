package application

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
)

type DistanceRequestService struct {
	repo IDistanceRepository
}

func NewDistanceRequestService(distanceRepo IDistanceRepository) (DistanceRequestService, error) {
	return DistanceRequestService{repo: distanceRepo}, nil
}

func (r *DistanceRequestService) GetDistances() ([]domain.Distance, error) {
	return r.repo.FindAllDistances()
}

func (r *DistanceRequestService) GetDistance(id int) (domain.Distance, error) {
	return r.repo.FindDistanceByID(id)
}

func (r *DistanceRequestService) PostDistance(Distance domain.Distance) error {
	return r.repo.CreateDistance(Distance)
}

func (r *DistanceRequestService) PatchDistance(id int, modifiedDistance domain.Distance) error {
	storedDistance, err := r.repo.FindDistanceByID(id)
	if err != nil {
		return err
	}
	return r.repo.UpdateDistance(id, storedDistance, modifiedDistance)
}

func (r *DistanceRequestService) DeleteDistance(id int) error {
	return r.repo.DeleteDistance(id)
}
