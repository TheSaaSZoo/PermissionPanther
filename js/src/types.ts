export interface PantherConfig {
  // Your API Key
  key: string

  // The endpoint of the PermissionPanther service, ex: `localhost:8080`
  endpoint: string
}

export interface CheckPermissionInput {
  // The entity you are checking the permission for
  entity: string

  // The permission you are checking
  permission: string

  // The object to are checking if the `entity` has the `permission` on
  object: string
}

export interface CheckPermissionResponse {
  // Whether the permission was valid
  valid: boolean

  // At what recursion level the permission was found, `0` means they had direct permission
  //
  // `1` means they belonged to a group that had that permission, etc.
  recursion: number
}
