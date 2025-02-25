package presenter

import "time"

type CreateGoalRequest struct {
	RequestUserID uint      `swaggerignore:"true"`
	GroupID       uint      `json:"groupID"`
	Deadline      time.Time `json:"deadline" validate:"required" example:"2021-08-01T00:00:00Z"`
	PageRange     int       `json:"pageRange"`
}

type CreateGoalResponse struct {
	GoalID    uint      `json:"goalID"`
	Deadline  time.Time `json:"deadline"`
	PageRange int       `json:"pageRange"`
}

type ListGoalsRequest struct {
	RequestUserID uint `swaggerignore:"true"`
	GroupID       uint `json:"groupID"`
}

type ListGoalsResponse struct {
	Goals []GoalResponse `json:"goals"`
}

type GoalResponse struct {
	ID        uint      `json:"id"`
	Deadline  time.Time `json:"deadline"`
	PageRange int       `json:"pageRange"`
}

type PatchGoalRequest struct {
	RequestUserID uint       `swaggerignore:"true"`
	GoalID        uint       `param:"goal-id"`
	Deadline      *time.Time `json:"deadline,omitempty" validate:"omitempty" example:"2021-08-01T00:00:00Z"`
	PageRange     *int       `json:"pageRange,omitempty" validate:"omitempty"`
}

type DeleteGoalRequest struct {
	RequestUserID uint `swaggerignore:"true"`
	GoalID        uint `param:"goal-id"`
}
