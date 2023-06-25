CREATE TABLE IF NOT EXISTS user_transaction (
                id SERIAL PRIMARY KEY,
                balance DOUBLE PRECISION
);

CREATE INDEX idx_user_transaction_id ON user_transaction (id);