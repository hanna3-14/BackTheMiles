package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

func GetDistances() ([]models.Distance, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return []models.Distance{}, err
	}
	defer db.Close()

	return selectAllDistancesFromDB(db)
}

func GetDistanceById(id string) (models.Distance, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return models.Distance{}, err
	}
	defer db.Close()

	return selectDistanceByIdFromDB(db, id)
}

func PostDistance(distance models.Distance) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	err = createDistancesTable(db)
	if err != nil {
		return err
	}

	return insertDistanceIntoDB(db, distance)
}

func PatchDistance(id string, modifiedDistance models.Distance) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	distance, err := selectDistanceByIdFromDB(db, id)
	if err != nil {
		return err
	}

	return updateDistanceInDB(db, distance, modifiedDistance)
}

func DeleteDistance(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	return deleteDistanceFromDB(db, id)
}
