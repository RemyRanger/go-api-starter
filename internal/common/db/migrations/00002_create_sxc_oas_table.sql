-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS sxc_oas (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    api_version VARCHAR(256) NOT NULL,
    openapi_version VARCHAR(1024),
    json_spec JSONB NOT NULL,  -- Champ JSONB
    api_id uuid,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_apis_oas FOREIGN KEY (api_id) REFERENCES sxc_api (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS sxc_oas;