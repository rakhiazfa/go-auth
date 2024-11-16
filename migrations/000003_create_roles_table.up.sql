CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_roles_deleted_at ON roles (deleted_at);

ALTER TABLE roles
    ADD COLUMN search_text TEXT GENERATED ALWAYS AS (
        LOWER(COALESCE(name, ''))
    ) STORED;

CREATE INDEX IF NOT EXISTS idx_roles_search_text ON roles (search_text);