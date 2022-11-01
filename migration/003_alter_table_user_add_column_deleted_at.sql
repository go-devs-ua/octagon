-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE "user"
    ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE "user"
    DROP COLUMN deleted_at;
