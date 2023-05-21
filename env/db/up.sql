/*
 *** Database ***

As PostgreSQL doesn't support IF NOT EXISTS to create a database; we have to use this approach to avoid creating it if
already exists.
 */
SELECT 'CREATE DATABASE transactions_api'
    WHERE NOT EXISTS
        (SELECT FROM pg_database WHERE datname = 'transactions_api')\gexec

/* Using transactions_api database from now on */
\c transactions_api

/* Set Brazil's TZ from Sao Paulo */
SET timezone = 'America/Sao_Paulo';

/* Table accounts */
CREATE TABLE IF NOT EXISTS accounts (
    id         SERIAL      PRIMARY KEY,
    document   VARCHAR(14) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
COMMENT ON COLUMN accounts.document IS 'Document length may vary between 11 (CPF) and 14 (CNPJ) chars.';

/* Table operations */
CREATE TABLE IF NOT EXISTS operations (
    id          SERIAL      PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

/* Table transactions */
CREATE TABLE IF NOT EXISTS transactions (
    id           SERIAL         PRIMARY KEY,
    account_id   INTEGER        NOT NULL,
    operation_id INTEGER        NOT NULL,
    amount       NUMERIC(15, 1) NOT NULL,
    created_at   TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_accounts
        FOREIGN KEY (account_id) REFERENCES accounts(id),
    CONSTRAINT fk_operations
        FOREIGN KEY (operation_id) REFERENCES operations(id)
);
