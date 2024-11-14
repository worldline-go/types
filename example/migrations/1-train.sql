CREATE TABLE IF NOT EXISTS train (
    id BIGSERIAL PRIMARY KEY,
    details jsonb,
    additionals jsonb,
    price numeric,
    last_price numeric,
    rate numeric,
    custom_number integer
);