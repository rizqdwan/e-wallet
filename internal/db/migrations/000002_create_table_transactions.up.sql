CREATE TABLE IF NOT EXISTS transactions(
    id               UUID PRIMARY KEY  DEFAULT gen_random_uuid(),
    from_wallet_id   UUID NULL,
    to_wallet_id     UUID NULl,
    transaction_type TEXT NOT NULL CHECK (transaction_type IN ('deposit', 'withdrawal', 'transfer')),
    status           TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed')),
    amount           BIGINT NOT NULL CHECK ( amount > 0 ),
    reference_id     VARCHAR(100) UNIQUE,
    idempotency_key  VARCHAR(100) UNIQUE,
    description      TEXT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT NOW()
)