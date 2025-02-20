package presenter

import "time"

type CreateGoalRequest struct {
	RequestUserID uint      `swaggerignore:"true"`
	GroupID       uint      `json:"group_id"`
	Deadline      time.Time `json:"deadline" validate:"required, datetime=2006-01-02"`
	PageRange     int       `json:"page_range"`
}

type CreateGoalResponse struct {
}
