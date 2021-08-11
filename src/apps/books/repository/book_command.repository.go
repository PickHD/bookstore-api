package repository

import (
	"github.com/PickHD/bookstore-api/src/apps/books/model"
	"gorm.io/gorm"
)

type BookCommandRepository struct {
	DB *gorm.DB
}

func NewBookCommandRepository(database *gorm.DB) BookCommandRepository {
	return BookCommandRepository{DB: database}
}

func (repository BookCommandRepository) Save(books model.Books) (insertedUUID string, err error) {
	result := repository.DB.Create(&books)
	if result.Error != nil {
		return "", result.Error
	}

	insertedUUID = books.UUID
	return
}

func (repository BookCommandRepository) Update(books model.Books, bookUUID string) (updatedUUID string, err error) {
	result := repository.DB.Model(&model.Books{}).Where("uuid=?", bookUUID).Updates(&model.Books{
		Title:           books.Title,
		ISBN:            books.ISBN,
		Year:            books.Year,
		Price:           books.Price,
		NoPages:         books.NoPages,
		BookDescription: books.BookDescription,
	})

	if result.Error != nil {
		return "", result.Error
	}

	updatedUUID = bookUUID
	return
}

func (repository BookCommandRepository) Delete(bookUUID string) (deletedUUID string, err error) {
	result := repository.DB.Model(&model.Books{}).Where("uuid=?", bookUUID).Delete(&model.Books{})
	if result.Error != nil {
		return "", result.Error
	}

	deletedUUID = bookUUID
	return
}
