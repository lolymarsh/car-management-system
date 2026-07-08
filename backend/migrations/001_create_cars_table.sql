-- +goose Up
CREATE TABLE cars (
    car_id              UUID NOT NULL DEFAULT gen_random_uuid(),
    PRIMARY KEY (car_id),
    registration_number VARCHAR(50)  NOT NULL,
    brand               VARCHAR(100) NOT NULL,
    model               VARCHAR(100) NOT NULL,
    color               VARCHAR(50)  NOT NULL DEFAULT '',
    year                INTEGER      NOT NULL DEFAULT 0,
    notes               TEXT         NOT NULL DEFAULT '',
    created_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

CREATE INDEX idx_cars_deleted_at ON cars(deleted_at);
CREATE INDEX idx_cars_registration ON cars(registration_number) WHERE deleted_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS idx_cars_registration;
DROP INDEX IF EXISTS idx_cars_deleted_at;
DROP TABLE IF EXISTS cars;
