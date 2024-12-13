// Package models contains the models for the URL shortener application.
package models

import "time"

type Url struct {
	ID        int       `db:"id"`
	Url       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
}

// ShortUrl is a model for the short URL
type ShortUrl struct {
	ID        int       `db:"id"`
	UrlID     int       `db:"url_id"`
	ShortUrl  string    `db:"short_url"`
	CreatedAt time.Time `db:"created_at"`
}
