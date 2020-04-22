-- +migrations Up
CREATE TABLE downloadedAnimations(
    id INTEGER PRIMARY KEY AUTOINCREMENT ,
    title TEXT NOT NULL ,
    part INTEGER NOT NULL
);
-- +migrations Down
DROP TABLE downloadedAnimations;