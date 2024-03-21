package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const selectResults string = `
	SELECT
	rowid,
	event_id,
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
		event_id INT,
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
			&result.ResultID,
			&result.EventID,
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
		&result.ResultID,
		&result.EventID,
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
		event_id,
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
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`

	stmt, err := db.Prepare(insert)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		result.EventID,
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

func updateResultInDB(db *sql.DB, storedResult models.Result, modifiedResult models.Result) error {

	const update string = `
	UPDATE results SET
	event_id = ?,
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
	if modifiedResult.EventID == 0 {
		modifiedResult.EventID = storedResult.Finisher.Total
	} else {
		updateEventResultByResultID(db, storedResult.ResultID, modifiedResult.EventID)
	}
	if len(modifiedResult.Date) == 0 {
		modifiedResult.Date = storedResult.Date
	}
	if len(modifiedResult.Distance) == 0 {
		modifiedResult.Distance = storedResult.Distance
	}
	if modifiedResult.TimeGross.Hours == 0 {
		modifiedResult.TimeGross.Hours = storedResult.TimeGross.Hours
	}
	if modifiedResult.TimeGross.Minutes == 0 {
		modifiedResult.TimeGross.Minutes = storedResult.TimeGross.Minutes
	}
	if modifiedResult.TimeGross.Seconds == 0 {
		modifiedResult.TimeGross.Seconds = storedResult.TimeGross.Seconds
	}
	if modifiedResult.TimeNet.Hours == 0 {
		modifiedResult.TimeNet.Hours = storedResult.TimeNet.Hours
	}
	if modifiedResult.TimeNet.Minutes == 0 {
		modifiedResult.TimeNet.Minutes = storedResult.TimeNet.Minutes
	}
	if modifiedResult.TimeNet.Seconds == 0 {
		modifiedResult.TimeNet.Seconds = storedResult.TimeNet.Seconds
	}
	if len(modifiedResult.Category) == 0 {
		modifiedResult.Category = storedResult.Category
	}
	if len(modifiedResult.Agegroup) == 0 {
		modifiedResult.Agegroup = storedResult.Agegroup
	}
	if modifiedResult.Place.Total == 0 {
		modifiedResult.Place.Total = storedResult.Place.Total
	}
	if modifiedResult.Place.Category == 0 {
		modifiedResult.Place.Category = storedResult.Place.Category
	}
	if modifiedResult.Place.Agegroup == 0 {
		modifiedResult.Place.Agegroup = storedResult.Place.Agegroup
	}
	if modifiedResult.Finisher.Total == 0 {
		modifiedResult.Finisher.Total = storedResult.Finisher.Total
	}
	if modifiedResult.Finisher.Category == 0 {
		modifiedResult.Finisher.Category = storedResult.Finisher.Category
	}
	if modifiedResult.Finisher.Agegroup == 0 {
		modifiedResult.Finisher.Agegroup = storedResult.Finisher.Agegroup
	}

	_, err = stmt.Exec(
		modifiedResult.EventID,
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
		modifiedResult.ResultID,
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
