-- db/migrations/0001_initial.down.sql

-- Удаление таблицы торговых точек (терминалов)
DROP TABLE IF EXISTS terminals;

-- Удаление таблицы фискальных модулей
DROP TABLE IF EXISTS fiscal_modules;

-- Удаление таблицы пользователей
DROP TABLE IF EXISTS users;