CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    photo VARCHAR(255),
    client_id UUID REFERENCES users(id),
    worker_id UUID REFERENCES workers(id),
    service_id UUID REFERENCES services(id),
    status VARCHAR(20) CHECK (status IN ('pending', 'in_progress', 'completed', 'canceled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);