package presenter

import "github.com/roka-crew/domain"

type CreateUserRequest struct {
	Nickname   string  `json:"nickname" validate:"required"`
	Resolution *string `json:"resolution" validate:"omitempty"`
}

func NewCreateUserResponse(user *domain.User) CreateUserResponse {
	return CreateUserResponse{
		Nickname:   user.Nickname,
		Resolution: user.Resolution,
	}
}

type CreateUserResponse struct {
	Nickname   string  `json:"nickname"`
	Resolution *string `json:"resolution,omitempty"`
}

type FindUserByMeRequest struct {
	RequestUserID uint
}

func NewFindUserByMeRequest(user *domain.User) FindUserByMeResponse {
	return FindUserByMeResponse{
		Nickname:   user.Nickname,
		Resolution: user.Resolution,
	}
}

type FindUserByMeResponse struct {
	Nickname   string  `json:"nickname"`
	Resolution *string `json:"resolution,omitempty"`
}
