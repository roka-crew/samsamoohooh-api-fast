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
