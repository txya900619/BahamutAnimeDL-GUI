-- +migrate Up
CREATE TABLE downloadStatus(
  sn INTEGER NOT NULL ,
  part INTEGER NOT NULL ,
  success INTEGER NOT NULL DEFAULT 0
);

-- +migrate Down
DROP TABLE downloadStatus;
