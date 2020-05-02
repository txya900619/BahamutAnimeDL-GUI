-- +migrate Up
CREATE TABLE downloadState(
  id INTEGER PRIMARY KEY  AUTOINCREMENT ,
  sn INTEGER NOT NULL ,
  part TEXT NOT NULL ,
  success INTEGER NOT NULL DEFAULT 0
);

-- +migrate Down
DROP TABLE downloadState;
