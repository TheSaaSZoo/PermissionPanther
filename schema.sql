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
  max_recursions INT8 NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS keys_by_ns ON keys(ns);

CREATE TABLE IF NOT EXISTS permission_groups (
  name TEXT NOT NULL,
  ns TEXT NOT NULL,
  perms TEXT[] NOT NULL,
  PRIMARY KEY (ns, name)
);

CREATE TABLE IF NOT EXISTS permission_group_membership (
  group_name TEXT NOT NULL,
  entity TEXT NOT NULL,
  ns TEXT NOT NULL,
  object TEXT NOT NULL, -- for faster permission change propagation
  PRIMARY KEY(ns, group_name, entity)
);
CREATE INDEX IF NOT EXISTS pgm_inverted_pkey ON permission_group_membership(entity, group_name);
