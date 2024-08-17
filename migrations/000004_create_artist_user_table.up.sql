CREATE TABLE IF NOT EXISTS artist_user (
    artist_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    PRIMARY KEY (artist_id, user_id),
    FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO artist_user(artist_id, user_id) VALUES
(1, 1), (1, 2), (1, 3),
(2, 1), (2, 2), (2, 3), (2, 4), (2, 5);