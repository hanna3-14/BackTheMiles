package db

const selectEventResults string = `
	SELECT DISTINCT
	event_id,
	result_id
	FROM eventResults
	`

func (repo *SQLDBAdapter) CreateEventResult(eventId int) error {
	const createStmt string = `
	CREATE TABLE IF NOT EXISTS eventResults (
		event_id INT,
		result_id INT
	);`

	_, err := repo.Database.Exec(createStmt)
	if err != nil {
		return err
	}

	// get the resultId of the latest result
	response, err := repo.Database.Query("SELECT MAX(rowid) FROM results")
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
	stmt, err := repo.Database.Prepare("INSERT INTO eventResults (event_id, result_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventId, resultId)
	return err
}

func (repo *SQLDBAdapter) FindAllEventResults() (map[int][]int, error) {
	eventResults := map[int][]int{}
	response, err := repo.Database.Query(selectEventResults)
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

func (repo *SQLDBAdapter) UpdateEventResult(id int, resultId int, eventId int) error {
	stmt, err := repo.Database.Prepare("UPDATE eventResults SET event_id = ? WHERE result_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventId, resultId)
	return err
}

func (repo *SQLDBAdapter) DeleteEventResult(resultId int) error {
	stmt, err := repo.Database.Prepare("DELETE FROM eventResults WHERE result_id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(resultId)
	return err
}
