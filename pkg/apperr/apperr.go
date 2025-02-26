package apperr

import (
	stderrors "errors"
	"fmt"
	"regexp"
	"strings"

	pkgerrors "github.com/pkg/errors"
)

type Error struct {
	// 오류의 짧은 제목
	Title string `json:"title"`
	// HTTP 상태 코드
	Status int `json:"status,omitempty"`
	// 오류의 상세 설명
	Detail string `json:"detail,omitempty"`
	// 오류가 발생한 특정 인스턴스를 식별하는 데 사용할 수 있는 URI입니다.
	Instance string `json:"instance,omitempty"`
}

func New(title string) *Error {
	return &Error{
		Title: title,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Detail)
}

func (e *Error) SetStatus(status int) *Error {
	e.Status = status
	return e
}

func (e *Error) SetDetail(detail string) *Error {
	e.Detail = detail
	return e
}

func (e *Error) SetInstance(instance string) *Error {
	e.Instance = instance
	return e
}

func Restore(err error) *Error {
	if err == nil {
		return &Error{}
	}

	var appErr *Error
	if stderrors.As(err, &appErr) {
		return appErr
	}

	return New(err.Error())
}

func Is(err error, target error) bool {
	return stderrors.Is(err, target)
}

type InternalError struct {
	Err error
}

func NewInternalError(err error) error {
	return &InternalError{Err: pkgerrors.WithStack(err)}
}

func (e *InternalError) Error() string {
	return e.Err.Error()
}

func (e *InternalError) StackTrace() string {
	if e.Err == nil {
		return "no stack trace available"
	}

	return fmt.Sprintf("%+v", e.Err)
}

func (e *InternalError) Pretty() string {
	stackTrace := e.StackTrace()

	// 정규식을 사용하여 함수 호출과 파일 위치를 함께 매칭
	re := regexp.MustCompile(`(?m)^(.*)\n\s+(.*:\d+)`)
	formatted := re.ReplaceAllString(stackTrace, "🔹 $1\n        $2")

	// 몇 가지 추가 정리
	lines := strings.Split(formatted, "\n")
	var result []string

	result = append(result, "❌ 에러 메시지: "+e.Error())
	result = append(result, "📋 스택 트레이스:")

	inStackTrace := false
	for _, line := range lines {
		if strings.HasPrefix(line, "🔹") {
			inStackTrace = true
			result = append(result, "    "+line) // 함수 호출에 들여쓰기 적용
		} else if strings.HasPrefix(line, "        ") && inStackTrace {
			result = append(result, "    "+line) // 파일 위치에 더 많은 들여쓰기 적용
		} else if len(line) > 0 && !strings.HasPrefix(line, e.Error()) && inStackTrace {
			result = append(result, "        ℹ️ "+line) // 기타 정보에 추가 들여쓰기
		}
	}

	return strings.Join(result, "\n")
}
