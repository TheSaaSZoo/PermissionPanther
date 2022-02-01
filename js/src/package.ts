import { CheckPermissionInput, CheckPermissionResponse, PantherConfig } from "./types";
import grpc from '@grpc/grpc-js'
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
    const req = new CheckDirectReq()
    req.setKey(this.key)
    req.setEntity(input.entity)
    req.setPermission(input.permission)
    req.setObject(input.object)
    try {
      const res = await promisify(this.client.checkDirectPermission)(req) as CheckDirectRes
      return {
        recursion: res.getRecursion(),
        valid: res.getValid()
      }
    } catch (error: any) {
      switch (error.code) {
        case grpc.status.PERMISSION_DENIED:
          throw new PermissionDenied()

        default:
          throw new Error(`unknown error, details: ${error.details}`)
      }
    }
  }
}
