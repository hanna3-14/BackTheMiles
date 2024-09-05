package domain

type IDistanceRepository interface {
	CreateDistance(distance Distance) error
	FindAllDistances() ([]Distance, error)
	FindDistanceByID(id int) (Distance, error)
	UpdateDistance(id int, distance, modifiedDistance Distance) error
	DeleteDistance(id int) error
}
