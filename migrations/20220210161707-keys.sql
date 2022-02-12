
-- +migrate Up
CREATE TABLE IF NOT EXISTS keys (
  id TEXT NOT NULL,
  secret_hash TEXT NOT NULL,
  ns TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS keys_by_ns ON keys(ns);

-- +migrate Down
DROP INDEX IF EXISTS keys_by_ns;
DROP TABLE IF EXISTS keys;
