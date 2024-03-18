package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const resultsFile = "/results.db"

func GetResults() ([]models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return []models.Result{}, err
	}
	defer db.Close()

	return selectAllResultsFromDB(db)
}

func GetResultById(id string) (models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return models.Result{}, err
	}
	defer db.Close()

	return selectResultByIdFromDB(db, id)
}

func PostResult(result models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	err = createResultsTable(db)
	if err != nil {
		return err
	}

	return insertResultIntoDB(db, result)
}

func PatchResult(id string, modifiedResult models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := selectResultByIdFromDB(db, id)
	if err != nil {
		return err
	}

	return updateResultInDB(db, result, modifiedResult)
}

func DeleteResult(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	return deleteResultFromDB(db, id)
}
