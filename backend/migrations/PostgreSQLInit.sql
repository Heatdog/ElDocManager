CREATE EXTENSION citext;
CREATE DOMAIN EMAIL AS citext
  CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );

CREATE TABLE IF NOT EXISTS Users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email EMAIL NOT NULL,
    password VARCHAR(255) NOT NULL,
    role INTEGER NOT NULL,
    created TIMESTAMP DEFAULT Now(),
    is_created BOOLEAN DEFAULT FALSE
);