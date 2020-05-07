-- +migrate Up

CREATE TABLE downloadedAnimations(
    sn INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL ,
    episode TEXT NOT NULL,
    spacial INTEGER NOT NULL  DEFAULT 0
);

-- +migrate Down

DROP TABLE downloadedAnimations;