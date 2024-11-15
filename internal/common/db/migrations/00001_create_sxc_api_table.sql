-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS sxc_api (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    description VARCHAR(1024),
    version VARCHAR(32) NOT NULL,
    base_url VARCHAR(1024) NOT NULL,
    gateway_mode VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    active_oas_id uuid,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS sxc_api;