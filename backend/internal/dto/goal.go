package dto

import (
	"github.com/google/uuid"
)

type CreateGoalRequest struct {
	Name                string
	TargetAmount        float64
	SavedAmount         float64
	MonthlyContribution float64
}

type UpdateGoalRequest struct {
	Name                string
	TargetAmount        float64
	SavedAmount         float64
	MonthlyContribution float64
}

type GoalResponse struct {
	ID                  uuid.UUID
	Name                string
	TargetAmount        float64
	SavedAmount         float64
	MonthlyContribution float64
}
