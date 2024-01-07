package data

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const goalsFile = "/goals.db"

func GetGoals() ([]models.Goal, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return []models.Goal{}, err
	}

	var goals []models.Goal
	response, err := db.Query("SELECT rowid, distance, time FROM goals")
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

func GetGoalById(id string) (models.Goal, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return models.Goal{}, err
	}

	stmt, err := db.Prepare("SELECT rowid, distance, time FROM goals WHERE rowid = ?")
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

func PostGoal(goal models.Goal) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return err
	}

	const create string = `
	CREATE TABLE IF NOT EXISTS goals (
		distance TEXT,
		time TEXT
	);`

	_, err = db.Exec(create)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO goals(distance, time) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goal.Distance, goal.Time)
	if err != nil {
		return err
	}

	return nil
}

func PatchGoal(id string, modifiedGoal models.Goal) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("SELECT rowid, distance, time FROM goals WHERE rowid = ?")
	if err != nil {
		return err
	}

	var goal models.Goal
	err = stmt.QueryRow(id).Scan(&goal.ID, &goal.Distance, &goal.Time)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("UPDATE goals SET distance = ?, time = ? WHERE rowid = ?")
	if err != nil {
		return err
	}

	if len(modifiedGoal.Distance) == 0 {
		modifiedGoal.Distance = goal.Distance
	}
	if len(modifiedGoal.Time) == 0 {
		modifiedGoal.Time = goal.Time
	}

	_, err = stmt.Exec(modifiedGoal.Distance, modifiedGoal.Time, modifiedGoal.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteGoal(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM goals WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
