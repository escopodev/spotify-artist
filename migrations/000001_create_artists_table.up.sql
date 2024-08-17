CREATE TABLE IF NOT EXISTS artists (
    id BIGSERIAL PRIMARY KEY,
    uid VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    popularity INT NOT NULL,
    genres TEXT []
);

INSERT INTO artists (uid, name, popularity, genres) VALUES 
('7HnkRhoGqYLTasI52iJoE7', 'Chico Science', 44, '{"afrofuturismo brasileiro", "brazilian hip hop", "brazilian reggae", "coco", "hard rock brasileiro", "manguebeat", "mpb", "samba"}');

INSERT INTO artists (uid, name, popularity, genres) VALUES 
('7jy3rLJdDQY21OgRLCZ9sD', 'Foo Fighters', 75, '{"alternative metal", "alternative rock", "modern rock", "permanent wave", "post-grunge", "rock"}');


