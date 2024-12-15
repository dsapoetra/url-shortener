package routes

import (
	"Backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, urlHandler *handlers.UrlHandler) {
	app.Get("/:shortUrl", urlHandler.GetUrlByShortUrl)

	api := app.Group("/api")

	api.Post("/urls", urlHandler.CreateUrl)
}
