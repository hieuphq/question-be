-- +migrate Up
CREATE TABLE topics
(
    id SERIAL PRIMARY KEY,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz DEFAULT NULL,
    name varchar(200),
    code varchar(20),
    status varchar(20)
);

-- +migrate Down
DROP TABLE IF EXISTS topics;

