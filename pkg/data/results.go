package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const databaseFile = "/marathon.db"

func GetResults() ([]models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return []models.Result{}, err
	}
	defer db.Close()

	return selectAllResultsFromDB(db)
}

func GetResultById(id string) (models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return models.Result{}, err
	}
	defer db.Close()

	return selectResultByIdFromDB(db, id)
}

func PostResult(result models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	// insert result into results table
	err = createResultsTable(db)
	if err != nil {
		return err
	}

	err = insertResultIntoDB(db, result)
	if err != nil {
		return err
	}

	// insert eventResults into the according db table
	err = createEventResultsTable(db)
	if err != nil {
		return err
	}

	return insertEventResultIntoDB(db, result.EventID)
}

func PatchResult(id string, modifiedResult models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	storedResult, err := selectResultByIdFromDB(db, id)
	if err != nil {
		return err
	}

	return updateResultInDB(db, storedResult, modifiedResult)
}

func DeleteResult(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	err = deleteEventResultByResultID(db, id)
	if err != nil {
		return err
	}

	return deleteResultFromDB(db, id)
}
