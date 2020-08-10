
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE reviews ADD user_id integer;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE reviews DROP user_id;
