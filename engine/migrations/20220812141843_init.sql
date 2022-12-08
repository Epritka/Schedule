-- +goose Up
-- +goose StatementBegin
CREATE TABLE app_user (
    id INT GENERATED ALWAYS AS IDENTITY,
    telegram_user_id INT NOT NULL email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_superuser BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE app_user;

-- +goose StatementEnd