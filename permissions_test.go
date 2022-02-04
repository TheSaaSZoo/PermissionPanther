package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/utils"
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
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list entity permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success with filter")
		perm := "access"
		relations, err := ListEntityPermissions("nspc", "user1", &perm)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success")
		relations, err := ListObjectPermissions("nspc", "obj1", nil)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "access"
		relations, err := ListObjectPermissions("nspc", "obj1", &perm)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission empty with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "accesseeeee"
		relations, err := ListObjectPermissions("nspc", "obj1", &perm)
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("Upsert relation", func(t *testing.T) {
		log.Println("\n\n\n### test upsert relation")
		err := UpsertRelation("nspc", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)

		// Get it to validate
		perm := "tperm"
		relations, err := ListEntityPermissions("nspc", "tuser", &perm)
		utils.HandleTestError(t, err)
		if len(relations) != 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("Delete relation", func(t *testing.T) {
		log.Println("\n\n\n### test delete relation")
		err := DeleteRelation("nspc", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)

		// Get it to validate
		relations, err := ListEntityPermissions("nspc", "tuser", nil)
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})
}
