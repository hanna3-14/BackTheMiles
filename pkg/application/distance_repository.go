package application

import "github.com/hanna3-14/BackTheMiles/pkg/domain"

type IDistanceRepository interface {
	CreateDistance(distance domain.Distance) error
	FindAllDistances() ([]domain.Distance, error)
	FindDistanceByID(id int) (domain.Distance, error)
	UpdateDistance(id int, distance, modifiedDistance domain.Distance) error
	DeleteDistance(id int) error
}
