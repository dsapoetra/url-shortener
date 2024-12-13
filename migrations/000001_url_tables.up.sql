-- Create the urls table and short_urls table
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_urls_url ON urls(url);

CREATE TABLE short_urls (
    id SERIAL PRIMARY KEY,
    url_id INT NOT NULL,
    short_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (url_id) REFERENCES urls(id)
);

-- Indexes
CREATE INDEX idx_short_urls_short_url ON short_urls(short_url);
