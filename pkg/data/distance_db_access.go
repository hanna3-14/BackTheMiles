package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const selectDistances string = `
	SELECT
	rowid,
	name,
	distanceInMeters
	FROM distances
	`

func createDistancesTable(db *sql.DB) error {

	const createStmt string = `
	CREATE TABLE IF NOT EXISTS distances (
		name TEXT,
		distanceInMeters INT
	);`

	_, err := db.Exec(createStmt)
	return err
}

func selectAllDistancesFromDB(db *sql.DB) ([]models.Distance, error) {

	var distances []models.Distance
	response, err := db.Query(selectDistances)
	if err != nil {
		return []models.Distance{}, err
	}

	for response.Next() {
		var distance models.Distance
		err = response.Scan(&distance.ID, &distance.Name, &distance.DistanceInMeters)
		if err != nil {
			return []models.Distance{}, err
		}
		distances = append(distances, distance)
	}
	return distances, nil
}

func selectDistanceByIdFromDB(db *sql.DB, id string) (models.Distance, error) {

	stmt, err := db.Prepare(selectDistances + " WHERE rowid = ?")
	if err != nil {
		return models.Distance{}, err
	}

	var distance models.Distance
	err = stmt.QueryRow(id).Scan(&distance.ID, &distance.Name, &distance.DistanceInMeters)
	if err != nil {
		return models.Distance{}, err
	}
	return distance, nil
}

func insertDistanceIntoDB(db *sql.DB, distance models.Distance) error {

	stmt, err := db.Prepare("INSERT INTO distances (name, distanceInMeters) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(distance.Name, distance.DistanceInMeters)
	return err
}

func updateDistanceInDB(db *sql.DB, distance models.Distance, modifiedDistance models.Distance) error {

	stmt, err := db.Prepare("UPDATE distances SET name = ?, distanceInMeters = ? WHERE rowid = ?")
	if err != nil {
		return err
	}

	if len(modifiedDistance.Name) == 0 {
		modifiedDistance.Name = distance.Name
	}
	if modifiedDistance.DistanceInMeters == 0 {
		modifiedDistance.DistanceInMeters = distance.DistanceInMeters
	}

	_, err = stmt.Exec(distance.Name, distance.DistanceInMeters, distance.ID)
	return err
}

func deleteDistanceFromDB(db *sql.DB, id string) error {

	stmt, err := db.Prepare("DELETE FROM distances WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
