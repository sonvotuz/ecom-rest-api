CREATE TYPE statusenum AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders(
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total REAL NOT NULL,
    status statusenum NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP    
);