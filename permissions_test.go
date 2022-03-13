package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/TheSaaSZoo/PermissionPanther/utils"
)

func TestPermissions(t *testing.T) {
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
		relations, err := ListEntityPermissions("testns", "user1", "", 0)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list entity permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success with filter")
		perm := "access"
		relations, err := ListEntityPermissions("testns", "user1", perm, 0)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success")
		relations, err := ListObjectPermissions("testns", "obj1", "", 0)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "access"
		relations, err := ListObjectPermissions("testns", "obj1", perm, 0)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission empty with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "accesseeeee"
		relations, err := ListObjectPermissions("testns", "obj1", perm, 0)
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
		relations, err := ListEntityPermissions("testns", "tuser", perm, 0)
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
		relations, err := ListEntityPermissions("testns", "tuser", perm, 0)
		utils.HandleTestError(t, err)
		if len(relations) != 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}

		// Verify offset works
		relations, err = ListEntityPermissions("testns", "tuser", perm, 100)
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed offset"))
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
		relations, err := ListEntityPermissions("testns", "tuser", "", 0)
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
		relations, err := ListEntityPermissions("testns", "tuser", "", 0)
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	// t.Run("wildcard entity relation", func(t *testing.T) {
	// 	log.Println("\n\n\n### test wildcard entity relation")
	// 	err := UpsertRelation("testns", "t2obj", "t2perm", "*")
	// 	utils.HandleTestError(t, err)

	// 	// Validate
	// 	found := CheckPermissions("testns", "t2obj", "t2perm", "user1", 0, 4)
	// 	log.Println("Found wildcard entity relation at recursion", found)
	// 	if found == -1 {
	// 		panic("Not found!")
	// 	}

	// 	// Delete
	// 	err = DeleteRelation("testns", "t2obj", "t2perm", "*")
	// 	utils.HandleTestError(t, err)

	// 	// Check again
	// 	found = CheckPermissions("testns", "t2obj", "t2perm", "user1", 0, 4)
	// 	log.Println("Found wildcard relation at recursion", found)
	// 	if found != -1 {
	// 		panic("Not found!")
	// 	}
	// })

	// t.Run("wildcard object relation", func(t *testing.T) {
	// 	log.Println("\n\n\n### test wildcard object relation")
	// 	err := UpsertRelation("testns", "*", "t2perm", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Validate
	// 	found := CheckPermissions("testns", "yeye", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard object relation at recursion", found)
	// 	if found == -1 {
	// 		panic("Not found!")
	// 	}

	// 	// Delete
	// 	err = DeleteRelation("testns", "*", "t2perm", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Check again
	// 	found = CheckPermissions("testns", "yeye", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard relation at recursion", found)
	// 	if found != -1 {
	// 		panic("Not found!")
	// 	}
	// })

	// t.Run("wildcard permission relation", func(t *testing.T) {
	// 	log.Println("\n\n\n### test wildcard permission relation")
	// 	err := UpsertRelation("testns", "oo1", "*", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Validate
	// 	found := CheckPermissions("testns", "oo1", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard permission relation at recursion", found)
	// 	if found == -1 {
	// 		panic("Not found!")
	// 	}

	// 	// Delete
	// 	err = DeleteRelation("testns", "oo1", "*", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Check again
	// 	found = CheckPermissions("testns", "oo1", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard relation at recursion", found)
	// 	if found != -1 {
	// 		panic("Not found!")
	// 	}
	// })
}
