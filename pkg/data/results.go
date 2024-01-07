package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const file = "/results.db"

func GetResults() ([]models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+file)
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

func PostResult(result models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+file)
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
