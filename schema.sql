CREATE TABLE IF NOT EXISTS relations (
  object TEXT NOT NULL,
  entity TEXT NOT NULL,
  permission TEXT NOT NULL,
  ns TEXT,
  PRIMARY KEY (ns, entity, permission, object)
);

CREATE INDEX IF NOT EXISTS relations_inverted_key ON relations(ns, object, permission, entity);
