package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/danthegoodman1/PermissionPanther/utils"
)

func TestPermissionGroups(t *testing.T) {
	t.Run("create and remove permission group", func(t *testing.T) {
		log.Println("\n\n\n### test create and remove permission group")
		applied, err := CreatePermissionGroup("testns", "test_g_1", []string{"TEST_PERM"})
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied create"))
		}

		applied, err = CreatePermissionGroup("testns", "test_g_1", []string{"TEST_PERM"})
		utils.HandleTestError(t, err)

		// Make sure we cannot do it again
		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied create"))
		}

		// Get it to validate
		conn, err := crdb.PGPool.Acquire(context.Background())
		utils.HandleTestError(t, err)
		group, err := query.New(conn).SelectPermissionGroup(context.Background(), query.SelectPermissionGroupParams{
			Ns:   "testns",
			Name: "test_g_1",
		})
		utils.HandleTestError(t, err)
		log.Println("Got permission group:")
		log.Println(group)

		// Delete the group
		applied, err = RemovePermissionGroup("testns", "test_g_1", true)
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied delete"))
		}

		// Make sure we cannot do it again
		applied, err = RemovePermissionGroup("testns", "test_g_1", true)
		utils.HandleTestError(t, err)

		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied delete"))
		}
	})

	t.Run("create and add member then remove member", func(t *testing.T) {
		log.Println("\n\n\n### test create and add member then remove member")
		applied, err := CreatePermissionGroup("testns", "test_g_2", []string{"TEST_PERM"})
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied create"))
		}

		applied, err = AddMemberToPermissionGroup("testns", "test_g_2", "test_ent_2", "test_obj_2")
		utils.HandleTestError(t, err)
		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied add member"))
		}

		// Get it to validate
		conn, err := crdb.PGPool.Acquire(context.Background())
		utils.HandleTestError(t, err)
		members, err := query.New(conn).ListEntitiesInPermissionGroup(context.Background(), "")
		utils.HandleTestError(t, err)
		if len(members) != 1 {
			utils.HandleTestError(t, fmt.Errorf("Failed to find single user in permission group"))
		}

		// Validate the relation
		_, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
			Ns:         "testns",
			Entity:     "test_ent_2",
			Permission: "test_g_2",
			Object:     "test_obj_2",
		})
		utils.HandleTestError(t, err)

		// TODO: Test propagation by checking the direct relation

		// Remove the member
		rows, err := query.New(conn).RemoveMemberFromPermissionGroup(context.Background(), query.RemoveMemberFromPermissionGroupParams{
			Ns:        "testns",
			GroupName: "test_g_2",
			Entity:    "test_ent_2",
		})
		utils.HandleTestError(t, err)

		if rows != 1 {
			utils.HandleTestError(t, fmt.Errorf("Did not have 1 row when removing users from group"))
		}

		// Delete the group
		applied, err = RemovePermissionGroup("testns", "test_g_2", true)
		utils.HandleTestError(t, err)

		// _, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
		// 	Ns:         "testns",
		// 	Entity:     "test_ent_2",
		// 	Permission: "test_g_2",
		// 	Object:     "test_obj_2",
		// })
		// if err == nil {
		// 	utils.HandleTestError(t, fmt.Errorf("found the relation"))
		// } else if err != pgx.ErrNoRows {
		// 	utils.HandleTestError(t, err)
		// }

		// _, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
		// 	Ns:         "testns",
		// 	Entity:     "test_ent_2",
		// 	Permission: "TEST_PERM",
		// 	Object:     "test_obj_2",
		// })
		// if err == nil {
		// 	utils.HandleTestError(t, fmt.Errorf("found the relation"))
		// } else if err != pgx.ErrNoRows {
		// 	utils.HandleTestError(t, err)
		// }

		// TODO: comment in the above

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied delete"))
		}
	})
}
