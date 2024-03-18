package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const selectGoals string = `
	SELECT
	rowid,
	distance,
	time
	FROM goals
	`

func createGoalsTable(db *sql.DB) error {

	const createStmt string = `
	CREATE TABLE IF NOT EXISTS goals (
		distance TEXT,
		time TEXT
	);`

	_, err := db.Exec(createStmt)
	return err
}

func selectAllGoalsFromDB(db *sql.DB) ([]models.Goal, error) {

	var goals []models.Goal
	response, err := db.Query(selectGoals)
	if err != nil {
		return []models.Goal{}, err
	}

	for response.Next() {
		var goal models.Goal
		err = response.Scan(&goal.ID, &goal.Distance, &goal.Time)
		if err != nil {
			return []models.Goal{}, err
		}
		goals = append(goals, goal)
	}
	return goals, nil
}

func selectGoalByIdFromDB(db *sql.DB, id string) (models.Goal, error) {

	stmt, err := db.Prepare(selectGoals + " WHERE rowid = ?")
	if err != nil {
		return models.Goal{}, err
	}

	var goal models.Goal
	err = stmt.QueryRow(id).Scan(&goal.ID, &goal.Distance, &goal.Time)
	if err != nil {
		return models.Goal{}, err
	}
	return goal, nil
}

func insertGoalIntoDB(db *sql.DB, goal models.Goal) error {

	stmt, err := db.Prepare("INSERT INTO goals (distance, time) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(goal.Distance, goal.Time)
	return err
}

func updateGoalInDB(db *sql.DB, goal models.Goal, modifiedGoal models.Goal) error {

	stmt, err := db.Prepare("UPDATE goals SET distance = ?, time = ? WHERE rowid = ?")
	if err != nil {
		return err
	}

	if len(modifiedGoal.Distance) == 0 {
		modifiedGoal.Distance = goal.Distance
	}
	if len(modifiedGoal.Time) == 0 {
		modifiedGoal.Time = goal.Time
	}

	_, err = stmt.Exec(goal.Distance, goal.Time, goal.ID)
	return err
}

func deleteGoalFromDB(db *sql.DB, id string) error {

	stmt, err := db.Prepare("DELETE FROM goals WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
