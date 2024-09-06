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

func TestCreateDistance(t *testing.T) {
	tests := []struct {
		name        string
		distance    domain.Distance
		expectedErr error
	}{
		{
			name:        "database is closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "happy path",
			distance: domain.Distance{
				Name:             "Marathon",
				DistanceInMeters: 42195,
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		if tt.expectedErr != nil {
			mock.ExpectExec(`CREATE TABLE IF NOT EXISTS distances ( name TEXT, distanceInMeters INT );`).
				WithoutArgs().
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectExec(`CREATE TABLE IF NOT EXISTS distances ( name TEXT, distanceInMeters INT );`).
				WithoutArgs().
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectPrepare(`INSERT INTO distances (name, distanceInMeters) VALUES (?, ?)`)
			mock.ExpectExec(`INSERT INTO distances (name, distanceInMeters) VALUES (?, ?)`).
				WithArgs(tt.distance.Name, tt.distance.DistanceInMeters).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err = mockDBAdapter.CreateDistance(tt.distance)
		if tt.expectedErr != nil {
			require.NotNil(t, err)
			assert.Equal(t, tt.expectedErr, err)
		} else {
			require.Nil(t, err)
		}
	}
}

func TestFindAllDistances(t *testing.T) {
	tests := []struct {
		name              string
		expectedDistances []domain.Distance
		expectedErr       error
	}{
		{
			name:              "no rows in database",
			expectedErr:       sql.ErrNoRows,
			expectedDistances: []domain.Distance{},
		},
		{
			name: "happy path",
			expectedDistances: []domain.Distance{
				{
					ID:               0,
					Name:             "Marathon",
					DistanceInMeters: 42195,
				},
				{
					ID:               1,
					Name:             "Half marathon",
					DistanceInMeters: 21098,
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
			mock.ExpectQuery(`SELECT rowid, name, distanceInMeters FROM distances`).
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(`SELECT rowid, name, distanceInMeters FROM distances`).
				WillReturnRows(sqlmock.NewRows([]string{"rowid", "name", "distanceInMeters"}).
					AddRow(0, "Marathon", 42195).
					AddRow(1, "Half marathon", 21098))
		}

		actualDistances, err := mockDBAdapter.FindAllDistances()
		assert.Equal(t, tt.expectedDistances, actualDistances)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestFindDistanceByID(t *testing.T) {
	tests := []struct {
		name             string
		DistanceID       int
		expectedDistance domain.Distance
		expectedErr      error
	}{
		{
			name:             "no rows in database",
			DistanceID:       0,
			expectedErr:      sql.ErrNoRows,
			expectedDistance: domain.Distance{},
		},
		{
			name:       "happy path",
			DistanceID: 0,
			expectedDistance: domain.Distance{
				ID:               0,
				Name:             "Marathon",
				DistanceInMeters: 42195,
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare("SELECT rowid, name, distanceInMeters FROM distances WHERE rowid = ?")
		if tt.expectedErr != nil {
			mock.ExpectQuery(`SELECT rowid, name, distanceInMeters FROM distances WHERE rowid = ?`).
				WithArgs(tt.DistanceID).WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(`SELECT rowid, name, distanceInMeters FROM distances WHERE rowid = ?`).
				WithArgs(tt.DistanceID).WillReturnRows(sqlmock.NewRows([]string{"rowid", "name", "distanceInMeters"}).
				AddRow(0, "Marathon", 42195).
				AddRow(1, "Half marathon", 21098))
		}

		actualDistance, err := mockDBAdapter.FindDistanceByID(tt.DistanceID)
		assert.Equal(t, tt.expectedDistance, actualDistance)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestDeleteDistance(t *testing.T) {
	tests := []struct {
		name        string
		distanceID  int
		expectedErr error
	}{
		{
			name:        "database closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name:       "happy path",
			distanceID: 3,
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare(`DELETE FROM distances WHERE rowid = ?`).
			WillReturnError(tt.expectedErr)
		if tt.expectedErr == nil {
			mock.ExpectExec(`DELETE FROM distances WHERE rowid = ?`).
				WithArgs(tt.distanceID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := mockDBAdapter.DeleteDistance(tt.distanceID)
		assert.Equal(t, tt.expectedErr, err)
	}
}
