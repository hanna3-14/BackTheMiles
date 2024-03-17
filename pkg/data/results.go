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

	const selectStmt string = `
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

	var results []models.Result
	response, err := db.Query(selectStmt)
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

func GetResultById(id string) (models.Result, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return models.Result{}, err
	}
	defer db.Close()

	const selectStmt string = `
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
	FROM results WHERE rowid = ?
	`

	stmt, err := db.Prepare(selectStmt)
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

func PostResult(result models.Result) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+resultsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	const create string = `
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

	_, err = db.Exec(create)
	if err != nil {
		return err
	}

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
	defer db.Close()

	const selectStmt string = `
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
	FROM results WHERE rowid = ?
	`

	stmt, err := db.Prepare(selectStmt)
	if err != nil {
		return err
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
		return err
	}

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

	stmt, err = db.Prepare(update)
	if err != nil {
		return err
	}

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
	defer db.Close()

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
