package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const selectResults string = `
	SELECT
	rowid,
	date,
	distance,
	time_gross_hours,
	time_gross_minutes,
	time_gross_seconds,
	time_net_hours,
	time_net_minutes,
	time_net_seconds,
	category,
	agegroup,
	place_total,
	place_category,
	place_agegroup,
	finisher_total,
	finisher_category,
	finisher_agegroup
	FROM results
	`

func createResultsTable(db *sql.DB) error {

	const createStmt string = `
	CREATE TABLE IF NOT EXISTS results (
		date TEXT,
		distance TEXT,
		time_gross_hours INT,
		time_gross_minutes INT,
		time_gross_seconds INT,
		time_net_hours INT,
		time_net_minutes INT,
		time_net_seconds INT,
		category TEXT,
		agegroup TEXT,
		place_total INT,
		place_category INT,
		place_agegroup INT,
		finisher_total INT,
		finisher_category INT,
		finisher_agegroup INT
	);`

	_, err := db.Exec(createStmt)
	return err
}

func selectAllResultsFromDB(db *sql.DB) ([]models.Result, error) {

	var results []models.Result
	response, err := db.Query(selectResults)
	if err != nil {
		return []models.Result{}, err
	}

	for response.Next() {
		var result models.Result
		err = response.Scan(
			&result.ID,
			&result.Date,
			&result.Distance,
			&result.TimeGross.Hours,
			&result.TimeGross.Minutes,
			&result.TimeGross.Seconds,
			&result.TimeNet.Hours,
			&result.TimeNet.Minutes,
			&result.TimeNet.Seconds,
			&result.Category,
			&result.Agegroup,
			&result.Place.Total,
			&result.Place.Category,
			&result.Place.Agegroup,
			&result.Finisher.Total,
			&result.Finisher.Category,
			&result.Finisher.Agegroup,
		)
		if err != nil {
			return []models.Result{}, err
		}
		results = append(results, result)
	}
	return results, nil
}

func selectResultByIdFromDB(db *sql.DB, id string) (models.Result, error) {

	stmt, err := db.Prepare(selectResults + " WHERE rowid = ?")
	if err != nil {
		return models.Result{}, err
	}

	var result models.Result
	err = stmt.QueryRow(id).Scan(
		&result.ID,
		&result.Date,
		&result.Distance,
		&result.TimeGross.Hours,
		&result.TimeGross.Minutes,
		&result.TimeGross.Seconds,
		&result.TimeNet.Hours,
		&result.TimeNet.Minutes,
		&result.TimeNet.Seconds,
		&result.Category,
		&result.Agegroup,
		&result.Place.Total,
		&result.Place.Category,
		&result.Place.Agegroup,
		&result.Finisher.Total,
		&result.Finisher.Category,
		&result.Finisher.Agegroup,
	)
	if err != nil {
		return models.Result{}, err
	}
	return result, nil
}

func insertResultIntoDB(db *sql.DB, result models.Result) error {

	const insert string = `
	INSERT INTO results (
		date,
		distance,
		time_gross_hours,
		time_gross_minutes,
		time_gross_seconds,
		time_net_hours,
		time_net_minutes,
		time_net_seconds,
		category,
		agegroup,
		place_total,
		place_category,
		place_agegroup,
		finisher_total,
		finisher_category,
		finisher_agegroup
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`

	stmt, err := db.Prepare(insert)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		result.Date,
		result.Distance,
		result.TimeGross.Hours,
		result.TimeGross.Minutes,
		result.TimeGross.Seconds,
		result.TimeNet.Hours,
		result.TimeNet.Minutes,
		result.TimeNet.Seconds,
		result.Category,
		result.Agegroup,
		result.Place.Total,
		result.Place.Category,
		result.Place.Agegroup,
		result.Finisher.Total,
		result.Finisher.Category,
		result.Finisher.Agegroup,
	)
	return err
}

func updateResultInDB(db *sql.DB, result models.Result, modifiedResult models.Result) error {

	const update string = `
	UPDATE results SET
	date = ?,
	distance = ?,
	time_gross_hours = ?,
	time_gross_minutes = ?,
	time_gross_seconds = ?,
	time_net_hours = ?,
	time_net_minutes = ?,
	time_net_seconds = ?,
	category = ?,
	agegroup = ?,
	place_total = ?,
	place_category = ?,
	place_agegroup = ?,
	finisher_total = ?,
	finisher_category = ?,
	finisher_agegroup = ?
	WHERE rowid = ?
	`

	stmt, err := db.Prepare(update)
	if err != nil {
		return err
	}

	// keep the original value if the modified value is empty
	if len(modifiedResult.Date) == 0 {
		modifiedResult.Date = result.Date
	}
	if len(modifiedResult.Distance) == 0 {
		modifiedResult.Distance = result.Distance
	}
	if modifiedResult.TimeGross.Hours == 0 {
		modifiedResult.TimeGross.Hours = result.TimeGross.Hours
	}
	if modifiedResult.TimeGross.Minutes == 0 {
		modifiedResult.TimeGross.Minutes = result.TimeGross.Minutes
	}
	if modifiedResult.TimeGross.Seconds == 0 {
		modifiedResult.TimeGross.Seconds = result.TimeGross.Seconds
	}
	if modifiedResult.TimeNet.Hours == 0 {
		modifiedResult.TimeNet.Hours = result.TimeNet.Hours
	}
	if modifiedResult.TimeNet.Minutes == 0 {
		modifiedResult.TimeNet.Minutes = result.TimeNet.Minutes
	}
	if modifiedResult.TimeNet.Seconds == 0 {
		modifiedResult.TimeNet.Seconds = result.TimeNet.Seconds
	}
	if len(modifiedResult.Category) == 0 {
		modifiedResult.Category = result.Category
	}
	if len(modifiedResult.Agegroup) == 0 {
		modifiedResult.Agegroup = result.Agegroup
	}
	if modifiedResult.Place.Total == 0 {
		modifiedResult.Place.Total = result.Place.Total
	}
	if modifiedResult.Place.Category == 0 {
		modifiedResult.Place.Category = result.Place.Category
	}
	if modifiedResult.Place.Agegroup == 0 {
		modifiedResult.Place.Agegroup = result.Place.Agegroup
	}
	if modifiedResult.Finisher.Total == 0 {
		modifiedResult.Finisher.Total = result.Finisher.Total
	}
	if modifiedResult.Finisher.Category == 0 {
		modifiedResult.Finisher.Category = result.Finisher.Category
	}
	if modifiedResult.Finisher.Agegroup == 0 {
		modifiedResult.Finisher.Agegroup = result.Finisher.Agegroup
	}

	_, err = stmt.Exec(
		modifiedResult.Date,
		modifiedResult.Distance,
		modifiedResult.TimeGross.Hours,
		modifiedResult.TimeGross.Minutes,
		modifiedResult.TimeGross.Seconds,
		modifiedResult.TimeNet.Hours,
		modifiedResult.TimeNet.Minutes,
		modifiedResult.TimeNet.Seconds,
		modifiedResult.Category,
		modifiedResult.Agegroup,
		modifiedResult.Place.Total,
		modifiedResult.Place.Category,
		modifiedResult.Place.Agegroup,
		modifiedResult.Finisher.Total,
		modifiedResult.Finisher.Category,
		modifiedResult.Finisher.Agegroup,
		modifiedResult.ID,
	)
	return err
}

func deleteResultFromDB(db *sql.DB, id string) error {

	stmt, err := db.Prepare("DELETE FROM results WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
