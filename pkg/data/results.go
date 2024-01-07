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
	name,
	distance,
	time_gross,
	time_net,
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
		err = response.Scan(&result.ID, &result.Name, &result.Distance, &result.TimeGross, &result.TimeNet, &result.Category, &result.Agegroup, &result.PlaceTotal, &result.PlaceCategory, &result.PlaceAgegroup, &result.FinisherTotal, &result.FinisherCategory, &result.FinisherAgegroup)
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
	name,
	distance,
	time_gross,
	time_net,
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
	err = stmt.QueryRow(id).Scan(&result.ID, &result.Name, &result.Distance, &result.TimeGross, &result.TimeNet, &result.Category, &result.Agegroup, &result.PlaceTotal, &result.PlaceCategory, &result.PlaceAgegroup, &result.FinisherTotal, &result.FinisherCategory, &result.FinisherAgegroup)
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
		name TEXT,
		distance TEXT,
		time_gross TEXT,
		time_net TEXT,
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
		name,
		distance,
		time_gross,
		time_net,
		category,
		agegroup,
		place_total,
		place_category,
		place_agegroup,
		finisher_total,
		finisher_category,
		finisher_agegroup
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
	`

	stmt, err := db.Prepare(insert)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(result.Name, result.Distance, result.TimeGross, result.TimeNet, result.Category, result.Agegroup, result.PlaceTotal, result.PlaceCategory, result.PlaceAgegroup, result.FinisherTotal, result.FinisherCategory, result.FinisherAgegroup)
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
	name,
	distance,
	time_gross,
	time_net,
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
	err = stmt.QueryRow(id).Scan(&result.ID, &result.Name, &result.Distance, &result.TimeGross, &result.TimeNet, &result.Category, &result.Agegroup, &result.PlaceTotal, &result.PlaceCategory, &result.PlaceAgegroup, &result.FinisherTotal, &result.FinisherCategory, &result.FinisherAgegroup)
	if err != nil {
		return err
	}

	const update string = `
	UPDATE results SET
	name = ?,
	distance = ?,
	time_gross = ?,
	time_net = ?,
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

	if len(modifiedResult.Name) == 0 {
		modifiedResult.Name = result.Name
	}
	if len(modifiedResult.Distance) == 0 {
		modifiedResult.Distance = result.Distance
	}
	if len(modifiedResult.TimeGross) == 0 {
		modifiedResult.TimeGross = result.TimeGross
	}
	if len(modifiedResult.TimeNet) == 0 {
		modifiedResult.TimeNet = result.TimeNet
	}
	if len(modifiedResult.Category) == 0 {
		modifiedResult.Category = result.Category
	}
	if len(modifiedResult.Agegroup) == 0 {
		modifiedResult.Agegroup = result.Agegroup
	}
	if modifiedResult.PlaceTotal == 0 {
		modifiedResult.PlaceTotal = result.PlaceTotal
	}
	if modifiedResult.PlaceCategory == 0 {
		modifiedResult.PlaceCategory = result.PlaceCategory
	}
	if modifiedResult.PlaceAgegroup == 0 {
		modifiedResult.PlaceAgegroup = result.PlaceAgegroup
	}
	if modifiedResult.FinisherTotal == 0 {
		modifiedResult.FinisherTotal = result.FinisherTotal
	}
	if modifiedResult.FinisherCategory == 0 {
		modifiedResult.FinisherCategory = result.FinisherCategory
	}
	if modifiedResult.FinisherAgegroup == 0 {
		modifiedResult.FinisherAgegroup = result.FinisherAgegroup
	}

	_, err = stmt.Exec(modifiedResult.Name, modifiedResult.Distance, modifiedResult.TimeGross, modifiedResult.TimeNet, modifiedResult.Category, modifiedResult.Agegroup, modifiedResult.PlaceTotal, modifiedResult.PlaceCategory, modifiedResult.PlaceAgegroup, modifiedResult.FinisherTotal, modifiedResult.FinisherCategory, modifiedResult.FinisherAgegroup, modifiedResult.ID)
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
