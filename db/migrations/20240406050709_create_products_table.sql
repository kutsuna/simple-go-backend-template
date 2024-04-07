-- +goose Up
CREATE TABLE products (
    id TEXT,
    name TEXT NOT NULL,
    price REAL NOT NULL,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE products;
