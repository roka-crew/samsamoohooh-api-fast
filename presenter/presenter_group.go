package presenter

type CreateGroupRequest struct {
	RequestUserID    uint    `swaggerignore:"true"`
	BookTitle        string  `json:"bookTitle"`
	BookAuthor       string  `json:"bookAuthor"`
	BookPageMax      int     `json:"bookPageMax"`
	BookPageCount    int     `json:"bookPageCount"`
	BookPublisher    *string `json:"bookPublisher"`
	BookIntroduction *string `json:"bookIntroduction"`
}

type CreateGroupResponse struct {
	BookTitle        string  `json:"bookTitle"`
	BookAuthor       string  `json:"bookAuthor"`
	BookPageMax      int     `json:"bookPageMax"`
	BookPageCount    int     `json:"bookPageCount"`
	BookPublisher    *string `json:"bookPublisher,omitempty"`
	BookIntroduction *string `json:"bookIntroduction,omitempty"`
}

type ListGroupsRequest struct {
	RequestUserID uint `swaggerignore:"true"`
}

type ListGroupsResponse struct {
	Groups []GroupsResponse `json:"groups"`
}

type GroupsResponse struct {
	ID               uint    `json:"id"`
	BookTitle        string  `json:"bookTitle"`
	BookAuthor       string  `json:"bookAuthor"`
	BookPageMax      int     `json:"bookPageMax"`
	BookPageCount    int     `json:"bookPageCount"`
	BookPublisher    *string `json:"bookPublisher,omitempty"`
	BookIntroduction *string `json:"bookIntroduction,omitempty"`
}

type PatchGroupRequest struct {
	RequestUserID    uint    `swaggerignore:"true"`
	GroupID          uint    `param:"group-id"`
	BookTitle        *string `json:"bookTitle" validate:"omitempty"`
	BookAuthor       *string `json:"bookAuthor" validate:"omitempty"`
	BookPageMax      *int    `json:"bookPageMax" validate:"omitempty"`
	BookPageCount    *int    `json:"bookPageCount" validate:"omitempty"`
	BookPublisher    *string `json:"bookPublisher,omitempty" validate:"omitempty"`
	BookIntroduction *string `json:"bookIntroduction,omitempty" validate:"omitempty"`
}

type DeleteGroupRequest struct {
	RequestUserID uint `swaggerignore:"true"`
	GroupID       uint `param:"group-id"`
}

type LeaveGroupRequest struct {
	RequestUserID uint `swaggerignore:"true"`
	GroupID       uint `param:"group-id"`
}

type JoinGroupRequest struct {
	RequestUserID uint `swaggerignore:"true"`
	GroupID       uint `param:"group-id"`
}
