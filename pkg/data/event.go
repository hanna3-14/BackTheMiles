package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const eventsFile = "/events.db"

func GetEvents() ([]models.Event, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return []models.Event{}, err
	}
	defer db.Close()

	const selectEvents string = `
	SELECT
	rowid,
	name,
	location
	FROM events
	`

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

func GetEventById(id string) (models.Event, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return models.Event{}, err
	}
	defer db.Close()

	const selectEvent string = `
	SELECT
	rowid,
	name,
	location,
	FROM events WHERE rowid = ?
	`

	stmt, err := db.Prepare(selectEvent)
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

func PostEvent(event models.Event) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	const createEvents string = `
	CREATE TABLE IF NOT EXISTS events (
	name TEXT,
	location TEXT
	);`

	_, err = db.Exec(createEvents)
	if err != nil {
		return err
	}

	const insertEvent string = `
	INSERT INTO events (
		name,
		location
	) VALUES (?, ?)
	`

	stmt, err := db.Prepare(insertEvent)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Location)
	if err != nil {
		return err
	}

	return nil
}

func PatchEvent(id string, modifiedEvent models.Event) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	const selectStmt string = `
	SELECT
	rowid,
	name,
	location
	FROM events WHERE rowid = ?
	`

	stmt, err := db.Prepare(selectStmt)
	if err != nil {
		return err
	}

	var event models.Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Location)
	if err != nil {
		return err
	}

	const updateStmt string = `
	UPDATE events SET
	name = ?,
	location = ?
	WHERE rowid = ?
	`

	stmt, err = db.Prepare(updateStmt)
	if err != nil {
		return err
	}

	if len(modifiedEvent.Name) == 0 {
		modifiedEvent.Name = event.Name
	}
	if len(modifiedEvent.Location) == 0 {
		modifiedEvent.Location = event.Location
	}

	_, err = stmt.Exec(modifiedEvent.Name, modifiedEvent.Location, modifiedEvent.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	const deleteStmt string = `
	DELETE FROM events WHERE rowid = ?
	`

	stmt, err := db.Prepare(deleteStmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
