package main

import (
	"log"
	"testing"
)

func TestGRPCClient(t *testing.T) {
	t.Run("direct success", func(t *testing.T) {
		log.Println("\n\n\n### direct success")
		found := CheckPermissions("nspc", "obj1", "access", "user1", 0, 4)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 1 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 1 success")
		found := CheckPermissions("nspc", "obj2", "access", "user1", 0, 4)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 2 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 2 success")
		found := CheckPermissions("nspc", "obj3", "access", "user1", 0, 4)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 3 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 3 success")
		found := CheckPermissions("nspc", "obj4", "access", "user1", 0, 4)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 4 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 4 success")
		found := CheckPermissions("nspc", "obj5", "access", "user1", 0, 4)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 5 fail", func(t *testing.T) {
		log.Println("\n\n\n### recursion 5 fail")
		found := CheckPermissions("nspc", "obj6", "access", "user1", 0, 4)
		if found > 0 {
			panic("Found!")
		}
	})
}
