-- Alter short_url column to have counter_visit column
ALTER TABLE short_urls ADD COLUMN counter_visit INT NOT NULL DEFAULT 0;