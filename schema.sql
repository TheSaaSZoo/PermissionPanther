CREATE TABLE IF NOT EXISTS relations (
  object TEXT NOT NULL,
  entity TEXT NOT NULL,
  permission TEXT NOT NULL,
  ns TEXT NOT NULL,
  PRIMARY KEY (ns, entity, permission, object)
);

CREATE INDEX IF NOT EXISTS relations_inverted_key ON relations(ns, object, permission, entity);

CREATE TABLE IF NOT EXISTS keys (
  id TEXT NOT NULL,
  secret_hash TEXT NOT NULL,
  ns TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS keys_by_ns ON keys(ns);
