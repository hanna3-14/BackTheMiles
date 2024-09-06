package db

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
)

const selectEvents string = `
	SELECT
	rowid,
	name,
	location
	FROM events
	`

func (repo *SQLDBAdapter) CreateEvent(event domain.Event) error {
	const createStmt string = `
	CREATE TABLE IF NOT EXISTS events (
		name TEXT,
		location TEXT
	);`

	_, err := repo.Database.Exec(createStmt)
	if err != nil {
		return err
	}

	const insertStmt string = `
	INSERT INTO events (name, location)
	VALUES (?, ?)
	`

	stmt, err := repo.Database.Prepare(insertStmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(event.Name, event.Location)
	return err
}

func (repo *SQLDBAdapter) FindAllEvents() ([]domain.Event, error) {
	var events []domain.Event
	response, err := repo.Database.Query(selectEvents)
	if err != nil {
		return []domain.Event{}, err
	}

	for response.Next() {
		var event domain.Event
		err = response.Scan(&event.ID, &event.Name, &event.Location)
		if err != nil {
			return []domain.Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (repo *SQLDBAdapter) FindEventByID(id int) (domain.Event, error) {
	stmt, err := repo.Database.Prepare(selectEvents + " WHERE rowid = ?")
	if err != nil {
		return domain.Event{}, err
	}

	var event domain.Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Location)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (repo *SQLDBAdapter) UpdateEvent(id int, event, modifiedEvent domain.Event) error {
	const updateStmt string = `
	UPDATE events SET
	name = ?,
	location = ?
	WHERE rowid = ?
	`

	stmt, err := repo.Database.Prepare(updateStmt)
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

func (repo *SQLDBAdapter) DeleteEvent(id int) error {
	const deleteStmt string = `
	DELETE FROM events WHERE rowid = ?
	`

	stmt, err := repo.Database.Prepare(deleteStmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
