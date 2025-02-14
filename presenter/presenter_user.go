package presenter

import "github.com/roka-crew/domain"

type CreateUserRequest struct {
	Nickname string
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
