package repository

import (
	"github.com/PickHD/bookstore-api/src/apps/books/model"
	"github.com/PickHD/bookstore-api/src/utils"
	"gorm.io/gorm"
)

type BookQueryRepository struct {
	DB *gorm.DB
}

func NewBookQueryRepository(database *gorm.DB) BookQueryRepository {
	return BookQueryRepository{DB: database}
}

func (repository BookQueryRepository) GetAll(pg *utils.Pagination) ([]model.Books, int64, error) {
	var getAllBooks []model.Books

	offset := (pg.Page - 1) * pg.Limit

	result := repository.DB.Model(&model.Books{}).Limit(pg.Limit).Offset(offset).Find(&getAllBooks)
	if result.Error != nil {
		return []model.Books{}, 0, result.Error
	}

	return getAllBooks, result.RowsAffected, nil
}

func (repository BookQueryRepository) GetOne(bookUUID string) (model.Books, error) {
	getOneBook := model.NewBooks()

	result := repository.DB.Model(&model.Books{}).Where("uuid=?", bookUUID).First(&getOneBook)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return getOneBook, nil
}
