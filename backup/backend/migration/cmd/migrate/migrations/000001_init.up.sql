CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
	logo TEXT NOT NULL,
	price INT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS oauth_states (
    state TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    provider_id BIGINT NOT NULL UNIQUE,
    login TEXT NOT NULL,
    display_name TEXT NOT NULL,
    email TEXT NOT NULL,
    balance INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),
    user_id BIGINT NOT NULL REFERENCES users(id),
    price INT NOT NULL CHECK (price >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
