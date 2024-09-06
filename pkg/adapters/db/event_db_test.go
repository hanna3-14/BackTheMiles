package db_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hanna3-14/BackTheMiles/pkg/adapters/db"
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	tests := []struct {
		name        string
		Event       domain.Event
		expectedErr error
	}{
		{
			name:        "database is closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "happy path",
			Event: domain.Event{
				Name:     "Baden Marathon",
				Location: "Karlsruhe",
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		if tt.expectedErr != nil {
			mock.ExpectExec(`CREATE TABLE IF NOT EXISTS events ( name TEXT, location TEXT );`).
				WithoutArgs().
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectExec(`CREATE TABLE IF NOT EXISTS events ( name TEXT, location TEXT );`).
				WithoutArgs().
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectPrepare(`INSERT INTO events (name, location) VALUES (?, ?)`)
			mock.ExpectExec(`INSERT INTO events (name, location) VALUES (?, ?)`).
				WithArgs(tt.Event.Name, tt.Event.Location).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err = mockDBAdapter.CreateEvent(tt.Event)
		if tt.expectedErr != nil {
			require.NotNil(t, err)
			assert.Equal(t, tt.expectedErr, err)
		} else {
			require.Nil(t, err)
		}
	}
}

func TestFindAllEvents(t *testing.T) {
	tests := []struct {
		name           string
		expectedEvents []domain.Event
		expectedErr    error
	}{
		{
			name:           "no rows in database",
			expectedErr:    sql.ErrNoRows,
			expectedEvents: []domain.Event{},
		},
		{
			name: "happy path",
			expectedEvents: []domain.Event{
				{
					ID:       0,
					Name:     "Baden Marathon",
					Location: "Karlsruhe",
				},
				{
					ID:       1,
					Name:     "Schwarzwald Marathon",
					Location: "Bräunlingen",
				},
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		if tt.expectedErr != nil {
			mock.ExpectQuery(`SELECT rowid, name, location FROM events`).
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(`SELECT rowid, name, location FROM events`).
				WillReturnRows(sqlmock.NewRows([]string{"rowid", "name", "location"}).
					AddRow(0, "Baden Marathon", "Karlsruhe").
					AddRow(1, "Schwarzwald Marathon", "Bräunlingen"))
		}

		actualEvents, err := mockDBAdapter.FindAllEvents()
		assert.Equal(t, tt.expectedEvents, actualEvents)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestFindEventByID(t *testing.T) {
	tests := []struct {
		name          string
		EventID       int
		expectedEvent domain.Event
		expectedErr   error
	}{
		{
			name:          "no rows in database",
			EventID:       0,
			expectedErr:   sql.ErrNoRows,
			expectedEvent: domain.Event{},
		},
		{
			name:    "happy path",
			EventID: 0,
			expectedEvent: domain.Event{
				ID:       0,
				Name:     "Baden Marathon",
				Location: "Karlsruhe",
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare("SELECT rowid, name, location FROM events WHERE rowid = ?")
		if tt.expectedErr != nil {
			mock.ExpectQuery(`SELECT rowid, name, location FROM events WHERE rowid = ?`).
				WithArgs(tt.EventID).WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(`SELECT rowid, name, location FROM events WHERE rowid = ?`).
				WithArgs(tt.EventID).WillReturnRows(sqlmock.NewRows([]string{"rowid", "name", "EventInMeters"}).
				AddRow(0, "Baden Marathon", "Karlsruhe").
				AddRow(1, "Schwarzwald Marathon", "Bräunlingen"))
		}

		actualEvent, err := mockDBAdapter.FindEventByID(tt.EventID)
		assert.Equal(t, tt.expectedEvent, actualEvent)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestDeleteEvent(t *testing.T) {
	tests := []struct {
		name        string
		EventID     int
		expectedErr error
	}{
		{
			name:        "database closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name:    "happy path",
			EventID: 3,
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare(`DELETE FROM events WHERE rowid = ?`).
			WillReturnError(tt.expectedErr)
		if tt.expectedErr == nil {
			mock.ExpectExec(`DELETE FROM events WHERE rowid = ?`).
				WithArgs(tt.EventID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := mockDBAdapter.DeleteEvent(tt.EventID)
		assert.Equal(t, tt.expectedErr, err)
	}
}
