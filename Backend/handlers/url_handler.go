package handlers

import (
	"log"
	"url-shortener/models"
	"url-shortener/services"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type UrlHandler struct {
	urlService services.UrlServiceInterface
}

func NewUrlHandler(urlService services.UrlServiceInterface) *UrlHandler {
	return &UrlHandler{urlService: urlService}
}

// Get the base URL from the environment variable using viper
var baseURL = viper.GetString("BASE_URL")

// const baseURL = "http://localhost:8080/api/urls/"

func (h *UrlHandler) CreateUrl(c *fiber.Ctx) error {
	// Parse request body
	var request struct {
		URL string `json:"url"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate URL
	if request.URL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL is required",
		})
	}

	// Create URL model
	url := &models.Url{
		Url: request.URL,
	}

	// Call service to create short URL
	shortUrl, err := h.urlService.CreateUrl(url)
	if err != nil {
		log.Println("Error creating short URL:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create short URL",
		})
	}

	// Return response with full short URL
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"original_url": url.Url,
		"short_url":    baseURL + shortUrl.ShortUrl,
		"created_at":   url.CreatedAt,
	})
}

func (h *UrlHandler) GetUrlByShortUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	url, err := h.urlService.GetUrlByShortUrl(shortUrl)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short URL not found",
		})
	}

	// redirect to the original URL
	return c.Redirect(url)
}
