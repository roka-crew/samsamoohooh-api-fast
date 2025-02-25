package validator

import (
	"github.com/roka-crew/pkg/apperr"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func New() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (v CustomValidator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return apperr.New("invalid request").
			SetStatus(http.StatusBadRequest).
			SetDetail(err.Error())
	}

	return nil
}
