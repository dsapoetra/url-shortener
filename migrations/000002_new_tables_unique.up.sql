-- Drop existing tables if they exist
DROP TABLE IF EXISTS short_urls;
DROP TABLE IF EXISTS urls;

-- Create the urls table with named unique constraint
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT urls_url_key UNIQUE (url)  -- Named unique constraint
);

-- Create the short_urls table with unique constraint
CREATE TABLE short_urls (
    id SERIAL PRIMARY KEY,
    url_id INT NOT NULL,
    short_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT short_urls_short_url_key UNIQUE (short_url),  -- Named unique constraint
    FOREIGN KEY (url_id) REFERENCES urls(id)
);

-- Indexes
CREATE INDEX idx_urls_url ON urls(url);
CREATE INDEX idx_short_urls_short_url ON short_urls(short_url);