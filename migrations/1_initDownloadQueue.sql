-- +migrate Up
CREATE TABLE downloadQueue(
    sn INTEGER PRIMARY KEY NOT NULL ,
    name TEXT NOT NULL ,
    ep TEXT NOT NULL ,
    sequence INTEGER NOT NULL ,
    Downloading INTEGER NOT NULL DEFAULT 0
);
-- +migrate Down
DROP TABLE downloadQueue;