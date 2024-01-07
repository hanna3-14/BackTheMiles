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

	var results []models.Result
	response, err := db.Query("SELECT rowid, name, distance, time, place FROM results")
	if err != nil {
		return []models.Result{}, err
	}

	for response.Next() {
		var result models.Result
		err = response.Scan(&result.ID, &result.Name, &result.Distance, &result.Time, &result.Place)
		if err != nil {
			return []models.Result{}, err
		}
		results = append(results, result)
	}

	return results, nil
}

func GetResultById(id string) (models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return models.Result{}, err
	}

	stmt, err := db.Prepare("SELECT rowid, name, distance, time, place FROM results WHERE rowid = ?")
	if err != nil {
		return models.Result{}, err
	}

	var result models.Result
	err = stmt.QueryRow(id).Scan(&result.ID, &result.Name, &result.Distance, &result.Time, &result.Place)
	if err != nil {
		return models.Result{}, err
	}

	return result, nil
}

func PostResult(result models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}

	const create string = `
	CREATE TABLE IF NOT EXISTS results (
		name TEXT,
		distance TEXT,
		time TEXT,
		place INT
	);`

	_, err = db.Exec(create)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO results(name, distance, time, place) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(result.Name, result.Distance, result.Time, result.Place)
	if err != nil {
		return err
	}

	return nil
}

func PatchResult(id string, modifiedResult models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("SELECT rowid, name, distance, time, place FROM results WHERE rowid = ?")
	if err != nil {
		return err
	}

	var result models.Result
	err = stmt.QueryRow(id).Scan(&result.ID, &result.Name, &result.Distance, &result.Time, &result.Place)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("UPDATE goals SET name = ?, distance = ?, time = ?, place = ? WHERE rowid = ?")
	if err != nil {
		return err
	}

	if len(modifiedResult.Name) == 0 {
		modifiedResult.Name = result.Name
	}
	if len(modifiedResult.Distance) == 0 {
		modifiedResult.Distance = result.Distance
	}
	if len(modifiedResult.Time) == 0 {
		modifiedResult.Time = result.Time
	}
	if modifiedResult.Place == 0 {
		modifiedResult.Place = result.Place
	}

	_, err = stmt.Exec(result.Name, result.Distance, result.Time, result.Place, result.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteResult(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM results WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
