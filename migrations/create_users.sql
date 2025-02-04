CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    rating NUMERIC,
    role VARCHAR(20) CHECK (role IN ('admin', 'worker', 'client')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);