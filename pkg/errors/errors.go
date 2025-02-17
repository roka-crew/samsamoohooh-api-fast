package errors

import (
	stderrors "errors"
	"fmt"
	"net/http"
)

type Error struct {
	// 개발자가 읽을 수 있는 오류 메시지입니다. 보안 유출을 방지하기 위해 사용자에게 표시되지 않습니다.
	Err error `json:"-" xml:"-" yaml:"-"`
	// 오류의 짧은 제목
	Title string `json:"title"`
	// HTTP 상태 코드
	Status int `json:"status,omitempty"`
	// 오류의 상세 설명
	Detail string `json:"detail,omitempty"`
}

func New(message string) Error {
	return Error{
		Err: stderrors.New(message),
	}
}

func NewInternalError(err error) Error {
	return Error{
		Err:    err,
		Status: http.StatusInternalServerError,
	}
}

func (e Error) SetStatus(status int) Error {
	e.Status = status
	return e
}

func (e Error) SetTitle(title string) Error {
	e.Title = title
	return e
}

func (e Error) SetDetail(detail string) Error {
	e.Detail = detail
	return e
}

func (e Error) StatusCode() int {
	if e.Status == 0 {
		return http.StatusInternalServerError
	}
	return e.Status
}

func (e Error) DetailMsg() string {
	return e.Detail
}

func (e Error) Error() string {
	code := e.StatusCode()
	title := e.Title
	if title == "" {
		title = http.StatusText(code)
		if title == "" {
			title = "HTTP Error"
		}
	}
	msg := fmt.Sprintf("%d %s", code, title)

	detail := e.DetailMsg()
	if detail == "" {
		return msg
	}

	return fmt.Sprintf("%s: %s", msg, detail)
}

func (e Error) Unwrap() error {
	return e.Err
}

func Is(err error, target error) bool {
	return stderrors.Is(err, target)
}

func As(err error, target any) bool {
	return stderrors.As(err, target)
}
