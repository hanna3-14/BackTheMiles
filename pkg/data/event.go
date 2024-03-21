package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

func GetEvents() ([]models.Event, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return []models.Event{}, err
	}
	defer db.Close()

	// select all events from db
	events, err := selectAllEventsFromDB(db)
	if err != nil {
		return []models.Event{}, err
	}

	// select eventResults from db and add them to the according event
	eventResults, err := selectAllEventResultsFromDB(db)
	if err != nil {
		return []models.Event{}, err
	}

	for i := range events {
		events[i].ResultIDs = eventResults[events[i].ID]
	}

	return events, err
}

func GetEventById(id string) (models.Event, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return models.Event{}, err
	}
	defer db.Close()

	// select event from db
	event, err := selectEventByIdFromDB(db, id)
	if err != nil {
		return models.Event{}, err
	}

	// select eventResults from db and add them to the event
	eventResults, err := selectAllEventResultsFromDB(db)
	if err != nil {
		return models.Event{}, err
	}

	event.ResultIDs = eventResults[event.ID]
	return event, nil
}

func PostEvent(event models.Event) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	err = createEventsTable(db)
	if err != nil {
		return err
	}

	return insertEventIntoDB(db, event)
}

func PatchEvent(id string, modifiedEvent models.Event) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	event, err := selectEventByIdFromDB(db, id)
	if err != nil {
		return err
	}

	return updateEventInDB(db, event, modifiedEvent)
}

func DeleteEvent(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	err = deleteEventResultByResultID(db, id)
	if err != nil {
		return err
	}

	return deleteEventFromDB(db, id)
}
