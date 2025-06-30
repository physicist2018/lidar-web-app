BEGIN;

CREATE TABLE experiments (
    id bigserial PRIMARY KEY,
    title varchar(255) NOT NULL,
    user_id int NOT NULL REFERENCES users(id)
);

COMMIT;