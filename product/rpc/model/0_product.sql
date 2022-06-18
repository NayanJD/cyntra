CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS product (
    id text PRIMARY KEY,
    name text NOT NULL UNIQUE,
    price integer NOT NULL,
    description text NOT NULL,
    size text NOT NULL,
    color text NOT NULL,
    brand text NOT NULL,
    shade text NOT NULL,
    wash_care text,
    stretchable text,
    distress text,
    features text,
    fade text,
    fabric text NOT NULL,
    category text NOT NULL,
    countryOfOrigin text NOT NULL,
    discount integer NOT NULL DEFAULT 0,
    quantity integer NOT NULL,
    created_at  timestamptz NOT NULL,
    updated_at  timestamptz NOT NULL,
    archived_at  timestamptz
)