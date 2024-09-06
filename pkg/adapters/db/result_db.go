package db

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"

	_ "github.com/mattn/go-sqlite3"
)

const selectResults string = `
	SELECT
	rowid,
	event_id,
	date,
	distance_id,
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

func (repo *SQLDBAdapter) CreateResult(result domain.Result) error {
	const createStmt string = `
	CREATE TABLE IF NOT EXISTS results (
		event_id INT,
		date TEXT,
		distance_id INT,
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

	_, err := repo.Database.Exec(createStmt)
	if err != nil {
		return err
	}

	const insertStmt string = `
	INSERT INTO results (
		event_id,
		date,
		distance_id,
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

	stmt, err := repo.Database.Prepare(insertStmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		result.EventID,
		result.Date,
		result.DistanceID,
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

func (repo *SQLDBAdapter) FindAllResults() ([]domain.Result, error) {
	var results []domain.Result
	response, err := repo.Database.Query(selectResults)
	if err != nil {
		return []domain.Result{}, err
	}

	for response.Next() {
		var result domain.Result
		err = response.Scan(
			&result.ResultID,
			&result.EventID,
			&result.Date,
			&result.DistanceID,
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
			return []domain.Result{}, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (repo *SQLDBAdapter) FindResultByID(id int) (domain.Result, error) {
	stmt, err := repo.Database.Prepare(selectResults + " WHERE rowid = ?")
	if err != nil {
		return domain.Result{}, err
	}

	var result domain.Result
	err = stmt.QueryRow(id).Scan(
		&result.ResultID,
		&result.EventID,
		&result.Date,
		&result.DistanceID,
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
		return domain.Result{}, err
	}
	return result, nil
}

func (repo *SQLDBAdapter) UpdateResult(id int, storedResult, modifiedResult domain.Result) error {
	const update string = `
	UPDATE results SET
	event_id = ?,
	date = ?,
	distance_id = ?,
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

	stmt, err := repo.Database.Prepare(update)
	if err != nil {
		return err
	}

	// keep the original value if the modified value is empty
	if modifiedResult.EventID == 0 {
		modifiedResult.EventID = storedResult.Finisher.Total
	} else {
		repo.updateEventResultByResultID(storedResult.ResultID, modifiedResult.EventID)
	}
	if len(modifiedResult.Date) == 0 {
		modifiedResult.Date = storedResult.Date
	}
	if modifiedResult.DistanceID == 0 {
		modifiedResult.DistanceID = storedResult.DistanceID
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
		modifiedResult.DistanceID,
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

func (repo *SQLDBAdapter) DeleteResult(id int) error {
	stmt, err := repo.Database.Prepare("DELETE FROM results WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}

func (repo *SQLDBAdapter) updateEventResultByResultID(resultId int, eventId int) error {

	stmt, err := repo.Database.Prepare("UPDATE eventResults SET event_id = ? WHERE result_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventId, resultId)
	return err
}
