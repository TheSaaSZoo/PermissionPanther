
-- +migrate Up
CREATE TABLE IF NOT EXISTS keys (
  id TEXT NOT NULL,
  secret_hash TEXT NOT NULL,
  ns TEXT NOT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS keys;
