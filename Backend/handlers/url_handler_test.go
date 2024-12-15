// Backend/handlers/url_handler_test.go
package handlers

import (
	"Backend/models"
	mock_services "Backend/services/mocks"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"net/http/httptest"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUrl_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)

	// Setup expected service response
	expectedShortUrl := &models.ShortUrl{
		ShortUrl: "abc123",
	}
	mockService.EXPECT().CreateUrl(gomock.Any()).Return(expectedShortUrl, nil)

	// Create test app and request
	app := fiber.New()
	app.Post("/shorten", handler.CreateUrl)

	// Fix: Use valid JSON with a URL
	req := httptest.NewRequest("POST", "/shorten", strings.NewReader(`{"url": "https://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Contains(t, result["short_url"], "abc123")
}

func TestCreateUrl_EmptyURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)

	app := fiber.New()
	app.Post("/shorten", handler.CreateUrl)

	req := httptest.NewRequest("POST", "/shorten", strings.NewReader(`{"url": ""}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestCreateUrl_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)

	app := fiber.New()
	app.Post("/shorten", handler.CreateUrl)

	req := httptest.NewRequest("POST", "/shorten", strings.NewReader(`{invalid json}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestCreateUrl_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)

	mockService.EXPECT().CreateUrl(gomock.Any()).Return(nil, errors.New("service error"))

	app := fiber.New()
	app.Post("/shorten", handler.CreateUrl)

	req := httptest.NewRequest("POST", "/shorten", strings.NewReader(`{"url": "https://example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestGetUrlByShortUrl_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)

	mockService.EXPECT().GetUrlByShortUrl("abc123").Return("https://example.com", nil)

	app := fiber.New()
	app.Get("/:shortUrl", handler.GetUrlByShortUrl)

	req := httptest.NewRequest("GET", "/abc123", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusFound, resp.StatusCode) // 302 for redirect
}

func TestGetUrlByShortUrl_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockUrlServiceInterface(ctrl)
	handler := NewUrlHandler(mockService)

	mockService.EXPECT().GetUrlByShortUrl("notfound").Return("", errors.New("not found"))

	app := fiber.New()
	app.Get("/:shortUrl", handler.GetUrlByShortUrl)

	req := httptest.NewRequest("GET", "/notfound", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}
