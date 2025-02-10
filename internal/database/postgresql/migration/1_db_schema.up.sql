BEGIN;


UPDATE schema_migrations SET dirty=false;
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS original_links
(
    id                       BIGSERIAL PRIMARY KEY,
    original_link                     VARCHAR(255) UNIQUE,
    shorter_link_id          BIGINT UNIQUE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);

CREATE TABLE IF NOT EXISTS shorter_links
(
    id                       BIGSERIAL PRIMARY KEY,
    shorter_link             VARCHAR(255) UNIQUE,
    created_at               TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() at TIME zone 'utc')
);