import { CheckPermissionInput, CheckPermissionResponse, GroupMembership, ListRelationsInput, ListRelationsResponse, PantherConfig } from "./types";
import { PermissionPantherClient } from './pb/main_grpc_pb';
export default class PermissionPanther {
    keyID: string;
    keySecret: string;
    target: string;
    client: PermissionPantherClient;
    constructor(keyId: string, keySecret: string, config?: PantherConfig);
    /**
     * Checks whether an entity has a permission on an object. Optionally specify explicity deny permission, and group inheritance checks.
     */
    CheckPermission(entity: string, permission: string, object: string, options?: CheckPermissionInput): Promise<CheckPermissionResponse>;
    /**
     * Lists an entity's relations to find what objects they have permission on. Optionally specify a `permission` to look for objects that the entity has a specific permission on.
     */
    ListEntityRelations(entity: string, options?: ListRelationsInput): Promise<ListRelationsResponse>;
    /**
     * Lists an object's relations to find what entities have permission on it. Optionally specify a `permission` to look for entities who have a specific permission on the object.
     */
    ListObjectRelations(object: string, options?: ListRelationsInput): Promise<ListRelationsResponse>;
    /**
     * Sets a permission.
     * Returns whether the relation was created (did not exist).
     */
    SetPermission(entity: string, permission: string, object: string): Promise<boolean>;
    /**
     * Create a Permission Group. Returns whether it was created.
     * @param initialPermissions The list of initial permissions to add to the group
     */
    CreatePermissionGroup(groupName: string, initialPermissions?: string[]): Promise<boolean>;
    /**
     * @param initialPermissions The list of initial permissions to add to the group
     * Returns whether it was created.
     * @param propagate Whether after deleting the group, every member of this group will have their permissions removed that were included in the group. This has major performance implications for large groups.
     * Delete a Permission Group. Returns whether the group was successfully deleted.
     */
    DeletePermissionGroup(groupName: string, propagate: boolean): Promise<boolean>;
    /**
     * Adds a permission to a group. Returns whether the permission was added.
     * @param propagate Whether after deleting the group, every member of this group will have their permissions removed that were included in the group. This has major performance implications for large groups.
     */
    AddPermissionToGroup(groupName: string, permission: string, propagate: boolean): Promise<boolean>;
    /**
     * Removes a permission from a group. Returns whether the permission was removed.
     * @param propagate Whether after deleting the group, every member of this group will have their permissions removed that were included in the group. This has major performance implications for large groups.
     */
    RemovePermissionFromGroup(groupName: string, permission: string, propagate: boolean): Promise<boolean>;
    /**
     * Lists entities in a permission group.
     * @param entityOffset If provided, the pagination will continue from this entity
     */
    ListEntitiesInPermissionGroup(groupName: string, entityOffset?: string): Promise<GroupMembership[]>;
    /**
     * Removes a permission.
     * Returns whether the relation was deleted (existed).
     */
    RemovePermission(entity: string, permission: string, object: string): Promise<unknown>;
    /**
     * Give all permissions defined within this group, if it exists
     */
    PermissionGroup(groupName: string): string;
    /**
     * Inherit relationships from another permission on an object
     */
    Inherit(permission: string, object: string): string;
}
