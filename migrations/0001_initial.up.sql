-- db/migrations/0001_initial.up.sql

-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    inn VARCHAR(12) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    is_admin BOOLEAN NOT NULL DEFAULT false
);

-- Создание таблицы фискальных модулей
CREATE TABLE IF NOT EXISTS fiscal_modules (
    id SERIAL PRIMARY KEY,
    factory_number VARCHAR(50) NOT NULL,
    fiscal_number VARCHAR(50) NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id)
);

-- Создание таблицы торговых точек (терминалов)
CREATE TABLE IF NOT EXISTS terminals (
    id SERIAL PRIMARY KEY,
    inn VARCHAR(12) NOT NULL,
    company_name VARCHAR(100) NOT NULL,
    address VARCHAR(200),
    cash_register_number VARCHAR(50),
    module_number VARCHAR(50),
    assembly_number VARCHAR(50),
    last_request_date TIMESTAMP,
    update_date TIMESTAMP,
    status VARCHAR(20),
    user_id INTEGER NOT NULL REFERENCES users(id),
    database_update_date TIMESTAMP
);
