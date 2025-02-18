package presenter

type CreateUserRequest struct {
	Nickname   string  `json:"nickname" validate:"required"`
	Resolution *string `json:"resolution" validate:"omitempty"`
}

type CreateUserResponse struct {
	Nickname   string  `json:"nickname"`
	Resolution *string `json:"resolution,omitempty"`
}

type FindUserByMeRequest struct {
	RequestUserID uint `swaggerignore:"true"`
}

type FindUserByMeResponse struct {
	Nickname   string  `json:"nickname"`
	Resolution *string `json:"resolution,omitempty"`
}

type PatchUserByMeRequest struct {
	RequestUserID uint    `swaggerignore:"true"`
	Nickname      *string `json:"nickname" validate:"omitempty"`
	Resolution    *string `json:"resolution" validate:"omitempty"`
}

type DeleteUserByMeRequest struct {
	RequestUserID uint `swaggerignore:"true"`
}
