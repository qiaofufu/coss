package api

import (
	"coss/service/access/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func initAPI(app *fiber.App) {
	app.Post("/:bucketName/:objectName", handler.CreateObject())
	app.Get("/:bucketName/:objectName", handler.GetObject())
}

func MustStartAPI(port string) {
	var (
		app = fiber.New(fiber.Config{
			BodyLimit: 1024 * 1024 * 1024,
		})
	)
	initAPI(app)
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
