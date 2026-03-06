-- Criar bases de dados
CREATE DATABASE orders_db;
CREATE DATABASE payments_db;
CREATE DATABASE inventory_db;

-- Tabela orders
\c orders_db;
CREATE TABLE IF NOT EXISTS orders (
    id           VARCHAR(36) PRIMARY KEY,
    customer_id  VARCHAR(36) NOT NULL,
    items        JSONB NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status       VARCHAR(20) NOT NULL,
    created_at   TIMESTAMP NOT NULL
);

-- Tabela payments
\c payments_db;
CREATE TABLE IF NOT EXISTS payments (
    id         VARCHAR(36) PRIMARY KEY,
    order_id   VARCHAR(36) NOT NULL,
    amount     DECIMAL(10,2) NOT NULL,
    status     VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Tabela inventory
\c inventory_db;
CREATE TABLE IF NOT EXISTS inventory (
    id         VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) NOT NULL,
    quantity   INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);