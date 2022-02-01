import * as grpc from '@grpc/grpc-js'

import { CheckPermissionInput, CheckPermissionResponse, PantherConfig } from "./types";
import { PermissionPantherClient } from './pb/main_grpc_pb'
import { CheckDirectReq, CheckDirectRes } from './pb/permissions_pb'
import { promisify } from "util"
import { PermissionDenied } from "./errors";

export default class PermissionPanther {
  key: string
  target: string
  client: PermissionPantherClient
  constructor(config: PantherConfig) {
    this.key = config.key
    this.target = config.endpoint
    this.client = new PermissionPantherClient(this.target, grpc.credentials.createInsecure())
  }

  /**
   * Check permission
   */
  async CheckPermission(input: CheckPermissionInput): Promise<CheckPermissionResponse> {
    return new Promise((resolve, reject) => {
      const req = new CheckDirectReq()
      req.setKey(this.key)
      req.setEntity(input.entity)
      req.setPermission(input.permission)
      req.setObject(input.object)
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
}
