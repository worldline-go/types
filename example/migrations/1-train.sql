CREATE TABLE IF NOT EXISTS train (
    id BIGSERIAL PRIMARY KEY,
    details jsonb,
    additionals jsonb,
    price numeric
);