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

const (
	createResultsTableStmt = `CREATE TABLE IF NOT EXISTS results (
		event_id INT,
		date TEXT,
		distance_id INT,
		time_gross_hours INT,
		time_gross_minutes INT,
		time_gross_seconds INT,
		time_net_hours INT,
		time_net_minutes INT,
		time_net_seconds INT,
		category TEXT,
		agegroup TEXT,
		place_total INT,
		place_category INT,
		place_agegroup INT,
		finisher_total INT,
		finisher_category INT,
		finisher_agegroup INT
	);`
	insertResultStmt = `INSERT INTO results (
		event_id,
		date,
		distance_id,
		time_gross_hours,
		time_gross_minutes,
		time_gross_seconds,
		time_net_hours,
		time_net_minutes,
		time_net_seconds,
		category,
		agegroup,
		place_total,
		place_category,
		place_agegroup,
		finisher_total,
		finisher_category,
		finisher_agegroup
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	selectResultStmt = `SELECT rowid, event_id, date, distance_id, time_gross_hours, time_gross_minutes, time_gross_seconds, time_net_hours, time_net_minutes, time_net_seconds, category, agegroup, place_total, place_category, place_agegroup, finisher_total, finisher_category, finisher_agegroup FROM results`
)

func TestCreateResult(t *testing.T) {
	tests := []struct {
		name        string
		Result      domain.Result
		expectedErr error
	}{
		{
			name:        "database is closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "happy path",
			Result: domain.Result{
				ResultID:   1,
				EventID:    2,
				Date:       "today",
				DistanceID: 3,
				TimeGross: domain.RaceTime{
					Hours:   4,
					Minutes: 5,
					Seconds: 6,
				},
				TimeNet: domain.RaceTime{
					Hours:   7,
					Minutes: 8,
					Seconds: 9,
				},
				Category: "the best category",
				Agegroup: "very old",
				Place: domain.CategoryNumbers{
					Total:    3,
					Category: 2,
					Agegroup: 1,
				},
				Finisher: domain.CategoryNumbers{
					Total:    6,
					Category: 5,
					Agegroup: 4,
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
			mock.ExpectExec(createResultsTableStmt).
				WithoutArgs().
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectExec(createResultsTableStmt).
				WithoutArgs().
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectPrepare(insertResultStmt)
			mock.ExpectExec(insertResultStmt).
				WithArgs(tt.Result.EventID, tt.Result.Date, tt.Result.DistanceID, tt.Result.TimeGross.Hours, tt.Result.TimeGross.Minutes, tt.Result.TimeGross.Seconds, tt.Result.TimeNet.Hours, tt.Result.TimeNet.Minutes, tt.Result.TimeNet.Seconds, tt.Result.Category, tt.Result.Agegroup, tt.Result.Place.Total, tt.Result.Place.Category, tt.Result.Place.Agegroup, tt.Result.Finisher.Total, tt.Result.Finisher.Category, tt.Result.Finisher.Agegroup).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err = mockDBAdapter.CreateResult(tt.Result)
		if tt.expectedErr != nil {
			require.NotNil(t, err)
			assert.Equal(t, tt.expectedErr, err)
		} else {
			require.Nil(t, err)
		}
	}
}

func TestFindAllResults(t *testing.T) {
	tests := []struct {
		name            string
		expectedResults []domain.Result
		expectedErr     error
	}{
		{
			name:            "no rows in database",
			expectedErr:     sql.ErrNoRows,
			expectedResults: []domain.Result{},
		},
		{
			name: "happy path",
			expectedResults: []domain.Result{
				{
					ResultID:   1,
					EventID:    2,
					Date:       "today",
					DistanceID: 3,
					TimeGross: domain.RaceTime{
						Hours:   4,
						Minutes: 5,
						Seconds: 6,
					},
					TimeNet: domain.RaceTime{
						Hours:   7,
						Minutes: 8,
						Seconds: 9,
					},
					Category: "the best category",
					Agegroup: "very old",
					Place: domain.CategoryNumbers{
						Total:    3,
						Category: 2,
						Agegroup: 1,
					},
					Finisher: domain.CategoryNumbers{
						Total:    6,
						Category: 5,
						Agegroup: 4,
					},
				},
				{
					ResultID:   2,
					EventID:    3,
					Date:       "yesterday",
					DistanceID: 4,
					TimeGross: domain.RaceTime{
						Hours:   4,
						Minutes: 5,
						Seconds: 6,
					},
					TimeNet: domain.RaceTime{
						Hours:   7,
						Minutes: 8,
						Seconds: 9,
					},
					Category: "the worst category",
					Agegroup: "very young",
					Place: domain.CategoryNumbers{
						Total:    3,
						Category: 2,
						Agegroup: 1,
					},
					Finisher: domain.CategoryNumbers{
						Total:    6,
						Category: 5,
						Agegroup: 4,
					},
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
			mock.ExpectQuery(selectResultStmt).
				WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(selectResultStmt).
				WillReturnRows(sqlmock.NewRows([]string{"rowid", "event_id", "date", "distance_id", "time_gross_hours",
					"time_gross_minutes", "time_gross_seconds", "time_net_hours",
					"time_net_minutes", "time_net_seconds", "category", "agegroup",
					"place_total", "place_category", "place_agegroup", "finisher_total", "finisher_category",
					"finisher_agegroup"}).
					AddRow(1, 2, "today", 3, 4, 5, 6, 7, 8, 9, "the best category", "very old", 3, 2, 1, 6, 5, 4).
					AddRow(2, 3, "yesterday", 4, 4, 5, 6, 7, 8, 9, "the worst category", "very young", 3, 2, 1, 6, 5, 4))
		}

		actualResults, err := mockDBAdapter.FindAllResults()
		assert.Equal(t, tt.expectedResults, actualResults)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestFindResultByID(t *testing.T) {
	tests := []struct {
		name           string
		ResultID       int
		expectedResult domain.Result
		expectedErr    error
	}{
		{
			name:           "no rows in database",
			ResultID:       0,
			expectedErr:    sql.ErrNoRows,
			expectedResult: domain.Result{},
		},
		{
			name:     "happy path",
			ResultID: 0,
			expectedResult: domain.Result{
				ResultID:   1,
				EventID:    2,
				Date:       "today",
				DistanceID: 3,
				TimeGross: domain.RaceTime{
					Hours:   4,
					Minutes: 5,
					Seconds: 6,
				},
				TimeNet: domain.RaceTime{
					Hours:   7,
					Minutes: 8,
					Seconds: 9,
				},
				Category: "the best category",
				Agegroup: "very old",
				Place: domain.CategoryNumbers{
					Total:    3,
					Category: 2,
					Agegroup: 1,
				},
				Finisher: domain.CategoryNumbers{
					Total:    6,
					Category: 5,
					Agegroup: 4,
				},
			},
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare(selectResultStmt + " WHERE rowid = ?")
		if tt.expectedErr != nil {
			mock.ExpectQuery(selectResultStmt + " WHERE rowid = ?").
				WithArgs(tt.ResultID).WillReturnError(tt.expectedErr)
		} else {
			mock.ExpectQuery(selectResultStmt + " WHERE rowid = ?").
				WithArgs(tt.ResultID).WillReturnRows(sqlmock.NewRows([]string{"rowid", "event_id", "date", "distance_id", "time_gross_hours",
				"time_gross_minutes", "time_gross_seconds", "time_net_hours",
				"time_net_minutes", "time_net_seconds", "category", "agegroup",
				"place_total", "place_category", "place_agegroup", "finisher_total", "finisher_category",
				"finisher_agegroup"}).
				AddRow(1, 2, "today", 3, 4, 5, 6, 7, 8, 9, "the best category", "very old", 3, 2, 1, 6, 5, 4))
		}

		actualResult, err := mockDBAdapter.FindResultByID(tt.ResultID)
		assert.Equal(t, tt.expectedResult, actualResult)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestDeleteResult(t *testing.T) {
	tests := []struct {
		name        string
		ResultID    int
		expectedErr error
	}{
		{
			name:        "database closed",
			expectedErr: sql.ErrConnDone,
		},
		{
			name:     "happy path",
			ResultID: 3,
		},
	}

	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.Nil(t, err)
	defer database.Close()

	mockDBAdapter := db.SQLDBAdapter{Database: database}

	for _, tt := range tests {
		mock.ExpectPrepare(`DELETE FROM results WHERE rowid = ?`).
			WillReturnError(tt.expectedErr)
		if tt.expectedErr == nil {
			mock.ExpectExec(`DELETE FROM results WHERE rowid = ?`).
				WithArgs(tt.ResultID).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := mockDBAdapter.DeleteResult(tt.ResultID)
		assert.Equal(t, tt.expectedErr, err)
	}
}
