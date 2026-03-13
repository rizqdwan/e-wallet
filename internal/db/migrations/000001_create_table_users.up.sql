CREATE TABLE IF NOT EXISTS users
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid() ,
    username    VARCHAR(100) NOT NULL,
    email       VARCHAR(100) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP    NOT NULL DEFAULT NOW(),

    CONSTRAINT users_username_unique  UNIQUE (username),
    CONSTRAINT users_email_unique    UNIQUE (email)
);
