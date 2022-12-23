-- +goose Up
-- +goose StatementBegin
ALTER TABLE app_user ADD CONSTRAINT email_unique UNIQUE (email);
ALTER TABLE app_user ADD CONSTRAINT tg_user_id_unique UNIQUE (telegram_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE app_user DROP CONSTRAINT email_unique;
ALTER TABLE app_user DROP CONSTRAINT tg_user_id_unique;
-- +goose StatementEnd
