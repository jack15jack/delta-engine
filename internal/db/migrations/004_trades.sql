CREATE TABLE IF NOT EXISTS positions (
    id SERIAL PRIMARY KEY,

    portfolio_id INT NOT NULL,
    ticker TEXT NOT NULL,

    quantity DOUBLE PRECISION NOT NULL,
    avg_cost DOUBLE PRECISION NOT NULL,

    updated_at TIMESTAMP DEFAULT NOW()
);