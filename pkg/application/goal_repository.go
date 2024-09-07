package application

import "github.com/hanna3-14/BackTheMiles/pkg/domain"

type IGoalRepository interface {
	CreateGoal(goal domain.Goal) error
	FindAllGoals() ([]domain.Goal, error)
	FindGoalByID(id int) (domain.Goal, error)
	UpdateGoal(id int, goal, modifiedGoal domain.Goal) error
	DeleteGoal(id int) error
}
