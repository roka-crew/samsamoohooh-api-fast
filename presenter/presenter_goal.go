package presenter

import "time"

type CreateGoalRequest struct {
	RequestUserID uint      `swaggerignore:"true"`
	GroupID       uint      `json:"groupID"`
	Deadline      time.Time `json:"deadline" validate:"required, datetime=2006-01-02" example:"2021-08-01"`
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
