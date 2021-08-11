package handler

import (
	"errors"
	"net/http"

	"github.com/PickHD/bookstore-api/src/apps/books/controller"
	"github.com/PickHD/bookstore-api/src/apps/books/presenter"
	"github.com/PickHD/bookstore-api/src/apps/books/validator"
	"github.com/PickHD/bookstore-api/src/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookHandler struct {
	DB *gorm.DB
}

func NewBookHandler(database *gorm.DB) BookHandler {
	return BookHandler{DB: database}
}

func (handler BookHandler) CreateBookHandler(c *fiber.Ctx) error {
	bookValidator := validator.NewBookValidator()

	if err := c.BodyParser(&bookValidator); err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Failed Body Parsing", err, fiber.Map{"error": map[string]string{"body_parser_error": err.Error()}})
	}

	errValidate := bookValidator.Validate()
	if errValidate != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Invalid Form", nil, fiber.Map{"error": errValidate})
	}

	tx := handler.DB.Begin()

	insertedUUID, err := controller.NewBookController(tx).CreateBook(bookValidator)
	if err != nil {
		tx.Rollback()
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Insert Data", err, fiber.Map{"error": map[string]string{"create_book_error": err.Error()}})
	}

	tx.Commit()
	return utils.NewResponse(c).ResponseFormatter(fiber.StatusCreated, "Create Book Successfully", nil, fiber.Map{"insertedUUID": insertedUUID})
}

func (handler BookHandler) GetAllBooksHandler(c *fiber.Ctx) error {
	getLimit := c.Query("limit", "5")
	getPage := c.Query("page", "1")

	pg, err := utils.GeneratePaginator(getLimit, getPage)
	if err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Generate Pagination", err, fiber.Map{"error": map[string]string{"generate_pagination_error": err.Error()}})
	}

	getBookModel, total, err := controller.NewBookController(handler.DB).GetAllBooks(&pg)
	if err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Fetch Data", err, fiber.Map{"error": map[string]string{"get_all_books_error": err.Error()}})
	}

	bookListPresenter := presenter.NewBookListPresenter()
	getAllBooks := bookListPresenter.Build(getBookModel, total)

	countTotalPage := utils.CountTotalPage(int(total), &pg)

	return utils.NewResponse(c).ResponseFormatter(fiber.StatusOK, "List All Books", nil, fiber.Map{
		"data":         getAllBooks,
		"total_data":   total,
		"current_page": pg.Page,
		"total_page":   countTotalPage,
	})
}

func (handler BookHandler) GetOneBooksHandler(c *fiber.Ctx) error {
	getBookUUID := c.Params("uuid")

	if !utils.ValidateUuid(getBookUUID) {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Invalid UUID", nil, fiber.Map{"error": "Invalid UUID"})
	}

	getBookModel, err := controller.NewBookController(handler.DB).GetOneBooks(getBookUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusNotFound, "Book Not Found", nil, fiber.Map{"error": "Book Not Found", "uuid": getBookUUID})
	} else if err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Fetch Detail", nil, fiber.Map{"error": map[string]string{"get_one_book_error": err.Error()}})
	}

	bookDetailPresenter := presenter.NewBookDetailPresenter()
	getOneBooks := bookDetailPresenter.Build(getBookModel)

	return utils.NewResponse(c).ResponseFormatter(http.StatusOK, "Get Book Detail", nil, fiber.Map{"detail": getOneBooks})
}

func (handler BookHandler) UpdateBookHandler(c *fiber.Ctx) error {
	bookValidator := validator.NewBookValidator()
	getBookUUID := c.Params("uuid")

	if !utils.ValidateUuid(getBookUUID) {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Invalid UUID", nil, fiber.Map{"error": "Invalid UUID"})
	}

	if err := c.BodyParser(&bookValidator); err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Failed Body Parsing", err, fiber.Map{"error": map[string]string{"body_parser_error": err.Error()}})
	}

	errValidate := bookValidator.Validate()
	if errValidate != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Invalid Form", nil, fiber.Map{"error": errValidate})
	}

	_, err := controller.NewBookController(handler.DB).GetOneBooks(getBookUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusNotFound, "Book Not Found", nil, fiber.Map{"error": "Book Not Found", "uuid": getBookUUID})
	} else if err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Fetch Detail", nil, fiber.Map{"error": map[string]string{"get_one_book_error": err.Error()}})
	}

	tx := handler.DB.Begin()

	updatedUUID, err := controller.NewBookController(tx).UpdateBook(bookValidator, getBookUUID)
	if err != nil {
		tx.Rollback()
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Update Data", err, fiber.Map{"error": map[string]string{"update_book_error": err.Error()}})
	}

	tx.Commit()
	return utils.NewResponse(c).ResponseFormatter(fiber.StatusOK, "Update Book Successfully", nil, fiber.Map{"updatedUUID": updatedUUID})
}

func (handler BookHandler) DeleteBookHandler(c *fiber.Ctx) error {
	getBookUUID := c.Params("uuid")

	if !utils.ValidateUuid(getBookUUID) {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusBadRequest, "Invalid UUID", nil, fiber.Map{"error": "Invalid UUID"})
	}

	_, err := controller.NewBookController(handler.DB).GetOneBooks(getBookUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusNotFound, "Book Not Found", nil, fiber.Map{"error": "Book Not Found", "uuid": getBookUUID})
	} else if err != nil {
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Fetch Detail", nil, fiber.Map{"error": map[string]string{"get_one_book_error": err.Error()}})
	}

	tx := handler.DB.Begin()

	deletedUUID, err := controller.NewBookController(tx).DeleteBook(getBookUUID)
	if err != nil {
		tx.Rollback()
		return utils.NewResponse(c).ResponseFormatter(fiber.StatusInternalServerError, "Failed Fetch Detail", nil, fiber.Map{"error": map[string]string{"get_one_book_error": err.Error()}})
	}

	tx.Commit()
	return utils.NewResponse(c).ResponseFormatter(fiber.StatusOK, "Deleted Book Successfully", nil, fiber.Map{"deletedUUID": deletedUUID})
}
