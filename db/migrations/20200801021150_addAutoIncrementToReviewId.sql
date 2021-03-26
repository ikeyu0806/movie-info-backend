
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE reviews CHANGE id id integer AUTO_INCREMENT;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE reviews CHANGE id id integer NOT NULL;
