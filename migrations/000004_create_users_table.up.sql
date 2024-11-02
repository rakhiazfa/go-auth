CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    avatar VARCHAR(255) NULL,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

ALTER TABLE users
    ADD COLUMN search_text TEXT GENERATED ALWAYS AS (
        LOWER(COALESCE(name, '') || ' ' || COALESCE(username, '') || ' ' || COALESCE(email, ''))
    ) STORED;

CREATE INDEX IF NOT EXISTS idx_users_search_text ON users (search_text);