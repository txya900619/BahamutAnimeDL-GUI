-- +migrate Up

CREATE TABLE downloadedAnimations(
    sn INTEGER PRIMARY KEY NOT NULL,
    title TEXT NOT NULL ,
    episode TEXT NOT NULL
);

-- +migrate Down

DROP TABLE downloadedAnimations;