import * as grpc from '@grpc/grpc-js'

import { CheckPermissionInput, CheckPermissionResponse, GroupMembership, ListRelationsInput, ListRelationsResponse, PantherConfig, Relationship } from "./types";
import { PermissionPantherClient } from './pb/main_grpc_pb'
import { CheckDirectReq, CreatePermissionGroupReq, DeletePermissionGroupReq, ListEntityRelationsReq, ListObjectRelationsReq, ListPermissionGroupReq, ModifyPermissionGroupReq, RelationReq } from './pb/permissions_pb'
import { PermissionDenied } from "./errors";

export default class PermissionPanther {
  keyID: string
  keySecret: string
  target: string
  client: PermissionPantherClient
  constructor(keyId: string, keySecret: string, config?: PantherConfig) {
    this.keyID = keyId
    this.keySecret = keySecret
    if (config?.endpoint) {
      this.target = config.endpoint
    } else {
      this.target = 'https://api.permissionpanther.com'
    }
    if (config?.insecure === true) {
      this.client = new PermissionPantherClient(this.target, grpc.credentials.createInsecure())
    } else {
      this.client = new PermissionPantherClient(this.target, grpc.credentials.createSsl())
    }
  }

  /**
   * Checks whether an entity has a permission on an object. Optionally specify explicity deny permission, and group inheritance checks.
   */
  async CheckPermission(entity: string, permission: string, object: string, options?: CheckPermissionInput): Promise<CheckPermissionResponse> {
    return new Promise((resolve, reject) => {
      const req = new CheckDirectReq()
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setEntity(entity)
      req.setPermission(permission)
      req.setObject(object)
      if (options?.denyPermission) {
        req.setDenypermission(options.denyPermission)
      }
      if (options?.inheritance === false) {
        req.setRecursive(false)
      } else {
        req.setRecursive(true)
      }
      this.client.checkDirectPermission(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
          return
        }
        resolve({
          recursion: res.getRecursion(),
          valid: res.getValid()
        })
      })
    })
  }

  /**
   * Lists an entity's relations to find what objects they have permission on. Optionally specify a `permission` to look for objects that the entity has a specific permission on.
   */
  async ListEntityRelations(entity: string, options?: ListRelationsInput): Promise<ListRelationsResponse> {
    return new Promise((resolve, reject) => {
      const req = new ListEntityRelationsReq()
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setEntity(entity)
      if (options?.permission) {
        req.setPermission(options.permission)
      }
      if (options?.offset) {
        req.setOffset(options.offset)
      }
      this.client.listEntityRelations(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        const rel: Relationship[] = []
        for (const r of res.getRelationsList()) {
          rel.push({
            entity: r.getEntity(),
            object: r.getObject(),
            permission: r.getPermission()
          })
        }
        resolve({
          relations: rel
        })
      })
    })
  }

  /**
   * Lists an object's relations to find what entities have permission on it. Optionally specify a `permission` to look for entities who have a specific permission on the object.
   */
  async ListObjectRelations(object: string, options?: ListRelationsInput): Promise<ListRelationsResponse> {
    return new Promise((resolve, reject) => {
      const req = new ListObjectRelationsReq()
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setObject(object)
      if (options?.permission) {
        req.setPermission(options.permission)
      }
      if (options?.offset) {
        req.setOffset(options.offset)
      }
      this.client.listObjectRelations(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        const rel: Relationship[] = []
        for (const r of res.getRelationsList()) {
          rel.push({
            entity: r.getEntity(),
            object: r.getObject(),
            permission: r.getPermission()
          })
        }
        resolve({
          relations: rel
        })
      })
    })
  }

  /**
   * Sets a permission.
   * Returns whether the relation was created (did not exist).
   */
  async SetPermission(entity: string, permission: string, object: string): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const req = new RelationReq()
      req.setEntity(entity)
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setObject(object)
      req.setPermission(permission)
      this.client.setPermission(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(res.getApplied())
      })
    })
  }

  /**
   * Create a Permission Group. Returns whether it was created.
   * @param initialPermissions The list of initial permissions to add to the group
   */
  async CreatePermissionGroup(groupName: string, initialPermissions?: string[]): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const req = new CreatePermissionGroupReq()
      req.setGroupname(groupName)
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setPermissionsList(initialPermissions || [])
      this.client.createPermissionGroup(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(res.getApplied())
      })
    })
  }

  /**
   * @param initialPermissions The list of initial permissions to add to the group
   * Returns whether it was created.
   * @param propagate Whether after deleting the group, every member of this group will have their permissions removed that were included in the group. This has major performance implications for large groups.
   * Delete a Permission Group. Returns whether the group was successfully deleted.
   */
  async DeletePermissionGroup(groupName: string, propagate: boolean): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const req = new DeletePermissionGroupReq()
      req.setGroupname(groupName)
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setPropagate(propagate)
      this.client.deletePermissionGroup(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(res.getApplied())
      })
    })
  }

  /**
   * Adds a permission to a group. Returns whether the permission was added.
   * @param propagate Whether after deleting the group, every member of this group will have their permissions removed that were included in the group. This has major performance implications for large groups.
   */
  async AddPermissionToGroup(groupName: string, permission: string, propagate: boolean): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const req = new ModifyPermissionGroupReq()
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setGroupname(groupName)
      req.setPermission(permission)
      req.setPropagate(propagate)
      this.client.addPermissionToGroup(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(res.getApplied())
      })
    })
  }

  /**
   * Removes a permission from a group. Returns whether the permission was removed.
   * @param propagate Whether after deleting the group, every member of this group will have their permissions removed that were included in the group. This has major performance implications for large groups.
   */
  async RemovePermissionFromGroup(groupName: string, permission: string, propagate: boolean): Promise<boolean> {
    return new Promise((resolve, reject) => {
      const req = new ModifyPermissionGroupReq()
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setGroupname(groupName)
      req.setPropagate(propagate)
      req.setPermission(permission)
      this.client.removePermissionFromGroup(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(res.getApplied())
      })
    })
  }

  /**
   * Lists entities in a permission group.
   * @param entityOffset If provided, the pagination will continue from this entity
   */
  async ListEntitiesInPermissionGroup(groupName: string, entityOffset?: string): Promise<GroupMembership[]> {
    return new Promise((resolve, reject) => {
      const req = new ListPermissionGroupReq()
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setGroupname(groupName)
      req.setOffset(entityOffset || "")
      this.client.listEntitiesInGroup(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        const rel: GroupMembership[] = []
        for (const r of res.getMembersList()) {
          rel.push({
            entity: r.getEntity(),
            object: r.getObject()
          })
        }
        resolve(rel)
      })
    })
  }

  /**
   * Removes a permission.
   * Returns whether the relation was deleted (existed).
   */
  async RemovePermission(entity: string, permission: string, object: string) {
    return new Promise((resolve, reject) => {
      const req = new RelationReq()
      req.setEntity(entity)
      req.setKeyid(this.keyID)
      req.setKeysecret(this.keySecret)
      req.setObject(object)
      req.setPermission(permission)
      this.client.removePermission(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(res.getApplied())
      })
    })
  }

  /**
   * Give all permissions defined within this group, if it exists
   */
  PermissionGroup(groupName: string): string {
    return `$${groupName}`
  }

  /**
   * Inherit relationships from another permission on an object
   */
  Inherit(permission: string, object: string): string {
    return `~${object}#${permission}`
  }
}
