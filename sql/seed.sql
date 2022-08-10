CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(200) NOT NULL
);

INSERT INTO users (id, name, email) VALUES (1, 'Elle', 'elle@elle.com') ON CONFLICT DO NOTHING
