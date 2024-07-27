CREATE TABLE IF NOT EXISTS fiscal_modules (
    id SERIAL PRIMARY KEY,
    factory_number VARCHAR(255) UNIQUE NOT NULL,
    fiscal_number VARCHAR(255) UNIQUE NOT NULL,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);
