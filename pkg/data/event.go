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

	return selectAllEventsFromDB(db)
}

func GetEventById(id string) (models.Event, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return models.Event{}, err
	}
	defer db.Close()

	return selectEventByIdFromDB(db, id)
}

func PostEvent(event models.Event) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+eventsFile)
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

	db, err := sql.Open("sqlite3", path+eventsFile)
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

	db, err := sql.Open("sqlite3", path+eventsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	return deleteEventFromDB(db, id)
}
