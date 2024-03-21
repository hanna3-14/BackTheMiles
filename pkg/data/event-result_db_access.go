package data

import (
	"database/sql"
)

const selectEventResults string = `
	SELECT DISTINCT
	event_id,
	result_id
	FROM eventResults
	`

func createEventResultsTable(db *sql.DB) error {

	const createStmt string = `
	CREATE TABLE IF NOT EXISTS eventResults (
		event_id INT,
		result_id INT
	);`

	_, err := db.Exec(createStmt)
	return err
}

func selectAllEventResultsFromDB(db *sql.DB) (map[int][]int, error) {

	eventResults := map[int][]int{}
	response, err := db.Query(selectEventResults)
	if err != nil {
		return eventResults, err
	}

	for response.Next() {
		var eventId int
		var resultId int
		err = response.Scan(&eventId, &resultId)
		if err != nil {
			return eventResults, err
		}
		eventResults[eventId] = append(eventResults[eventId], resultId)
	}
	return eventResults, nil
}

func insertEventResultIntoDB(db *sql.DB, eventId int) error {

	// get the resultId of the latest result
	response, err := db.Query("SELECT MAX(rowid) FROM results")
	if err != nil {
		return err
	}

	var resultId int
	for response.Next() {
		err = response.Scan(&resultId)
		if err != nil {
			return err
		}
	}

	// insert eventId and resultId into eventResults
	stmt, err := db.Prepare("INSERT INTO eventResults (event_id, result_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventId, resultId)
	return err
}

func updateEventResultByResultID(db *sql.DB, resultId int, eventId int) error {

	stmt, err := db.Prepare("UPDATE eventResults SET event_id = ? WHERE result_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventId, resultId)
	return err
}

func deleteEventResultByResultID(db *sql.DB, resultId string) error {

	stmt, err := db.Prepare("DELETE FROM eventResults WHERE result_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(resultId)
	return err
}
