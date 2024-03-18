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
	defer db.Close()

	return selectAllGoalsFromDB(db)
}

func GetGoalById(id string) (models.Goal, error) {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return models.Goal{}, err
	}
	defer db.Close()

	return selectGoalByIdFromDB(db, id)
}

func PostGoal(goal models.Goal) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	err = createGoalsTable(db)
	if err != nil {
		return err
	}

	return insertGoalIntoDB(db, goal)
}

func PatchGoal(id string, modifiedGoal models.Goal) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	goal, err := selectGoalByIdFromDB(db, id)
	if err != nil {
		return err
	}

	return updateGoalInDB(db, goal, modifiedGoal)
}

func DeleteGoal(id string) error {

	path := helpers.SafeGetEnv("PATH_TO_VOLUME")

	db, err := sql.Open("sqlite3", path+goalsFile)
	if err != nil {
		return err
	}
	defer db.Close()

	return deleteGoalFromDB(db, id)
}
