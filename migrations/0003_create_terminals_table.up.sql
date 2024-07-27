CREATE TABLE IF NOT EXISTS terminals (
    id SERIAL PRIMARY KEY,
    inn VARCHAR(12) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    cash_register_number VARCHAR(255) UNIQUE NOT NULL,
    module_number VARCHAR(255) UNIQUE NOT NULL,
    assembly_number VARCHAR(255) UNIQUE NOT NULL,
    last_request_date TIMESTAMP WITHOUT TIME ZONE,
    database_update_date TIMESTAMP WITHOUT TIME ZONE,
    free_record_balance INTEGER NOT NULL,
    status BOOLEAN DEFAULT true,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);
