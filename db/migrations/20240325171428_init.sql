-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quote (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  quote VARCHAR(255) UNIQUE NOT NULL,
  character VARCHAR(255) NOT NULL,
  season SMALLINT NOT NULL,
  episode SMALLINT NOT NULL,
  title VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE quote;
-- +goose StatementEnd
