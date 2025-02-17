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
