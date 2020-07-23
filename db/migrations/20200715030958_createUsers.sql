
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id integer AUTO_INCREMENT,
    name varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
