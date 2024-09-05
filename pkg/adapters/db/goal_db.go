package db

import (
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
)

const selectGoals string = `
	SELECT
	rowid,
	distance,
	time
	FROM goals
	`

func (repo *SQLDBAdapter) CreateGoal(goal domain.Goal) error {
	const createStmt string = `
		CREATE TABLE IF NOT EXISTS goals (
			distance TEXT,
			time TEXT
		);`

	_, err := repo.Database.Exec(createStmt)
	if err != nil {
		return err
	}

	stmt, err := repo.Database.Prepare("INSERT INTO goals (distance, time) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(goal.Distance, goal.Time)
	return err
}

func (repo *SQLDBAdapter) FindAllGoals() ([]domain.Goal, error) {
	var goals []domain.Goal
	response, err := repo.Database.Query(selectGoals)
	if err != nil {
		return []domain.Goal{}, err
	}

	for response.Next() {
		var goal domain.Goal
		err = response.Scan(&goal.ID, &goal.Distance, &goal.Time)
		if err != nil {
			return []domain.Goal{}, err
		}
		goals = append(goals, goal)
	}
	return goals, nil
}

func (repo *SQLDBAdapter) FindGoalByID(id int) (domain.Goal, error) {
	stmt, err := repo.Database.Prepare(selectGoals + " WHERE rowid = ?")
	if err != nil {
		return domain.Goal{}, err
	}

	var goal domain.Goal
	err = stmt.QueryRow(id).Scan(&goal.ID, &goal.Distance, &goal.Time)
	if err != nil {
		return domain.Goal{}, err
	}
	return goal, nil
}

func (repo *SQLDBAdapter) UpdateGoal(id int, goal, modifiedGoal domain.Goal) error {
	stmt, err := repo.Database.Prepare("UPDATE goals SET distance = ?, time = ? WHERE rowid = ?")
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

func (repo *SQLDBAdapter) DeleteGoal(id int) error {
	stmt, err := repo.Database.Prepare("DELETE FROM goals WHERE rowid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
