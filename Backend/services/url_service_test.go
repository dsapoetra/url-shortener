package services

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"Backend/models"
	mock_repositories "Backend/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGenerateShortUrl(t *testing.T) {
	shortUrl := generateShortUrl("https://www.google.com")
	fmt.Println(shortUrl)
	assert.Equal(t, len(shortUrl), 6)
}

func TestGenerateShortUrlWithSameUrl(t *testing.T) {
	shortUrl := generateShortUrl("https://www.google.com")
	time.Sleep(1 * time.Second)
	shortUrl2 := generateShortUrl("https://www.google.com")
	assert.NotEqual(t, shortUrl, shortUrl2, "Short URLs should be different even for the same input URL")
}

// Test get url by short url
func TestGetUrlByShortUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)

	// Set expectations
	mockRepo.EXPECT().GetUrlByShortUrl(gomock.Any()).Return("https://www.google.com", nil)
	mockRepo.EXPECT().IncrementCounterVisit(gomock.Any()).Return(nil)

	// Create service with mock
	urlService := NewUrlService(mockRepo)

	// Test
	url, err := urlService.GetUrlByShortUrl("abc123")
	assert.NoError(t, err)
	assert.Equal(t, "https://www.google.com", url)
}

func TestGenerateShortUrl_InvalidUrl(t *testing.T) {
	shortUrl := generateShortUrl("")
	assert.Equal(t, "", shortUrl, "Empty URL input should return an empty short URL")
}

func TestGetUrlByShortUrl_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)

	// Simulate a repository error
	mockRepo.EXPECT().GetUrlByShortUrl(gomock.Any()).Return("", errors.New("database error"))
	mockRepo.EXPECT().IncrementCounterVisit(gomock.Any()).Return(nil)

	service := UrlService{urlRepo: mockRepo}
	url, err := service.GetUrlByShortUrl("invalid_short_url")

	assert.Error(t, err, "An error should be returned when the repository fails")
	assert.Empty(t, url, "URL should be empty on repository error")
}

func TestGenerateShortUrl_LongUrl(t *testing.T) {
	longUrl := "https://www.example.com/" + string(make([]byte, 5000)) // Simulate a very long URL
	shortUrl := generateShortUrl(longUrl)
	assert.Equal(t, len(shortUrl), 6, "Generated short URL should still have the correct length")
}

func TestGetUrlByShortUrl_EmptyResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)

	// Simulate empty result
	mockRepo.EXPECT().GetUrlByShortUrl(gomock.Any()).Return("", nil)
	mockRepo.EXPECT().IncrementCounterVisit(gomock.Any()).Return(nil)

	service := UrlService{urlRepo: mockRepo}
	url, err := service.GetUrlByShortUrl("nonexistent_short_url")

	assert.NoError(t, err, "No error should be returned for empty results")
	assert.Empty(t, url, "URL should be empty for nonexistent short URL")
}

func TestCreateUrl_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)
	url := &models.Url{Url: "https://example.com"}
	expectedShortUrl := &models.ShortUrl{ShortUrl: "abc123"}

	mockRepo.EXPECT().InsertWithShortUrl(url, gomock.Any()).Return(expectedShortUrl, nil)

	service := NewUrlService(mockRepo)
	result, err := service.CreateUrl(url)

	assert.NoError(t, err)
	assert.Equal(t, expectedShortUrl, result)
}

func TestCreateUrl_DuplicateRetry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)
	url := &models.Url{Url: "https://example.com"}
	duplicateErr := fmt.Errorf("duplicate key value")
	successShortUrl := &models.ShortUrl{ShortUrl: "abc123"}

	// First attempt fails with duplicate, second succeeds
	mockRepo.EXPECT().InsertWithShortUrl(url, gomock.Any()).Return(nil, duplicateErr)
	mockRepo.EXPECT().InsertWithShortUrl(url, gomock.Any()).Return(successShortUrl, nil)

	service := NewUrlService(mockRepo)
	result, err := service.CreateUrl(url)

	assert.NoError(t, err)
	assert.Equal(t, successShortUrl, result)
}

func TestCreateUrl_MaxRetriesExceeded(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)
	url := &models.Url{Url: "https://example.com"}
	duplicateErr := fmt.Errorf("duplicate key value")

	// Expect 5 attempts all failing with duplicate error
	for i := 0; i < 5; i++ {
		mockRepo.EXPECT().InsertWithShortUrl(url, gomock.Any()).Return(nil, duplicateErr)
	}

	service := NewUrlService(mockRepo)
	result, err := service.CreateUrl(url)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to generate unique short URL after 5 attempts")
	assert.Nil(t, result)
}

func TestCreateUrl_NonDuplicateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockUrlRepositoryInterface(ctrl)
	url := &models.Url{Url: "https://example.com"}
	dbErr := fmt.Errorf("database connection error")

	mockRepo.EXPECT().InsertWithShortUrl(url, gomock.Any()).Return(nil, dbErr)

	service := NewUrlService(mockRepo)
	result, err := service.CreateUrl(url)

	assert.Error(t, err)
	assert.Equal(t, dbErr, err)
	assert.Nil(t, result)
}
