import { CheckPermissionInput, CheckPermissionResponse, PantherConfig } from "./types";
import { PermissionPantherClient } from './pb/main_grpc_pb';
export default class PermissionPanther {
    key: string;
    target: string;
    client: PermissionPantherClient;
    constructor(config: PantherConfig);
    /**
     * Check permission
     */
    CheckPermission(input: CheckPermissionInput): Promise<CheckPermissionResponse>;
}
