package domain

type IGoalRepository interface {
	CreateGoal(goal Goal) error
	FindAllGoals() ([]Goal, error)
	FindGoalByID(id int) (Goal, error)
	UpdateGoal(id int, goal, modifiedGoal Goal) error
	DeleteGoal(id int) error
}
