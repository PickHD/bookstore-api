package model

import (
	"github.com/PickHD/bookstore-api/src/utils"
	"gorm.io/gorm"
)

type Books struct {
	utils.BaseModel
	Title           string
	ISBN            string `gorm:"type:varchar(12);"`
	Year            string `gorm:"type:varchar(4);"`
	Price           int
	NoPages         int
	BookDescription string
}

func NewBooks() Books {
	return Books{}
}

func (model *Books) BeforeCreate(tx *gorm.DB) (err error) {
	model.UUID = utils.GenerateUuid()

	if !utils.ValidateUuid(model.UUID) {
		return err
	}

	return nil
}
