
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE reviews (
    id integer UNIQUE NOT NULL,
    public_id integer UNIQUE NOT NULL,
    movie_id integer,
    comment varchar(2000),
    score float NOT NULL,
    spoiler boolean,
    appreciation_day datetime,
    appreciation_method varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE reviews;
