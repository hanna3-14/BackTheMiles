package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const selectEvents string = `
	SELECT
	rowid,
	name,
	location
	FROM events
	`

func createEventsTable(db *sql.DB) error {

	const createStmt string = `
	CREATE TABLE IF NOT EXISTS events (
		name TEXT,
		location TEXT
	);`

	_, err := db.Exec(createStmt)
	return err
}

func selectAllEventsFromDB(db *sql.DB) ([]models.Event, error) {

	var events []models.Event
	response, err := db.Query(selectEvents)
	if err != nil {
		return []models.Event{}, err
	}

	for response.Next() {
		var event models.Event
		err = response.Scan(&event.ID, &event.Name, &event.Location)
		if err != nil {
			return []models.Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}

func selectEventByIdFromDB(db *sql.DB, id string) (models.Event, error) {

	stmt, err := db.Prepare(selectEvents + " WHERE rowid = ?")
	if err != nil {
		return models.Event{}, err
	}

	var event models.Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Location)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func insertEventIntoDB(db *sql.DB, event models.Event) error {

	const insertStmt string = `
	INSERT INTO events (name, location)
	VALUES (?, ?)
	`

	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Location)
	return err
}

func updateEventInDB(db *sql.DB, event models.Event, modifiedEvent models.Event) error {

	const updateStmt string = `
	UPDATE events SET
	name = ?,
	location = ?
	WHERE rowid = ?
	`

	stmt, err := db.Prepare(updateStmt)
	if err != nil {
		return err
	}

	if len(modifiedEvent.Name) == 0 {
		modifiedEvent.Name = event.Name
	}
	if len(modifiedEvent.Location) == 0 {
		modifiedEvent.Location = event.Location
	}

	_, err = stmt.Exec(modifiedEvent.Name, modifiedEvent.Location, event.ID)
	return err
}

func deleteEventFromDB(db *sql.DB, id string) error {

	const deleteStmt string = `
	DELETE FROM events WHERE rowid = ?
	`

	stmt, err := db.Prepare(deleteStmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
