package main

import (
	bookHandler "github.com/PickHD/bookstore-api/src/apps/books/handler"
	"github.com/PickHD/bookstore-api/src/utils"

	"github.com/PickHD/bookstore-api/src/utils/db"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	databaseConnection := db.GetDB()

	v1 := app.Group("/api/v1")
	{
		handler := bookHandler.NewBookHandler(databaseConnection)
		v1.Post("/book", handler.CreateBookHandler)
		v1.Get("/book", handler.GetAllBooksHandler)
		v1.Get("/book/:uuid", handler.GetOneBooksHandler)
		v1.Put("/book/:uuid", handler.UpdateBookHandler)
		v1.Delete("/book/:uuid", handler.DeleteBookHandler)
	}

	app.Use(func(c *fiber.Ctx) error {
		return utils.NewResponse(c).ResponseFormatter(404, "Route Not Found", nil, nil)
	})
}
