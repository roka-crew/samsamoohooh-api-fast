package service

import (
	"context"

	"github.com/roka-crew/domain"
	goalStore "github.com/roka-crew/internal/goal/store"
	groupStore "github.com/roka-crew/internal/group/store"
	"github.com/roka-crew/presenter"
)

type GoalService struct {
	groupStore *groupStore.GroupStore
	goalStore  *goalStore.GoalStore
}

func NewGoalService(
	goalStore *goalStore.GoalStore,
	groupStore *groupStore.GroupStore,
) *GoalService {
	return &GoalService{
		groupStore: groupStore,
		goalStore:  goalStore,
	}
}

func (s GoalService) CreateGoal(ctx context.Context, request presenter.CreateGoalRequest) (presenter.CreateGoalResponse, error) {
	// (1) 생성하고자 하는 사용자가 속한 group인지 확인
	exists, err := s.groupStore.ExistsGroupUser(ctx, presenter.ExistsGroupUserParams{
		UserID:  request.RequestUserID,
		GroupID: request.GroupID,
	})
	if err != nil {
		return presenter.CreateGoalResponse{}, err
	}
	if !exists {
		return presenter.CreateGoalResponse{}, domain.ErrGroupNotMember
	}

	// TODO: 로직 어떻게 변경할 건지 생각하기
	// (2) Deadline 보다 이상인 Goal이 있는지 확인
	// 목표는 진행중인 목표가 없어야지만 가능하다.
	goals, err := s.goalStore.ListGoals(ctx, presenter.ListGoalsParams{
		GroupIDs:    []uint{request.GroupID},
		GteDeadline: request.Deadline,
	})
	if err != nil {
		return presenter.CreateGoalResponse{}, err
	}

	if len(goals) > 0 {
		return presenter.CreateGoalResponse{}, domain.ErrFutureGoalExists
	}

	return presenter.CreateGoalResponse{}, nil
}
