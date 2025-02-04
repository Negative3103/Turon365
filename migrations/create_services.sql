CREATE TABLE IF NOT EXISTS services (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    photo VARCHAR(255),
    worker_id UUID REFERENCES workers(id),
    category_id UUID REFERENCES categories(id),
    location_id UUID REFERENCES locations(id),
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);