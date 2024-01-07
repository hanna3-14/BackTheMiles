package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const file string = "./volume/results.db"

func ResultsData() []models.Result {
	return []models.Result{
		{
			ID:       1,
			Name:     "Baden-Marathon",
			Distance: "HM",
			Time:     "02:21:40",
			Place:    2336,
		},
		{
			ID:       2,
			Name:     "Schwarzwald-Marathon",
			Distance: "HM",
			Time:     "02:09:45",
			Place:    535,
		},
		{
			ID:       3,
			Name:     "Bienwald-Marathon",
			Distance: "HM",
			Time:     "02:09:14",
			Place:    928,
		},
		{
			ID:       4,
			Name:     "Freiburg-Marathon",
			Distance: "M",
			Time:     "05:29:09",
			Place:    916,
		},
	}
}

func PostResult(result models.Result) error {

	db, err := sql.Open("sqlite3", file)
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
