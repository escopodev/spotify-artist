CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR NOT NULL
);

INSERT INTO users (username) VALUES 
('jose'), ('maria'), ('pedro'), ('lucas'), ('ana');