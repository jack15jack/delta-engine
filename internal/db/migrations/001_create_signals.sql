CREATE TABLE IF NOT EXISTS signals (
    id SERIAL PRIMARY KEY,

    ticker TEXT NOT NULL,
    strategy TEXT NOT NULL,

    signal_type TEXT NOT NULL, -- BUY / SELL / HOLD
    price DOUBLE PRECISION NOT NULL,
    strength DOUBLE PRECISION,

    timestamp TIMESTAMP NOT NULL,

    created_at TIMESTAMP DEFAULT NOW()
);