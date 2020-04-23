-- +migrate Up
CREATE TABLE downloadStatus(
  id INTEGER PRIMARY KEY  AUTOINCREMENT ,
  sn INTEGER NOT NULL ,
  part INTEGER NOT NULL ,
  success INTEGER NOT NULL DEFAULT 0
);

-- +migrate Down
DROP TABLE downloadStatus;
