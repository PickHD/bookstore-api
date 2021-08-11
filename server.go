package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"time"

	bookModel "github.com/PickHD/bookstore-api/src/apps/books/model"
	database "github.com/PickHD/bookstore-api/src/utils/db"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&bookModel.Books{})
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(logger.New())
	app.Use(requestid.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
		return
	}
	time.Local = loc

	db := database.Init()

	Migrate(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	SetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
