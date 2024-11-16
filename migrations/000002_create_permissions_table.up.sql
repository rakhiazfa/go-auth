CREATE TABLE IF NOT EXISTS permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    service_key VARCHAR(255) NULL,
    name VARCHAR(255) NOT NULL,
    method VARCHAR(50) NOT NULL,
    path VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE(service_key, method, path)
);

CREATE INDEX IF NOT EXISTS idx_permissions_method ON permissions (method);

CREATE INDEX IF NOT EXISTS idx_permissions_path ON permissions (path);

CREATE INDEX IF NOT EXISTS idx_permissions_deleted_at ON permissions (deleted_at);

ALTER TABLE permissions
    ADD COLUMN search_text TEXT GENERATED ALWAYS AS (
        LOWER(COALESCE(name, '') || ' ' || COALESCE(method, '') || ' ' || COALESCE(path, ''))
    ) STORED;

CREATE INDEX IF NOT EXISTS idx_permissions_search_text ON permissions (search_text);