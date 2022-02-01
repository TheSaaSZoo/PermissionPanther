export interface PantherConfig {
    key: string;
    endpoint: string;
}
export interface CheckPermissionInput {
    entity: string;
    permission: string;
    object: string;
}
export interface CheckPermissionResponse {
    valid: boolean;
    recursion: number;
}
