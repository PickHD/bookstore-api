package controller

import (
	"github.com/PickHD/bookstore-api/src/apps/books/model"
	"github.com/PickHD/bookstore-api/src/apps/books/repository"
	"github.com/PickHD/bookstore-api/src/apps/books/validator"
	"github.com/PickHD/bookstore-api/src/utils"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(database *gorm.DB) BookController {
	return BookController{DB: database}
}

func (controller BookController) CreateBook(data validator.BookValidator) (insertedUUID string, err error) {
	newBook := model.NewBooks()

	newBook.Title = data.Title
	newBook.ISBN = data.ISBN
	newBook.Year = data.Year
	newBook.Price = data.Price
	newBook.NoPages = data.NoPages
	newBook.BookDescription = data.BookDescription

	insertedUUID, err = repository.NewBookCommandRepository(controller.DB).Save(newBook)
	if err != nil {
		return "", err
	}

	return insertedUUID, nil
}

func (controller BookController) GetAllBooks(pg *utils.Pagination) ([]model.Books, int64, error) {
	getAllBooks, total, err := repository.NewBookQueryRepository(controller.DB).GetAll(pg)
	if err != nil {
		return []model.Books{}, 0, err
	}

	return getAllBooks, total, nil
}

func (controller BookController) GetOneBooks(bookUUID string) (model.Books, error) {
	getBook, err := repository.NewBookQueryRepository(controller.DB).GetOne(bookUUID)
	if err != nil {
		return model.Books{}, err
	}

	return getBook, nil
}

func (controller BookController) UpdateBook(data validator.BookValidator, bookUUID string) (updatedUUID string, err error) {
	updBook := model.NewBooks()

	updBook.Title = data.Title
	updBook.ISBN = data.ISBN
	updBook.Year = data.Year
	updBook.Price = data.Price
	updBook.NoPages = data.NoPages
	updBook.BookDescription = data.BookDescription

	updatedUUID, err = repository.NewBookCommandRepository(controller.DB).Update(updBook, bookUUID)
	if err != nil {
		return "", err
	}

	return updatedUUID, nil
}

func (controller BookController) DeleteBook(bookUUID string) (deletedUUID string, err error) {
	deletedUUID, err = repository.NewBookCommandRepository(controller.DB).Delete(bookUUID)
	if err != nil {
		return "", err
	}

	return deletedUUID, nil
}
