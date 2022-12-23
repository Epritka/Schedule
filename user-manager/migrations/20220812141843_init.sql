-- +goose Up
-- +goose StatementBegin
CREATE TABLE app_user (
    id INT GENERATED ALWAYS AS IDENTITY,
    telegram_user_id INT NOT NULL,
    PRIMARY KEY(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE app_user;

-- +goose StatementEnd