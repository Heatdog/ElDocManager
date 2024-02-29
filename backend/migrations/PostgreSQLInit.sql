DROP TABLE Users;

CREATE EXTENSION citext;
CREATE DOMAIN EMAIL AS citext
  CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );

CREATE type users_role AS enum ('Administrator', 'HeadOfDepatment', 'Worker')

CREATE TABLE IF NOT EXISTS Users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email EMAIL NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role users_role NOT NULL DEFAULT 'Worker',
    created NOT NULL TIMESTAMP DEFAULT Now(),
    is_confirmed NOT NULL BOOLEAN DEFAULT FALSE,
    CHECK (LENGTH(login) > 3 AND LENGTH(password) > 5)
);