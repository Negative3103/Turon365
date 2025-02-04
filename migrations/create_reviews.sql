CREATE TABLE IF NOT EXISTS reviews (
    id UUID PRIMARY KEY,
    job_id UUID REFERENCES jobs(id),
    rating INT CHECK (rating BETWEEN 1 AND 5) NOT NULL,
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);