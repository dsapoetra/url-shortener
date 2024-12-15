package services

import (
	"Backend/models"
	"Backend/repositories"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

type UrlService struct {
	urlRepo repositories.UrlRepositoryInterface // Change from *repositories.UrlRepository
}

func NewUrlService(urlRepo repositories.UrlRepositoryInterface) *UrlService {
	return &UrlService{urlRepo: urlRepo}
}

type UrlServiceInterface interface {
	CreateUrl(url *models.Url) (*models.ShortUrl, error)
	GetUrlByShortUrl(shortUrl string) (string, error)
}

func (s *UrlService) CreateUrl(url *models.Url) (*models.ShortUrl, error) {
	maxRetries := 5
	var shortUrl *models.ShortUrl
	var err error

	for i := 0; i < maxRetries; i++ {
		// Generate a new short URL each attempt
		shortUrl = &models.ShortUrl{
			ShortUrl: generateShortUrl(url.Url), // Your existing short URL generation function
		}

		shortUrl, err = s.urlRepo.InsertWithShortUrl(url, shortUrl)
		if err == nil {
			return shortUrl, nil
		}

		// If it's not a duplicate error, return the error
		if !strings.Contains(err.Error(), "duplicate key value") {
			return nil, err
		}

		// If it is a duplicate, continue the loop to try again
	}

	return nil, fmt.Errorf("failed to generate unique short URL after %d attempts", maxRetries)
}

func generateShortUrl(url string) string {
	if url == "" {
		return ""
	}
	// Define the character set for base62 encoding
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	// Create unique input by combining URL and current timestamp
	timestamp := time.Now().UnixNano()
	input := fmt.Sprintf("%s%d", url, timestamp)

	// Create hash
	hash := sha256.Sum256([]byte(input))
	seed := binary.BigEndian.Uint64(hash[:8])
	r := rand.New(rand.NewSource(uint64(seed)))

	// Generate random characters
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)
}

func (s *UrlService) GetUrlByShortUrl(shortUrl string) (string, error) {
	err := s.urlRepo.IncrementCounterVisit(shortUrl)
	if err != nil {
		return "", err
	}

	return s.urlRepo.GetUrlByShortUrl(shortUrl)
}
