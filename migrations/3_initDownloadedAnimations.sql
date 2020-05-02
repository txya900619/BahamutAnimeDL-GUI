-- +migrate Up

CREATE TABLE downloadedAnimations(
                                     id INTEGER PRIMARY KEY AUTOINCREMENT ,
                                     title TEXT NOT NULL ,
                                     part TEXT NOT NULL
);

-- +migrate Down

DROP TABLE downloadedAnimations;