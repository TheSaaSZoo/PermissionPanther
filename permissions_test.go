package main

import (
	"fmt"
	"log"
	"testing"
)

func TestPermissions(t *testing.T) {
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

	t.Run("list entity permission success", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success")
		relations, err := ListEntityPermissions("nspc", "user1", nil)
		HandleTestError(t, err)
		if len(relations) < 1 {
			HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list entity permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success with filter")
		perm := "access"
		relations, err := ListEntityPermissions("nspc", "user1", &perm)
		HandleTestError(t, err)
		if len(relations) < 1 {
			HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success")
		relations, err := ListObjectPermissions("nspc", "obj1", nil)
		HandleTestError(t, err)
		if len(relations) < 1 {
			HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "access"
		relations, err := ListObjectPermissions("nspc", "obj1", &perm)
		HandleTestError(t, err)
		if len(relations) < 1 {
			HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission empty with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "accesseeeee"
		relations, err := ListObjectPermissions("nspc", "obj1", &perm)
		HandleTestError(t, err)
		if len(relations) != 0 {
			HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})
}
