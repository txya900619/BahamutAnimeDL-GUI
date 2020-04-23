-- +migrate Up

CREATE TABLE downloadedAnimations(
                                     id INTEGER PRIMARY KEY AUTOINCREMENT ,
                                     title TEXT NOT NULL ,
                                     part INTEGER NOT NULL
);

-- +migrate Down

DROP TABLE downloadedAnimations;