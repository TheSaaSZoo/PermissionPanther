
-- +migrate Up
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
  object TEXT NOT NULL,
  PRIMARY KEY(ns, group_name, entity)
);
CREATE INDEX IF NOT EXISTS pgm_inverted_pkey ON permission_group_membership(entity, group_name);

-- +migrate Down
DROP INDEX IF EXISTS pgm_inverted_pkey;
DROP TABLE IF EXISTS permission_group_membership;
DROP TABLE IF EXISTS permission_groups;
