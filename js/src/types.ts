export interface PantherConfig {
  // The endpoint of the PermissionPanther service, ex: `localhost:8080`
  endpoint: string

  // If true, will use an insecure connection
  insecure?: boolean
}

export interface CheckPermissionInput {
  /**
   * Optionally specify an explicity "deny" permission. If a direct relation between and object and an entity with this permission is found then the check will be invalid no matter what.
   */
  denyPermission?: string

  /**
   * Whether to look through group inheritance relationships, up to the server set max recursion. Setting to false will only look for direct relationships. Default `true`.
   */
  inheritance?: boolean
}

export interface CheckPermissionResponse {
  // Whether the permission was valid
  valid: boolean

  /**
   * At what recursion level the permission was found, `0` means they had direct permission
   *
   * `1` means the entity belonged to a group that had that permission on the object, etc.
   */
  recursion: number
}

export interface ListRelationsInput {

  /**
   * Optional filter of results, will only check for relationships with this permission.
   */
  permission?: string

  offset?: number
}

export interface ListRelationsResponse {
  relations: Relationship[]
}

export interface Relationship {
  entity: string
  permission: string
  object: string
}

export interface GroupMembership {
  entity: string
  object: string
}
