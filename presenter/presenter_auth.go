package presenter

type IssueTokenRequest struct {
	Nickname string `json:"nickname"`
}

type IssueTokenResponse struct {
	Token string `json:"token"`
}

type ValidateRequest struct {
	RequestUserID uint `swaggerignore:"true"`
}

type ValidateResponse struct {
	Nickname   string  `json:"nickname"`
	Resolution *string `json:"resolution,omitempty"`
}
