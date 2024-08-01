-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS login_password_data(
    user_login VARCHAR(50) NOT NULL,
    name VARCHAR(1000) NOT NULL,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    meta VARCHAR(1000),
    PRIMARY KEY(user_login, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS login_password_data;
-- +goose StatementEnd
