package db

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
)

const selectDistances string = `
	SELECT
	rowid,
	name,
	distanceInMeters
	FROM distances
	`

func (repo *SQLDBAdapter) CreateDistance(distance domain.Distance) error {
	const createStmt string = `
	CREATE TABLE IF NOT EXISTS distances (
		name TEXT,
		distanceInMeters INT
	);`

	_, err := repo.Database.Exec(createStmt)
	if err != nil {
		return err
	}

	stmt, err := repo.Database.Prepare("INSERT INTO distances (name, distanceInMeters) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(distance.Name, distance.DistanceInMeters)
	return err
}

func (repo *SQLDBAdapter) FindAllDistances() ([]domain.Distance, error) {
	var distances []domain.Distance
	response, err := repo.Database.Query(selectDistances)
	if err != nil {
		return []domain.Distance{}, err
	}

	for response.Next() {
		var distance domain.Distance
		err = response.Scan(&distance.ID, &distance.Name, &distance.DistanceInMeters)
		if err != nil {
			return []domain.Distance{}, err
		}
		distances = append(distances, distance)
	}
	return distances, nil
}

func (repo *SQLDBAdapter) FindDistanceByID(id int) (domain.Distance, error) {
	stmt, err := repo.Database.Prepare(selectDistances + " WHERE rowid = ?")
	if err != nil {
		return domain.Distance{}, err
	}

	var distance domain.Distance
	err = stmt.QueryRow(id).Scan(&distance.ID, &distance.Name, &distance.DistanceInMeters)
	if err != nil {
		return domain.Distance{}, err
	}
	return distance, nil
}

func (repo *SQLDBAdapter) UpdateDistance(id int, distance, modifiedDistance domain.Distance) error {
	stmt, err := repo.Database.Prepare("UPDATE distances SET name = ?, distanceInMeters = ? WHERE rowid = ?")
	if err != nil {
		return err
	}

	if len(modifiedDistance.Name) == 0 {
		modifiedDistance.Name = distance.Name
	}
	if modifiedDistance.DistanceInMeters == 0 {
		modifiedDistance.DistanceInMeters = distance.DistanceInMeters
	}

	_, err = stmt.Exec(modifiedDistance.Name, modifiedDistance.DistanceInMeters, distance.ID)
	return err
}

func (repo *SQLDBAdapter) DeleteDistance(id int) error {
	stmt, err := repo.Database.Prepare("DELETE FROM distances WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
