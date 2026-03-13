CREATE TABLE IF NOT EXISTS wallets (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL,
    balance      BIGINT NOT NULL DEFAULT 0,
    wallet_number VARCHAR(20) NOT NULL,
    is_active    BOOLEAN NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES users(id),

    CONSTRAINT wallet_user_id_unique UNIQUE (user_id),
    CONSTRAINT wallets_number_unique UNIQUE (wallet_number),
    CONSTRAINT wallets_balance_positive CHECK (balance >= 0)
)