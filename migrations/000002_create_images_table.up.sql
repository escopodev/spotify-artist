CREATE TABLE IF NOT EXISTS images (
    id BIGSERIAL PRIMARY KEY,
    artist_id BIGINT NOT NULL,
    uid VARCHAR NOT NULL,
    width INT,
    height INT,
    FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE
);

INSERT INTO images (artist_id, uid, width, height) VALUES 
(1, 'ab6761610000e5eb7ae97ba8000c4ef0d1b84d82', 640, 640),
(1, 'ab676161000051747ae97ba8000c4ef0d1b84d82', 320, 320),
(1, 'ab6761610000f1787ae97ba8000c4ef0d1b84d82', 160, 160);

INSERT INTO images (artist_id, uid, width, height) VALUES 
(2, 'ab6761610000e5ebc884df599abc793c116cdf15', 640, 640),
(2, 'ab67616100005174c884df599abc793c116cdf15', 320, 320),
(2, 'ab6761610000f178c884df599abc793c116cdf15', 160, 160);