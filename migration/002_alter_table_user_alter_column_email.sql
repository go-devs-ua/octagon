-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE "user"
    ALTER COLUMN email TYPE VARCHAR(320);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE "user"
    ALTER COLUMN email TYPE VARCHAR(255);
