package presenter

import (
	"time"

	"github.com/roka-crew/domain"
)

type CreateGoalParams = domain.Goal

type ListGoalsParams struct {
	IDs         []uint
	GroupIDs    []uint
	Limit       int
	GteDeadline time.Time
	WithTopcis  bool
}

type PatchGoalParams struct {
	// condition
	GoalID uint

	// change
	Deadline  *time.Time
	PageRange *int
}

type DeleteGoalParams struct {
	GoalID uint
}
