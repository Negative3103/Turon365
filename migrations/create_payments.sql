CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY,
    job_id UUID REFERENCES jobs(id),
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('pending', 'paid', 'failed')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);