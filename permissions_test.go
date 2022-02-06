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
		found, err := CheckPermissions("nspc", "obj1", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 1 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 1 success")
		found, err := CheckPermissions("nspc", "obj2", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 2 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 2 success")
		found, err := CheckPermissions("nspc", "obj3", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 3 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 3 success")
		found, err := CheckPermissions("nspc", "obj4", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 4 success", func(t *testing.T) {
		log.Println("\n\n\n### recursion 4 success")
		found, err := CheckPermissions("nspc", "obj5", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found == -1 {
			panic("Not found!")
		}
	})

	t.Run("recursion 5 fail", func(t *testing.T) {
		log.Println("\n\n\n### recursion 5 fail")
		found, err := CheckPermissions("nspc", "obj6", "access", "user1", 0, 4)
		utils.HandleTestError(t, err)
		if found > 0 {
			panic("Found!")
		}
	})

	t.Run("list entity permission success", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success")
		relations, err := ListEntityPermissions("nspc", "user1", "")
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list entity permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list entity permission success with filter")
		perm := "access"
		relations, err := ListEntityPermissions("nspc", "user1", perm)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list entity permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success")
		relations, err := ListObjectPermissions("nspc", "obj1", "")
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission success with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "access"
		relations, err := ListObjectPermissions("nspc", "obj1", perm)
		utils.HandleTestError(t, err)
		if len(relations) < 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("list object permission empty with filter", func(t *testing.T) {
		log.Println("\n\n\n### list object permission success with filter")
		perm := "accesseeeee"
		relations, err := ListObjectPermissions("nspc", "obj1", perm)
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to list object permission"))
		}
		log.Printf("Found relations %+v\n", relations)
	})

	t.Run("insert relation", func(t *testing.T) {
		log.Println("\n\n\n### test insert relation")
		applied, err := UpsertRelation("nspc", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied"))
		}

		// Get it to validate
		perm := "tperm"
		relations, err := ListEntityPermissions("nspc", "tuser", perm)
		utils.HandleTestError(t, err)
		if len(relations) != 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("insert relation already exists", func(t *testing.T) {
		log.Println("\n\n\n### test insert relation already exists")
		applied, err := UpsertRelation("nspc", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)
		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied"))
		}

		// Get it to validate
		perm := "tperm"
		relations, err := ListEntityPermissions("nspc", "tuser", perm)
		utils.HandleTestError(t, err)
		if len(relations) != 1 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("Delete relation", func(t *testing.T) {
		log.Println("\n\n\n### test delete relation")
		applied, err := DeleteRelation("nspc", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)
		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied"))
		}

		// Get it to validate
		relations, err := ListEntityPermissions("nspc", "tuser", "")
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	t.Run("Delete relation not exists", func(t *testing.T) {
		log.Println("\n\n\n### test delete relation not exists")
		applied, err := DeleteRelation("nspc", "tobj", "tperm", "tuser")
		utils.HandleTestError(t, err)
		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied"))
		}

		// Get it to validate
		relations, err := ListEntityPermissions("nspc", "tuser", "")
		utils.HandleTestError(t, err)
		if len(relations) != 0 {
			utils.HandleTestError(t, fmt.Errorf("failed to validate upsert"))
		}
	})

	// t.Run("wildcard entity relation", func(t *testing.T) {
	// 	log.Println("\n\n\n### test wildcard entity relation")
	// 	err := UpsertRelation("nspc", "t2obj", "t2perm", "*")
	// 	utils.HandleTestError(t, err)

	// 	// Validate
	// 	found := CheckPermissions("nspc", "t2obj", "t2perm", "user1", 0, 4)
	// 	log.Println("Found wildcard entity relation at recursion", found)
	// 	if found == -1 {
	// 		panic("Not found!")
	// 	}

	// 	// Delete
	// 	err = DeleteRelation("nspc", "t2obj", "t2perm", "*")
	// 	utils.HandleTestError(t, err)

	// 	// Check again
	// 	found = CheckPermissions("nspc", "t2obj", "t2perm", "user1", 0, 4)
	// 	log.Println("Found wildcard relation at recursion", found)
	// 	if found != -1 {
	// 		panic("Not found!")
	// 	}
	// })

	// t.Run("wildcard object relation", func(t *testing.T) {
	// 	log.Println("\n\n\n### test wildcard object relation")
	// 	err := UpsertRelation("nspc", "*", "t2perm", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Validate
	// 	found := CheckPermissions("nspc", "yeye", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard object relation at recursion", found)
	// 	if found == -1 {
	// 		panic("Not found!")
	// 	}

	// 	// Delete
	// 	err = DeleteRelation("nspc", "*", "t2perm", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Check again
	// 	found = CheckPermissions("nspc", "yeye", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard relation at recursion", found)
	// 	if found != -1 {
	// 		panic("Not found!")
	// 	}
	// })

	// t.Run("wildcard permission relation", func(t *testing.T) {
	// 	log.Println("\n\n\n### test wildcard permission relation")
	// 	err := UpsertRelation("nspc", "oo1", "*", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Validate
	// 	found := CheckPermissions("nspc", "oo1", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard permission relation at recursion", found)
	// 	if found == -1 {
	// 		panic("Not found!")
	// 	}

	// 	// Delete
	// 	err = DeleteRelation("nspc", "oo1", "*", "u2")
	// 	utils.HandleTestError(t, err)

	// 	// Check again
	// 	found = CheckPermissions("nspc", "oo1", "t2perm", "u2", 0, 4)
	// 	log.Println("Found wildcard relation at recursion", found)
	// 	if found != -1 {
	// 		panic("Not found!")
	// 	}
	// })
}
