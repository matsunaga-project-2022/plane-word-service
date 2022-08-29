-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS words (
    id char(36) NOT NULL,
    user_id char(36) NOT NULL,
    word_name varchar(50) NOT NULL,
    meaning varchar(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS words;
-- +goose StatementEnd
