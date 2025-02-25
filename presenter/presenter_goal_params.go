package presenter

import (
	"time"

	"github.com/roka-crew/pkg/persistence"

	"github.com/roka-crew/domain"
)

type CreateGoalParams = domain.Goal

type ListGoalsParams struct {
	IDs           []uint
	GroupIDs      []uint
	Limit         int
	DeadlineOrder persistence.Order
	GteDeadline   time.Time
	WithTopics    bool
}

type PatchGoalParams struct {
	// condition
	ID uint

	// change
	Deadline  *time.Time
	PageRange *int
}

type DeleteGoalParams struct {
	ID uint
}
