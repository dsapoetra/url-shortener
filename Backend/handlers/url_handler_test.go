package handlers

import (
	"Backend/models"
	mock_services "Backend/services/mocks"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUrlHandler_CreateUrl(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)
	app := fiber.New()
	app.Post("/api/urls", handler.CreateUrl)

	// Test case
	t.Run("successful URL creation", func(t *testing.T) {
		// Prepare request
		req := httptest.NewRequest("POST", "/api/urls", strings.NewReader(`{"url":"https://www.google.com"}`))
		req.Header.Set("Content-Type", "application/json")

		// Set expectations
		expectedShortUrl := &models.ShortUrl{ShortUrl: "abc123"}
		mockService.EXPECT().
			CreateUrl(gomock.Any()).
			Return(expectedShortUrl, nil)

		// Execute request
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		// Parse response
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		assert.Contains(t, result["short_url"], "abc123")
	})
}

func TestUrlHandler_GetUrlByShortUrl(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)
	app := fiber.New()
	app.Get("/:shortUrl", handler.GetUrlByShortUrl)

	// Test case
	t.Run("successful URL retrieval", func(t *testing.T) {
		// Set expectations
		mockService.EXPECT().
			GetUrlByShortUrl("abc123").
			Return("https://www.google.com", nil)

		// Execute request
		req := httptest.NewRequest("GET", "/abc123", nil)
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusFound, resp.StatusCode) // 302 for redirect
		assert.Equal(t, "https://www.google.com", resp.Header.Get("Location"))
	})
}
