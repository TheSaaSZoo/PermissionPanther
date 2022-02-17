// Code generated by sqlc. DO NOT EDIT.

package query

import (
	"time"
)

type Key struct {
	ID            string
	SecretHash    string
	Ns            string
	CreatedAt     time.Time
	MaxRecursions int64
}

type PermissionGroup struct {
	Name  string
	Ns    string
	Perms []string
}

type PermissionGroupMembership struct {
	GroupName string
	Entity    string
	Ns        string
	Object    string
}

type Relation struct {
	Object     string
	Entity     string
	Permission string
	Ns         string
}
