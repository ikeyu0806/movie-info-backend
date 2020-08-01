
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE reviews ADD created_at datetime,
  ADD updated_at datetime,
  ADD deleted_at datetime;
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE reviews DROP created_at,
  DROP updated_at,
  DROP deleted_at;
