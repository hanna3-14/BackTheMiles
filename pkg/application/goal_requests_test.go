package application_test

import (
	"errors"
	"testing"

	"github.com/hanna3-14/BackTheMiles/pkg/application"
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetGoals(t *testing.T) {
	mockDBAdapter := &MockGoalDBAdapter{}
	goalRequestService, _ := application.NewGoalRequestService(mockDBAdapter)

	goals := []domain.Goal{
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
	}

	mockDBAdapter.On("FindAllGoals").Return(goals, nil)

	actualGoals, err := goalRequestService.GetGoals()
	assert.Equal(t, goals, actualGoals)
	assert.Nil(t, err)
}

func TestGetGoal(t *testing.T) {
	mockDBAdapter := &MockGoalDBAdapter{}
	goalRequestService, _ := application.NewGoalRequestService(mockDBAdapter)

	goal := domain.Goal{
		ID:       0,
		Distance: "Marathon",
		Time:     "forever",
	}

	mockDBAdapter.On("FindGoalByID", mock.AnythingOfType("int")).Return(goal, nil)

	actualGoal, err := goalRequestService.GetGoal(0)
	assert.Equal(t, goal, actualGoal)
	assert.Nil(t, err)
}

func TestPostGoal(t *testing.T) {
	mockDBAdapter := &MockGoalDBAdapter{}
	goalRequestService, _ := application.NewGoalRequestService(mockDBAdapter)

	mockDBAdapter.On("CreateGoal", mock.AnythingOfType("domain.Goal")).Return(nil)

	goal := domain.Goal{
		ID:       0,
		Distance: "Marathon",
		Time:     "forever",
	}
	err := goalRequestService.PostGoal(goal)
	assert.Nil(t, err)
}

func TestPatchGoal(t *testing.T) {
	tests := []struct {
		name         string
		storedGoal   domain.Goal
		modifiedGoal domain.Goal
		expectedErr  error
	}{
		{
			name:       "goal could not be found",
			storedGoal: domain.Goal{},
			modifiedGoal: domain.Goal{
				ID:       0,
				Distance: "Half marathon",
				Time:     "very fast",
			},
			expectedErr: errors.New("goal not found"),
		},
		{
			name: "happy path",
			storedGoal: domain.Goal{
				ID:       0,
				Distance: "Marathon",
				Time:     "forever",
			},
			modifiedGoal: domain.Goal{
				ID:       0,
				Distance: "Half marathon",
				Time:     "very fast",
			},
		},
	}

	for _, tt := range tests {
		mockDBAdapter := &MockGoalDBAdapter{}
		goalRequestService, _ := application.NewGoalRequestService(mockDBAdapter)

		mockDBAdapter.On("FindGoalByID", mock.AnythingOfType("int")).Return(tt.storedGoal, tt.expectedErr)
		mockDBAdapter.On("UpdateGoal", mock.AnythingOfType("int"), mock.AnythingOfType("domain.Goal"), mock.AnythingOfType("domain.Goal")).Return(nil)

		err := goalRequestService.PatchGoal(tt.modifiedGoal.ID, tt.modifiedGoal)
		assert.Equal(t, tt.expectedErr, err)
	}
}

func TestDeleteGoal(t *testing.T) {
	mockDBAdapter := &MockGoalDBAdapter{}
	goalRequestService, _ := application.NewGoalRequestService(mockDBAdapter)

	mockDBAdapter.On("DeleteGoal", mock.AnythingOfType("int")).Return(nil)

	err := goalRequestService.DeleteGoal(0)
	assert.Nil(t, err)
}

type MockGoalDBAdapter struct {
	mock.Mock
}

func (m *MockGoalDBAdapter) CreateGoal(goal domain.Goal) error {
	args := m.Called(goal)
	return args.Error(0)
}

func (m *MockGoalDBAdapter) FindAllGoals() ([]domain.Goal, error) {
	args := m.Called()
	return args.Get(0).([]domain.Goal), args.Error(1)
}

func (m *MockGoalDBAdapter) FindGoalByID(id int) (domain.Goal, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Goal), args.Error(1)
}

func (m *MockGoalDBAdapter) UpdateGoal(id int, goal, modifiedGoal domain.Goal) error {
	args := m.Called(id, goal, modifiedGoal)
	return args.Error(0)
}

func (m *MockGoalDBAdapter) DeleteGoal(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
