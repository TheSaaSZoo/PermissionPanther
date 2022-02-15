package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/utils"
)

func TestPermissionGroups(t *testing.T) {
	t.Run("direct success", func(t *testing.T) {
		log.Println("\n\n\n### direct success")
		found, err := CheckPermissions("testns", "obj1", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 1 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 1 success")
		found, err := CheckPermissions("testns", "obj2", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 2 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 2 success")
		found, err := CheckPermissions("testns", "obj3", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 3 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 3 success")
		found, err := CheckPermissions("testns", "obj4", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 4 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 4 success")
		found, err := CheckPermissions("testns", "obj5", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 5 fail", func(t *testing.T) {
		log.Println("\n\n\n### recursion 5 fail")
		found, err := CheckPermissions("testns", "obj6", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found > 0 {
			panic("Found!")
		}
	})

	t.Run("list entity permission success", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success")
		relations, err := ListEntityPermissions("testns", "user1", "")
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list entity permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success with filter")
		perm := "access"
		relations, err := ListEntityPermissions("testns", "user1", perm)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success")
		relations, err := ListObjectPermissions("testns", "obj1", "")
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "access"
		relations, err := ListObjectPermissions("testns", "obj1", perm)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission empty with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "accesseeeee"
		relations, err := ListObjectPermissions("testns", "obj1", perm)
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("insert relation", func(t *testing.T) {
		log.Println("\n\n\n### test insert relation")
		applied, err := UpsertRelation("testns", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied"))
		}

		// Get it to validate
		perm := "tperm"
		relations, err := ListEntityPermissions("testns", "tuser", perm)
		utils.HandleTestError(t, err)
		if len(relations) != 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("insert relation already exists", func(t *testing.T) {
		log.Println("\n\n\n### test insert relation already exists")
		applied, err := UpsertRelation("testns", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)
		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied"))
		}

		// Get it to validate
		perm := "tperm"
		relations, err := ListEntityPermissions("testns", "tuser", perm)
		utils.HandleTestError(t, err)
		if len(relations) != 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("Delete relation", func(t *testing.T) {
		log.Println("\n\n\n### test delete relation")
		applied, err := DeleteRelation("testns", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)
		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied"))
		}

		// Get it to validate
		relations, err := ListEntityPermissions("testns", "tuser", "")
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("Delete relation not exists", func(t *testing.T) {
		log.Println("\n\n\n### test delete relation not exists")
		applied, err := DeleteRelation("testns", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)
		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied"))
		}

		// Get it to validate
		relations, err := ListEntityPermissions("testns", "tuser", "")
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})
}
