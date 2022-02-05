import * as grpc from '@grpc/grpc-js'

import { CheckPermissionInput, CheckPermissionResponse, ListEntityRelationsInput, ListObjectRelationsInput, ListRelationsResponse, PantherConfig, Relationship } from "./types";
import { PermissionPantherClient } from './pb/main_grpc_pb'
import { CheckDirectReq, ListEntityRelationsReq, ListObjectRelationsReq, RelationReq } from './pb/permissions_pb'
import { PermissionDenied } from "./errors";

export default class PermissionPanther {
  key: string
  target: string
  client: PermissionPantherClient
  constructor(config: PantherConfig) {
    this.key = config.key
    this.target = config.endpoint
    this.client = new PermissionPantherClient(this.target, grpc.credentials.createSsl())
  }

  /**
   * Checks whether an entity has a permission on an object. Optionally specify explicity deny permission, and group inheritance checks.
   */
  async CheckPermission(input: CheckPermissionInput): Promise<CheckPermissionResponse> {
    return new Promise((resolve, reject) => {
      const req = new CheckDirectReq()
      req.setKey(this.key)
      req.setEntity(input.entity)
      req.setPermission(input.permission)
      req.setObject(input.object)
      if (input.denyPermission) {
        req.setDenypermission(input.denyPermission)
      }
      if (input.inheritance === false) {
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
  async ListEntityRelations(input: ListEntityRelationsInput): Promise<ListRelationsResponse> {
    return new Promise((resolve, reject) => {
      const req = new ListEntityRelationsReq()
      req.setKey(this.key)
      req.setEntity(input.entity)
      if (input.permission) {
        req.setPermission(input.permission)
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
          offset: "",
          relations: rel
        })
      })
    })
  }

  /**
   * Lists an object's relations to find what entities have permission on it. Optionally specify a `permission` to look for entities who have a specific permission on the object.
   */
  async ListObjectRelations(input: ListObjectRelationsInput): Promise<ListRelationsResponse> {
    return new Promise((resolve, reject) => {
      const req = new ListObjectRelationsReq()
      req.setKey(this.key)
      req.setObject(input.object)
      if (input.permission) {
        req.setPermission(input.permission)
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
          offset: "",
          relations: rel
        })
      })
    })
  }

  /**
   * Sets a permission. Is a no-op if the permission already exists.
   */
  async SetPermission(input: Relationship) {
    return new Promise((resolve, reject) => {
      const req = new RelationReq()
      req.setEntity(input.entity)
      req.setKey(this.key)
      req.setObject(input.object)
      req.setPermission(input.permission)
      this.client.setPermission(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(undefined)
      })
    })
  }

  /**
   * Removes a permission. Is a no-op if the permission does not exist.
   */
  async RemovePermission(input: Relationship) {
    return new Promise((resolve, reject) => {
      const req = new RelationReq()
      req.setEntity(input.entity)
      req.setKey(this.key)
      req.setObject(input.object)
      req.setPermission(input.permission)
      this.client.removePermission(req, (err, res) => {
        if (err) {
          switch (err.code) {
            case grpc.status.PERMISSION_DENIED:
              reject(new PermissionDenied())
            default:
              reject(err)
          }
        }
        resolve(undefined)
      })
    })
  }
}
