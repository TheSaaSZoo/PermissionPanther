import { CheckPermissionInput, CheckPermissionResponse, ListEntityRelationsInput, ListObjectRelationsInput, ListRelationsResponse, PantherConfig, Relationship } from "./types";
import { PermissionPantherClient } from './pb/main_grpc_pb';
export default class PermissionPanther {
    keyID: string;
    keySecret: string;
    target: string;
    client: PermissionPantherClient;
    constructor(config: PantherConfig);
    /**
     * Checks whether an entity has a permission on an object. Optionally specify explicity deny permission, and group inheritance checks.
     */
    CheckPermission(input: CheckPermissionInput): Promise<CheckPermissionResponse>;
    /**
     * Lists an entity's relations to find what objects they have permission on. Optionally specify a `permission` to look for objects that the entity has a specific permission on.
     */
    ListEntityRelations(input: ListEntityRelationsInput): Promise<ListRelationsResponse>;
    /**
     * Lists an object's relations to find what entities have permission on it. Optionally specify a `permission` to look for entities who have a specific permission on the object.
     */
    ListObjectRelations(input: ListObjectRelationsInput): Promise<ListRelationsResponse>;
    /**
     * Sets a permission.
     * Returns whether the relation was created (did not exist).
     */
    SetPermission(input: Relationship): Promise<boolean>;
    /**
     * Removes a permission.
     * Returns whether the relation was deleted (existed).
     */
    RemovePermission(input: Relationship): Promise<unknown>;
}
