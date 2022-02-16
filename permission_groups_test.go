package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/danthegoodman1/PermissionPanther/utils"
	"github.com/jackc/pgx/v4"
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

		// Verify can't do again
		applied, err = AddMemberToPermissionGroup("testns", "test_g_2", "test_ent_2", "test_obj_2")
		utils.HandleTestError(t, err)
		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied add member"))
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

		_, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
			Ns:         "testns",
			Entity:     "test_ent_2",
			Permission: "TEST_PERM",
			Object:     "test_obj_2",
		})
		utils.HandleTestError(t, err)

		// Remove the member
		applied, err = RemoveMemberFromPermissionGroup("testns", "test_g_2", "test_ent_2", "test_obj_2")
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Did not apply removal"))
		}

		// Verify not applied
		applied, err = RemoveMemberFromPermissionGroup("testns", "test_g_2", "test_ent_2", "test_obj_2")
		utils.HandleTestError(t, err)

		if applied {
			utils.HandleTestError(t, fmt.Errorf("applied removal"))
		}

		_, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
			Ns:         "testns",
			Entity:     "test_ent_2",
			Permission: "test_g_2",
			Object:     "test_obj_2",
		})
		if err == nil {
			utils.HandleTestError(t, fmt.Errorf("found the relation"))
		} else if err != pgx.ErrNoRows {
			utils.HandleTestError(t, err)
		}

		_, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
			Ns:         "testns",
			Entity:     "test_ent_2",
			Permission: "TEST_PERM",
			Object:     "test_obj_2",
		})
		if err == nil {
			utils.HandleTestError(t, fmt.Errorf("found the relation"))
		} else if err != pgx.ErrNoRows {
			utils.HandleTestError(t, err)
		}

		// Delete the group
		applied, err = RemovePermissionGroup("testns", "test_g_2", true)
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied delete"))
		}
	})

	t.Run("permission add and remove from group propagation", func(t *testing.T) {
		log.Println("\n\n\n### test permission add and remove from group propagation")
		applied, err := CreatePermissionGroup("testns", "test_g_3", []string{"TEST_PERM_A"})
		utils.HandleTestError(t, err)

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied create"))
		}

		applied, err = AddMemberToPermissionGroup("testns", "test_g_3", "test_ent_3", "test_obj_3")
		utils.HandleTestError(t, err)
		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied add member"))
		}

		// Add a permission and propagate
		applied, err = AddPermissionToGroup("testns", "test_g_3", "TEST_PERM_B", true)
		utils.HandleTestError(t, err)
		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied add permission 1"))
		}

		// Verify add
		conn, err := crdb.PGPool.Acquire(context.Background())
		utils.HandleTestError(t, err)
		_, err = query.New(conn).CheckRelationDirect(context.Background(), query.CheckRelationDirectParams{
			Ns:         "testns",
			Entity:     "test_ent_3",
			Permission: "TEST_PERM_A",
			Object:     "test_obj_3",
		})
		utils.HandleTestError(t, err)

		// Add another propagate for later delete

		// Add permission no propagate

		// Verify no propagate

		// Remove permission propagate

		// Verify propagate

		// Remove permission no propagate

		// Verify still has permission

		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied delete"))
		}
	})

	t.Run("membership pagination test", func(t *testing.T) {
		log.Println("\n\n\n### test membership pagination test")
		applied, err := CreatePermissionGroup("testns", "test_g_4", []string{"TEST_PERM"})
		utils.HandleTestError(t, err)
		if !applied {
			utils.HandleTestError(t, fmt.Errorf("Not applied create"))
		}
	})
}

// Check add permission to group with both add to array and propagate
// Check add permission to group with both add to array and propagate (for the 2 removes below)
// Check add permission to group with both add to array and NO propagate
// Check remove permission from group with array and propagate
// Check remove permission from group with array and NO propagate
// Pagination test with offset (can just be one, but make sure offset works)
