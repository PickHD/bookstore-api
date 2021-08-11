package validator

import (
	"github.com/PickHD/bookstore-api/src/utils"
	"github.com/go-playground/validator/v10"
)

type BookValidator struct {
	Title           string `json:"title" form:"title" validate:"required,min=10,max=255"`
	ISBN            string `json:"isbn" form:"isbn" validate:"required,min=12,max=12"`
	Year            string `json:"year" form:"year" validate:"required,min=4,max=4"`
	Price           int    `json:"price" form:"price" validate:"required,min=7"`
	NoPages         int    `json:"no_pages" form:"no_pages" validate:"required,min=10,max=255"`
	BookDescription string `json:"book_description" form:"book_description" validate:"required,min=10,max=255"`
}

func NewBookValidator() BookValidator {
	return BookValidator{}
}

func (book BookValidator) Validate() []*utils.ErrorValidateResponse {
	var errors []*utils.ErrorValidateResponse
	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element utils.ErrorValidateResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
