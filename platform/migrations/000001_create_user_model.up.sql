BEGIN;

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";


CREATE TABLE users(
    id serial primary key,
    login VARCHAR(255) NOT null,
    password_hash VARCHAR(255) NOT null,
    email VARCHAR(255) NOT null
);

COMMIT;