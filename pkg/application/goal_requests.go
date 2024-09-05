package application

import "github.com/hanna3-14/BackTheMiles/pkg/domain"

type IGoalRepository interface {
	CreateGoal(goal domain.Goal) error
	FindAllGoals() ([]domain.Goal, error)
	FindGoalByID(id int) (domain.Goal, error)
	UpdateGoal(id int, goal, modifiedGoal domain.Goal) error
	DeleteGoal(id int) error
}

type GoalRequestService struct {
	repo IGoalRepository
}

func NewGoalRequestService(goalRepo IGoalRepository) (GoalRequestService, error) {
	return GoalRequestService{repo: goalRepo}, nil
}

func (r *GoalRequestService) GetGoals() ([]domain.Goal, error) {
	return r.repo.FindAllGoals()
}

func (r *GoalRequestService) GetGoal(id int) (domain.Goal, error) {
	return r.repo.FindGoalByID(id)
}

func (r *GoalRequestService) PostGoal(goal domain.Goal) error {
	return r.repo.CreateGoal(goal)
}

func (r *GoalRequestService) PatchGoal(id int, modifiedGoal domain.Goal) error {
	storedGoal, err := r.repo.FindGoalByID(id)
	if err != nil {
		return err
	}
	return r.repo.UpdateGoal(id, storedGoal, modifiedGoal)
}

func (r *GoalRequestService) DeleteGoal(id int) error {
	return r.repo.DeleteGoal(id)
}
