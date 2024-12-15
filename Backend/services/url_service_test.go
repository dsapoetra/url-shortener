package services

import (
	"fmt"
	"testing"
	"time"

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
