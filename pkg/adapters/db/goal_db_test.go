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

func TestCreateGoal(t *testing.T) {
	tests := []struct {
		name        string
		goal        domain.Goal
		expectedErr error
	}{
		{
			name:        "database is closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "happy path",
			goal: domain.Goal{
				Distance: "Marathon",
				Time:     "forever",
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		if tt.expectedErr != nil {
			mock.ExpectExec(`CREATE TABLE IF NOT EXISTS goals ( distance TEXT, time TEXT );`).
				WithoutArgs().
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectExec(`CREATE TABLE IF NOT EXISTS goals ( distance TEXT, time TEXT );`).
				WithoutArgs().
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectPrepare(`INSERT INTO goals (distance, time) VALUES (?, ?)`)
			mock.ExpectExec(`INSERT INTO goals (distance, time) VALUES (?, ?)`).
				WithArgs(tt.goal.Distance, tt.goal.Time).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err = mockDBAdapter.CreateGoal(tt.goal)
		if tt.expectedErr != nil {
			require.NotNil(t, err)
			assert.Equal(t, tt.expectedErr, err)
		} else {
			require.Nil(t, err)
		}
	}
}

func TestFindAllGoals(t *testing.T) {
	tests := []struct {
		name          string
		expectedGoals []domain.Goal
		expectedErr   error
	}{
		{
			name:          "no rows in database",
			expectedErr:   sql.ErrNoRows,
			expectedGoals: []domain.Goal{},
		},
		{
			name: "happy path",
			expectedGoals: []domain.Goal{
				{
					ID:       0,
					Distance: "Marathon",
					Time:     "forever",
				},
				{
					ID:       1,
					Distance: "Half marathon",
					Time:     "very fast",
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
			mock.ExpectQuery(`SELECT rowid, distance, time FROM goals`).
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(`SELECT rowid, distance, time FROM goals`).
				WillReturnRows(sqlmock.NewRows([]string{"rowid", "distance", "time"}).
					AddRow(0, "Marathon", "forever").
					AddRow(1, "Half marathon", "very fast"))
		}

		actualGoals, err := mockDBAdapter.FindAllGoals()
		assert.Equal(t, tt.expectedGoals, actualGoals)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestFindGoalByID(t *testing.T) {
	tests := []struct {
		name         string
		goalID       int
		expectedGoal domain.Goal
		expectedErr  error
	}{
		{
			name:         "no rows in database",
			goalID:       0,
			expectedErr:  sql.ErrNoRows,
			expectedGoal: domain.Goal{},
		},
		{
			name:   "happy path",
			goalID: 0,
			expectedGoal: domain.Goal{
				ID:       0,
				Distance: "Marathon",
				Time:     "forever",
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare("SELECT rowid, distance, time FROM goals WHERE rowid = ?")
		if tt.expectedErr != nil {
			mock.ExpectQuery(`SELECT rowid, distance, time FROM goals WHERE rowid = ?`).
				WithArgs(tt.goalID).WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(`SELECT rowid, distance, time FROM goals WHERE rowid = ?`).
				WithArgs(tt.goalID).WillReturnRows(sqlmock.NewRows([]string{"rowid", "distance", "time"}).
				AddRow(0, "Marathon", "forever").
				AddRow(1, "Half marathon", "very fast"))
		}

		actualGoal, err := mockDBAdapter.FindGoalByID(tt.goalID)
		assert.Equal(t, tt.expectedGoal, actualGoal)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestUpdateGoal(t *testing.T) {
	tests := []struct {
		name         string
		storedGoal   domain.Goal
		modifiedGoal domain.Goal
	}{
		{
			name: "happy path - update distance",
			storedGoal: domain.Goal{
				ID:       1,
				Distance: "Marathon",
				Time:     "forever",
			},
			modifiedGoal: domain.Goal{
				Distance: "Half marathon",
			},
		},
		{
			name: "happy path - update time",
			storedGoal: domain.Goal{
				ID:       1,
				Distance: "Marathon",
				Time:     "forever",
			},
			modifiedGoal: domain.Goal{
				Time: "very fast",
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare("UPDATE goals SET distance = ?, time = ? WHERE rowid = ?")
		mock.ExpectExec(`UPDATE goals SET distance = ?, time = ? WHERE rowid = ?`).
			WithArgs(tt.storedGoal.Distance, tt.storedGoal.Time, tt.storedGoal.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := mockDBAdapter.UpdateGoal(tt.storedGoal.ID, tt.storedGoal, tt.modifiedGoal)
		assert.Nil(t, err)
	}
}

func TestDeleteGoal(t *testing.T) {
	tests := []struct {
		name        string
		goalID      int
		expectedErr error
	}{
		{
			name:        "database closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name:   "happy path",
			goalID: 3,
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare(`DELETE FROM goals WHERE rowid = ?`).
			WillReturnError(tt.expectedErr)
		if tt.expectedErr == nil {
			mock.ExpectExec(`DELETE FROM goals WHERE rowid = ?`).
				WithArgs(tt.goalID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := mockDBAdapter.DeleteGoal(tt.goalID)
		assert.Equal(t, tt.expectedErr, err)
	}
}
