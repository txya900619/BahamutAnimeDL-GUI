-- +migrate Up
CREATE TABLE downloadQueue(
    sn INTEGER NOT NULL ,
    name TEXT NOT NULL ,
    ep INTEGER NOT NULL ,
    sequence INTEGER PRIMARY KEY ,
    Downloading INTEGER NOT NULL DEFAULT 0
);
-- +migrate Down
DROP TABLE downloadQueue;