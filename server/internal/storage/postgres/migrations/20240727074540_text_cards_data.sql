-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS text_data(
    user_login VARCHAR(50) NOT NULL,
    name VARCHAR(1000) NOT NULL,
    text VARCHAR(1000) NOT NULL,
    meta VARCHAR(1000),
    PRIMARY KEY(user_login, name)
);
CREATE TABLE IF NOT EXISTS cards_data(
    user_login VARCHAR(50) NOT NULL,
    name VARCHAR(1000) NOT NULL,
    number VARCHAR(12) NOT NULL,
    expiration_month INTEGER NOT NULL,
    expiration_year INTEGER NOT NULL,
    code INTEGER NOT NULL,
    cardholder VARCHAR(256),
    meta VARCHAR(1000),
    PRIMARY KEY(user_login, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS text_data;
DROP TABLE IF EXISTS cards_data;
-- +goose StatementEnd
