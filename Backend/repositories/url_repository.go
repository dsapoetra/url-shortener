package repositories

import (
	"url-shortener/models"

	"github.com/jmoiron/sqlx"
)

type UrlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

type UrlRepositoryInterface interface {
	InsertWithShortUrl(url *models.Url, shortUrl *models.ShortUrl) (*models.ShortUrl, error)
	GetUrlByShortUrl(shortUrl string) (string, error)
	IncrementCounterVisit(shortUrl string) error
}

// Insert a new URL into the database and short URL into the database
func (r *UrlRepository) InsertWithShortUrl(url *models.Url, shortUrl *models.ShortUrl) (*models.ShortUrl, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	// Check if URL exists, if not insert it
	urlQuery := `
        INSERT INTO urls (url, created_at)
        VALUES ($1, NOW())
        ON CONFLICT ON CONSTRAINT urls_url_key  -- Specify the unique constraint
        DO UPDATE SET url=EXCLUDED.url
        RETURNING id, url, created_at
    `

	err = tx.QueryRowx(urlQuery, url.Url).StructScan(url)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Always insert a new short_url
	shortUrlQuery := `
        INSERT INTO short_urls (url_id, short_url, created_at)
        VALUES ($1, $2, NOW())
        RETURNING id, short_url, created_at
    `

	err = tx.QueryRowx(shortUrlQuery, url.ID, shortUrl.ShortUrl).StructScan(shortUrl)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}

func (r *UrlRepository) GetUrlByShortUrl(shortUrl string) (string, error) {
	var originalUrl string
	query := `
        SELECT u.url 
        FROM urls u 
        JOIN short_urls su ON u.id = su.url_id 
        WHERE su.short_url = $1
    `

	err := r.db.Get(&originalUrl, query, shortUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (r *UrlRepository) IncrementCounterVisit(shortUrl string) error {
	query := `
		UPDATE short_urls
		SET counter_visit = counter_visit + 1
		WHERE short_url = $1
	`
	_, err := r.db.Exec(query, shortUrl)
	if err != nil {
		return err
	}

	return nil
}
