package service

import (
	"context"
	"github.com/roka-crew/pkg/persistence"
	"time"

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

	now := time.Now()
	// (2) 데드라인 유효성 검사
	// 0. 아래 조건들을 만족해야지만 목표를 생성할 수 있다.
	// 1. 지금 시각 기준으로, 요청자의 구룹 목표의 최신 데드라인 과거어야 한다.
	// 2. 지금 시각 기준으로, 생성하고자 하는 데드라인은 미래어야 한다.
	// 과거 -- (GroupLastGoalDeadline) -- (now) -- (request.Deadline) -- > 미래

	// (2-1) 현재 구룹의 최신 목표를 가져온다.
	listGoals, err := s.goalStore.ListGoals(ctx, presenter.ListGoalsParams{
		GroupIDs:      []uint{request.GroupID},
		DeadlineOrder: persistence.OrderDESC,
		Limit:         1,
	})
	if err != nil {
		return presenter.CreateGoalResponse{}, err
	}

	// (2-2) 만약 최신 목표가 있다면, 데드라인을 비교한다.
	if !listGoals.IsEmpty() {
		// (2-2-1) 유효성 조건 1 참고
		// 현재 시간 기준으로, 최신 목표의 데드라인이 미래에 위치해 있다면
		if listGoals.First().Deadline.After(now) {
			return presenter.CreateGoalResponse{}, domain.ErrInvalidDeadline.SetDetail("현재 시각 기준으로, 최신 목표의 데드라인은 과거여야 합니다.")
		}
	}

	// (2-3) 유효성 조건 2 참고
	// 현재 시간 기준으로, 생성하고자 하는 데드라인이 과거에 위치해 있다면
	if request.Deadline.After(now) {
		return presenter.CreateGoalResponse{}, domain.ErrInvalidDeadline.SetDetail("현재 시각 기준으로, 생성하고자 하는 데드라인은 미래여야 합니다.")
	}

	// (3) 목표 생성
	createdGoal, err := s.goalStore.CreateGoal(ctx, presenter.CreateGoalParams{
		GroupID:   request.GroupID,
		Deadline:  request.Deadline,
		PageRange: request.PageRange,
	})
	if err != nil {
		return presenter.CreateGoalResponse{}, err
	}

	return presenter.CreateGoalResponse{
		GoalID:    createdGoal.ID,
		Deadline:  createdGoal.Deadline,
		PageRange: createdGoal.PageRange,
		GroupID:   createdGoal.GroupID,
	}, nil
}
