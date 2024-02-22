CREATE TABLE IF NOT EXISTS Users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
);