CREATE TABLE IF NOT EXISTS train (
    id BIGSERIAL PRIMARY KEY,
    details jsonb,
    price numeric
);