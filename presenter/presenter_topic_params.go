package presenter

import "github.com/roka-crew/domain"

type CreateTopicParams = domain.Topic

type ListTopicsParams struct {
	IDs     []uint
	GoalIDs []uint
	UserIDs []uint
	Limit   int
}

type PatchTopicParams struct {
	// condition
	ID uint

	// change
	Title   *string
	Content *string
}

type DeleteTopicParams struct {
	ID uint
}
